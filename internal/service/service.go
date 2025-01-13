package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tube-nutrition/internal/database"
	"tube-nutrition/internal/database/mapper"
	"tube-nutrition/internal/model"
	"tube-nutrition/internal/utils"
)

type NutritionService interface {
	NutritionPlan(ctx utils.MyContext, accessToken string) (model.NutritionPlan, error)
}

type ImplNutritionService struct {
	repo      database.NutritionRepository
	jwtSecret []byte
}

func NewNutritionService(repo database.NutritionRepository, jwtSecret []byte) *ImplNutritionService {
	return &ImplNutritionService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *ImplNutritionService) NutritionPlan(ctx utils.MyContext, accessToken string) (model.NutritionPlan, error) {
	profile, err := FetchUserProfile(accessToken)
	if err != nil {
		return model.NutritionPlan{}, fmt.Errorf("failed to fetch user profile: %w", err)
	}

	totalCalories, proteins, fats, carbohydrates := calculateDailyNeeds(profile)

	dbDishes, err := s.repo.GetDishes(ctx, totalCalories, proteins, fats, carbohydrates)
	if err != nil {
		return model.NutritionPlan{}, fmt.Errorf("failed to fetch dishes: %w", err)
	}

	dishes := mapper.FromDBDishes(dbDishes)

	plan := model.NutritionPlan{
		Calories:      totalCalories,
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
		Dishes:        dishes,
	}

	return plan, nil
}

func FetchUserProfile(accessToken string) (model.Profile, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://tube-profile:82/api/profile/"), nil)
	if err != nil {
		return model.Profile{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return model.Profile{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Profile{}, fmt.Errorf("failed to fetch user profile, status code: %d", resp.StatusCode)
	}

	var profile model.Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return model.Profile{}, err
	}

	return profile, nil
}

func calculateDailyNeeds(profile model.Profile) (int, float32, float32, float32) {
	var bmr float32

	if profile.Gender == "м" {
		bmr = 88.36 + (13.4 * profile.Weight) + (4.8 * profile.Height) - (5.7 * float32(profile.Age))
	} else {
		bmr = 447.6 + (9.2 * profile.Weight) + (3.1 * profile.Height) - (4.3 * float32(profile.Age))
	}

	var multiplier float32
	switch profile.Goal {
	case "набор":
		multiplier = 1.2
	case "похудение":
		multiplier = 0.8
	default:
		multiplier = 1.0
	}

	totalCalories := int(bmr * multiplier)
	proteins := float32(0.3 * float32(totalCalories) / 4)
	fats := float32(0.25 * float32(totalCalories) / 9)
	carbohydrates := float32(0.45 * float32(totalCalories) / 4)

	return totalCalories, proteins, fats, carbohydrates
}

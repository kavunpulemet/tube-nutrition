package mapper

import (
	"tube-nutrition/internal/database"
	"tube-nutrition/internal/model"
)

func FromDBDishes(dbDishes []database.Dish) []model.Dish {
	serviceDishes := make([]model.Dish, len(dbDishes))

	for i, dbDish := range dbDishes {
		serviceDishes[i] = model.Dish{
			Name:        dbDish.Name,
			Description: dbDish.Description,
			Calories:    dbDish.Calories,
			Proteins:    dbDish.Proteins,
			Carbs:       dbDish.Carbs,
			Fats:        dbDish.Fats,
		}
	}

	return serviceDishes
}

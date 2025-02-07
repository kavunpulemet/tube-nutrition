package model

type Profile struct {
	UserID string  `json:"user_id"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	Goal   string  `json:"goal"`
}

type NutritionPlan struct {
	Calories      int     `json:"calories"`
	Proteins      float32 `json:"proteins"`
	Carbohydrates float32 `json:"carbohydrates"`
	Fats          float32 `json:"fats"`
	Dishes        []Dish  `json:"dishes"`
}

type Dish struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Calories    int     `json:"calories"`
	Proteins    float32 `json:"proteins"`
	Carbs       float32 `json:"carbs"`
	Fats        float32 `json:"fats"`
}

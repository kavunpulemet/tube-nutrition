package database

type Dish struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Calories    int     `db:"calories"`
	Proteins    float32 `db:"proteins"`
	Carbs       float32 `db:"carbohydrates"`
	Fats        float32 `db:"fats"`
}

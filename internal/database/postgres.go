package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"tube-nutrition/internal/utils"
)

type NutritionRepository interface {
	GetDishes(ctx utils.MyContext, goal string) ([]Dish, error)
}

type NutritionPostgres struct {
	db *sqlx.DB
}

func NewNutritionPostgres(db *sqlx.DB) *NutritionPostgres {
	return &NutritionPostgres{db: db}
}

//go:embed sql/GetDishes.sql
var getDishes string

func (p *NutritionPostgres) GetDishes(ctx utils.MyContext, goal string) ([]Dish, error) {
	var dishes []Dish

	err := p.db.SelectContext(ctx.Ctx, &dishes, getDishes, goal)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

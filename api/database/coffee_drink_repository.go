package database

import (
	"context"
	"go-template/api/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type CoffeeDrinksRepository interface {
	GetCoffeeDrinks() (*model.CoffeeDrinks, error)
	GetCoffeeDrinksByTemperatureStyle(temperatureStyle model.TemperatureStyle) (*model.CoffeeDrinks, error)
	CreateCoffeeDrink(coffeeDrink model.CoffeeDrink) (*string, error)
}

func (db *PostgresDatabase) GetCoffeeDrinks() (*model.CoffeeDrinks, error) {
	coffeeDrinks := &model.CoffeeDrinks{}
	err := pgxscan.Select(context.Background(), db.Pool, coffeeDrinks, CoffeeDrinkQuery{}.SelectCoffeeDrinks())
	if err != nil {
		return nil, err
	}

	return coffeeDrinks, err
}

func (db *PostgresDatabase) GetCoffeeDrinksByTemperatureStyle(temperatureStyle model.TemperatureStyle) (*model.CoffeeDrinks, error) {
	coffeeDrinks := &model.CoffeeDrinks{}
	err := pgxscan.Select(
		context.Background(),
		db.Pool,
		coffeeDrinks,
		CoffeeDrinkQuery{}.SelectCoffeeDrinksByTemperatureStyle(),
		temperatureStyle,
	)
	if err != nil {
		return nil, err
	}

	return coffeeDrinks, err
}

func (db *PostgresDatabase) CreateCoffeeDrink(coffeeDrink model.CoffeeDrink) (*string, error) {
	coffeeDrink.Id = uuid.New().String()

	row := db.Pool.QueryRow(
		context.Background(),
		CoffeeDrinkQuery{}.InsertCoffeeDrink(),
		coffeeDrink.Id,
		coffeeDrink.Name,
		coffeeDrink.Description,
		coffeeDrink.Origin,
		coffeeDrink.TemperatureStyle,
	)

	var createdId *string
	err := row.Scan(&createdId)
	if err != nil {
		return nil, err
	}

	return createdId, err
}

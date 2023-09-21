package handler

import (
	"fmt"
	"net/http"

	"go-template/api/custom"
	"go-template/api/database"
	"go-template/api/model"

	"github.com/labstack/echo/v4"
)

func GetCoffeeDrinks(ctx echo.Context) error {
	coffeeDrinksParams := model.GetCoffeeDrinksParams{}
	err := ctx.Bind(&coffeeDrinksParams)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("unable to bind request data: %s", err))
	}

	repository := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.CoffeeDrinksRepository)

	if coffeeDrinksParams.TemperatureStyle != nil {
		temperatureStyle := *coffeeDrinksParams.TemperatureStyle

		switch temperatureStyle {
		case model.HOT:
			hotCoffeeDrinks, err := repository.GetCoffeeDrinksByTemperatureStyle(temperatureStyle)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to get %s coffee drinks: %s", temperatureStyle, err))
			}

			return ctx.JSON(http.StatusOK, hotCoffeeDrinks)
		case model.COLD:
			coldCoffeeDrinks, err := repository.GetCoffeeDrinksByTemperatureStyle(temperatureStyle)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to get %s coffee drinks: %s", temperatureStyle, err))
			}

			return ctx.JSON(http.StatusOK, coldCoffeeDrinks)
		default:
			return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("invalid value for parameter temperatureStyle: %s", err))
		}
	}

	coffeeDrinks, err := repository.GetCoffeeDrinks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to get coffee drinks: %s", err))
	}

	return ctx.JSON(http.StatusOK, coffeeDrinks)
}

func CreateCoffeeDrink(ctx echo.Context) error {
	coffeeDrink := model.CoffeeDrink{}
	err := ctx.Bind(&coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("unable to bind request data: %s", err))
	}

	err = model.Validator.Struct(&coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("invalid coffee drink: %s", err))
	}

	repository := ctx.Get(custom.DATABASE_MIDDLEWARE_KEY).(database.CoffeeDrinksRepository)
	coffeeDrinkId, err := repository.CreateCoffeeDrink(coffeeDrink)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to create coffee drink: %s", err))
	}

	return ctx.JSON(http.StatusCreated, coffeeDrinkId)
}

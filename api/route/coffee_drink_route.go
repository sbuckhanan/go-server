package route

import (
	"go-template/api/handler"

	"github.com/labstack/echo/v4"
)

const COFFEE_DRINKS_PATH = "/coffee-drinks"

func (echoRouter *EchoRouter) AddCoffeeDrinksRoutes(baseGroup *echo.Group) {
	coffeeDrinksGroup := baseGroup.Group(COFFEE_DRINKS_PATH)
	coffeeDrinksGroup.GET("", handler.GetCoffeeDrinks)
	coffeeDrinksGroup.POST("", handler.CreateCoffeeDrink)
}

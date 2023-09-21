package model

const (
	COLD TemperatureStyle = "COLD"
	HOT  TemperatureStyle = "HOT"
)

type TemperatureStyle string

type CoffeeDrink struct {
	Description      string           `json:"description" validate:"required"`
	Id               string           `json:"id"`
	Name             string           `json:"name" validate:"required,max=32"`
	Origin           *string          `json:"origin,omitempty" validate:"max=32"`
	TemperatureStyle TemperatureStyle `json:"temperatureStyle" validate:"required"`
}

type CoffeeDrinks = []CoffeeDrink

type GetCoffeeDrinksParams struct {
	TemperatureStyle *TemperatureStyle `query:"temperatureStyle"`
}

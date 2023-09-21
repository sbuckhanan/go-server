package factory

import (
	"go-template/api/model"
	"go-template/test/utility"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rs/zerolog/log"
)

type TemperatureStyleFake model.TemperatureStyle

type CoffeeDrinkFactory struct {
	Id               string  `fake:"{uuid}"`
	Name             string  `fake:"{word}"`
	Origin           *string `fake:"{word}"`
	Description      string  `fake:"{word}"`
	TemperatureStyle TemperatureStyleFake
}

type CoffeeDrinkFactoryOptions struct {
	Id               *string
	Name             *string
	Origin           *string
	Description      *string
	TemperatureStyle *model.TemperatureStyle
}

func NewCoffeeDrinkFactory() *CoffeeDrinkFactory {
	factory := &CoffeeDrinkFactory{}
	err := gofakeit.Struct(factory)
	if err != nil {
		log.Error().Msgf("unable to generate fake data: %s", err)
	}

	return factory
}

func (factory CoffeeDrinkFactory) Create(options *CoffeeDrinkFactoryOptions) model.CoffeeDrink {
	coffeeDrink := model.CoffeeDrink{
		Id:               factory.Id,
		Name:             factory.Name,
		Origin:           factory.Origin,
		Description:      factory.Description,
		TemperatureStyle: factory.TemperatureStyle.Fake(),
	}

	if options == nil {
		return coffeeDrink
	}
	if options.Id != nil {
		coffeeDrink.Id = *options.Id
	}

	return coffeeDrink
}

func (fake TemperatureStyleFake) Fake() model.TemperatureStyle {
	randomTemperatureStyle := utility.RandomValue([]model.TemperatureStyle{
		model.HOT,
		model.COLD,
	})

	return randomTemperatureStyle.(model.TemperatureStyle)
}

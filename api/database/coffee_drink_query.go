package database

type CoffeeDrinkQuery struct{}

func (CoffeeDrinkQuery) SelectCoffeeDrinks() string {
	return "SELECT * FROM coffee_drink"
}

func (CoffeeDrinkQuery) SelectCoffeeDrinksByTemperatureStyle() string {
	return `
		SELECT * FROM coffee_drink
		WHERE temperature_style = $1
	`
}

func (CoffeeDrinkQuery) InsertCoffeeDrink() string {
	return `
		INSERT INTO coffee_drink (id, name, description, origin, temperature_style)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
}

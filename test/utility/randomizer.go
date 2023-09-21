package utility

import "github.com/brianvoe/gofakeit/v6"

func RandomValue[T any](list []T) interface{} {
	randomIndex := gofakeit.IntRange(0, len(list)-1)

	return list[randomIndex]
}

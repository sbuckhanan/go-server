package handler

import (
	"bytes"
	"encoding/json"
	"go-template/api/custom"
	"go-template/api/model"
	"go-template/test/data/factory"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCoffeeDrinksHandler(t *testing.T) {
	t.Run("should return a list of coffee drinks", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()
		ctx := e.NewContext(request, recorder)
		databaseFake := DatabaseFake{}
		ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseFake)

		_ = GetCoffeeDrinks(ctx)

		assert.Equal(t, http.StatusOK, recorder.Code)
		expectedCoffeeDrinks, _ := databaseFake.GetCoffeeDrinks()
		data, _ := json.Marshal(expectedCoffeeDrinks)
		expectedCoffeeDrinksString := string(data)
		actualCoffeeDrinksString := strings.TrimSuffix(recorder.Body.String(), "\n")
		assert.Equal(t, expectedCoffeeDrinksString, actualCoffeeDrinksString)
	})

	t.Run("should return a list of", func(t *testing.T) {
		testCases := []struct {
			description      string
			temperatureStyle model.TemperatureStyle
		}{
			{description: "hot coffee drinks", temperatureStyle: hot},
			{description: "cold coffee drinks", temperatureStyle: cold},
		}
		for _, testCase := range testCases {
			t.Run(testCase.description, func(t *testing.T) {
				e := echo.New()
				request := httptest.NewRequest(http.MethodGet, "/", nil)
				request.URL.RawQuery = url.Values{"temperatureStyle": {string(testCase.temperatureStyle)}}.Encode()
				recorder := httptest.NewRecorder()
				ctx := e.NewContext(request, recorder)
				databaseFake := DatabaseFake{}
				ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseFake)

				_ = GetCoffeeDrinks(ctx)

				assert.Equal(t, http.StatusOK, recorder.Code)
				expectedCoffeeDrinks, _ := databaseFake.GetCoffeeDrinksByTemperatureStyle(testCase.temperatureStyle)
				data, _ := json.Marshal(expectedCoffeeDrinks)
				expectedCoffeeDrinksString := string(data)
				actualCoffeeDrinksString := strings.TrimSuffix(recorder.Body.String(), "\n")
				assert.Equal(t, expectedCoffeeDrinksString, actualCoffeeDrinksString)
			})
		}
	})

	// Example unit tests to cover negative scenarios
	t.Run("should return an error when the request data unable to bind to GetCoffeeDrinksParams", func(t *testing.T) {})
	t.Run("should return an error when the requested temperature style is invalid", func(t *testing.T) {})
	t.Run("should return an error when the database unable to retrieve list of coffee drinks", func(t *testing.T) {})
}

func TestCreateCoffeeDrinkHandler(t *testing.T) {
	t.Run("CreateCoffeeDrink", func(t *testing.T) {
		t.Run("should create a coffee drink", func(t *testing.T) {
			e := echo.New()
			coffeeDrink := factory.NewCoffeeDrinkFactory().Create(nil)
			data, _ := json.Marshal(coffeeDrink)
			requestBody := bytes.NewReader(data)
			request := httptest.NewRequest(http.MethodPost, "/", requestBody)
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			recorder := httptest.NewRecorder()
			ctx := e.NewContext(request, recorder)
			databaseFake := DatabaseFake{}
			ctx.Set(custom.DATABASE_MIDDLEWARE_KEY, &databaseFake)

			_ = CreateCoffeeDrink(ctx)

			assert.Equal(t, http.StatusCreated, recorder.Code)
			expected, _ := databaseFake.CreateCoffeeDrink(coffeeDrink)
			expectedBytes, _ := json.Marshal(expected)
			expectedIdString := string(expectedBytes)
			actualIdString := strings.TrimSuffix(recorder.Body.String(), "\n")
			assert.Equal(t, expectedIdString, actualIdString)
		})

		// Example unit tests to cover negative scenarios
		t.Run("should return an error when the request data unable to bind to CoffeeDrink", func(t *testing.T) {})
		t.Run("should return an error when the request data is invalid", func(t *testing.T) {})
		t.Run("should return an error when the database unable to create coffee drink", func(t *testing.T) {})
	})
}

var (
	hot            = model.HOT
	cold           = model.COLD
	hotCoffeeDrink = factory.NewCoffeeDrinkFactory().Create(&factory.CoffeeDrinkFactoryOptions{
		TemperatureStyle: &hot,
	})
	coldCoffeeDrink = factory.NewCoffeeDrinkFactory().Create(&factory.CoffeeDrinkFactoryOptions{
		TemperatureStyle: &cold,
	})
)

type DatabaseFake struct{}

func (db *DatabaseFake) GetCoffeeDrinks() (*model.CoffeeDrinks, error) {
	return &model.CoffeeDrinks{
		hotCoffeeDrink,
		coldCoffeeDrink,
	}, nil
}
func (db *DatabaseFake) GetCoffeeDrinksByTemperatureStyle(temperatureStyle model.TemperatureStyle) (*model.CoffeeDrinks, error) {
	if temperatureStyle == hot {
		return &model.CoffeeDrinks{hotCoffeeDrink}, nil
	} else if temperatureStyle == cold {
		return &model.CoffeeDrinks{coldCoffeeDrink}, nil
	} else {
		return nil, nil
	}
}
func (db *DatabaseFake) CreateCoffeeDrink(coffeeDrink model.CoffeeDrink) (*string, error) {
	return &coffeeDrink.Id, nil
}

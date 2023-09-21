package route

import (
	"go-template/api/custom"
	"go-template/api/database"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

const BASE_PATH = "api/v1"

type EchoRouter struct {
	*echo.Echo
}

func NewEchoRouter(db *database.PostgresDatabase) *EchoRouter {
	e := echo.New()
	e.Use(custom.DatabaseMiddleware(db))
	e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
				logger := zerolog.New(os.Stdout).
					With().
					Timestamp().
					Logger()
				logger.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Msg("request")

				return nil
			},
		}),
	)

	return &EchoRouter{e}
}

func (echoRouter *EchoRouter) RegisterRoutes() {
	baseGroup := echoRouter.Group(BASE_PATH)
	echoRouter.AddCoffeeDrinksRoutes(baseGroup)
}

package http

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/toy-simulator/internal/app/toy"
	"github.com/toy-simulator/internal/interface/http/controller"
	middlewares "github.com/toy-simulator/internal/interface/http/middleware"
)

type httpServer struct {
	api  *echo.Echo
	Port string
}

type HttpServer interface {
	StartServer() error
}

func NewHttpServer(
	ctx context.Context,
	port string,
	toyCar toy.Toy,
) httpServer {
	echoAPI := echo.New()

	echoAPI.Use(middleware.CORS())

	echoAPI.Validator = middlewares.NewCustomValidator(validator.New())

	logger := zerolog.Ctx(ctx)
	echoAPI.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Str("Method", v.Method).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	echoAPI.HTTPErrorHandler = middlewares.NewHttpErrorHandler(ctx, middlewares.NewErrorStatusCodeMaps()).Handler

	controller.MakeGetCarController(ctx, echoAPI, toyCar)
	controller.MakePlaceToyController(ctx, echoAPI, toyCar)
	controller.MakeMoveToyController(ctx, echoAPI, toyCar)
	controller.MakeLeftToyController(ctx, echoAPI, toyCar)
	controller.MakeRightToyController(ctx, echoAPI, toyCar)

	return httpServer{
		api:  echoAPI,
		Port: port,
	}
}

func (h httpServer) StartServer() error {
	return h.api.Start(h.Port)
}

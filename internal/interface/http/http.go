package http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/toy-simulator/internal/app/toy"
	"github.com/toy-simulator/internal/interface/http/controller"
)

type httpServer struct {
	api  *echo.Echo
	Port string
}

type HttpServer interface {
	StartServer() error
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewHttpServer(
	ctx context.Context,
	port string,
	toyCar toy.Toy,
) httpServer {
	echoAPI := echo.New()

	echoAPI.Validator = &CustomValidator{validator: validator.New()}

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

	echoAPI.HTTPErrorHandler = NewHttpErrorHandler(ctx, NewErrorStatusCodeMaps()).Handler

	controller.MakeGetCarController(ctx, echoAPI, toyCar)
	controller.MakePostCommandToyController(ctx, echoAPI, toyCar)

	return httpServer{
		api:  echoAPI,
		Port: port,
	}
}

func (h httpServer) StartServer() error {
	return h.api.Start(h.Port)
}

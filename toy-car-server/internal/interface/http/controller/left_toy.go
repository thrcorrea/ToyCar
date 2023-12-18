package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
)

func MakeLeftToyController(ctx context.Context, api *echo.Echo, toyCar toy.Toy) {
	api.POST("/toy/left", LeftToyCar(ctx, toyCar))
}

func LeftToyCar(ctx context.Context, toyCar toy.Toy) echo.HandlerFunc {
	return func(e echo.Context) error {

		err := toyCar.Left()
		if err != nil {
			return err
		}

		toyCarPosition, _ := toyCar.GetPosition()

		response := GetToyCarResponse{
			PositionX: toyCarPosition.PositionX,
			PositionY: toyCarPosition.PositionY,
			Direction: string(toyCarPosition.Direction),
		}

		return e.JSON(http.StatusAccepted, response)
	}
}

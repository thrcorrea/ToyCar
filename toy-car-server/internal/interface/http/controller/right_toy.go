package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
)

func MakeRightToyController(ctx context.Context, api *echo.Echo, toyCar toy.Toy) {
	api.POST("/toy/right", RightToyCar(ctx, toyCar))
}

func RightToyCar(ctx context.Context, toyCar toy.Toy) echo.HandlerFunc {
	return func(e echo.Context) error {

		err := toyCar.Right()
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

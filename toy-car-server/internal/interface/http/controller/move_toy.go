package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
)

func MakeMoveToyController(ctx context.Context, api *echo.Echo, toyCar toy.Toy) {
	api.POST("/toy/move", MoveToyCar(ctx, toyCar))
}

func MoveToyCar(ctx context.Context, toyCar toy.Toy) echo.HandlerFunc {
	return func(e echo.Context) error {

		err := toyCar.Move()
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

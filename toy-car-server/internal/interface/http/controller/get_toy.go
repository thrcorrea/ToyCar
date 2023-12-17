package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
)

type GetToyCarResponse struct {
	PositionX int    `json:"x"`
	PositionY int    `json:"y"`
	Direction string `json:"direction"`
}

func MakeGetCarController(ctx context.Context, api *echo.Echo, toyCar toy.Toy) {
	api.GET("/toy", GetToyCarPosition(ctx, toyCar))
}

func GetToyCarPosition(ctx context.Context, toyCar toy.Toy) echo.HandlerFunc {
	return func(e echo.Context) error {

		toyCarPosition, err := toyCar.GetPosition()

		if err != nil {
			return err
		}

		response := GetToyCarResponse{
			PositionX: toyCarPosition.PositionX,
			PositionY: toyCarPosition.PositionY,
			Direction: string(toyCarPosition.Direction),
		}

		return e.JSON(http.StatusOK, response)
	}
}

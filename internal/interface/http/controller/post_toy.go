package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/toy-simulator/internal/app/toy"
)

type PostCommandRequest struct {
	Command   string `json:"command" validate:"required,alpha"`
	PositionX int    `json:"x"`
	PositionY int    `json:"y"`
	Direction string `json:"direction"`
}

func MakePostCommandToyController(ctx context.Context, api *echo.Echo, toyCar toy.Toy) {
	api.POST("/toy", PostCommandToyCar(ctx, toyCar))
}

func PostCommandToyCar(ctx context.Context, toyCar toy.Toy) echo.HandlerFunc {
	return func(e echo.Context) error {
		var requestData PostCommandRequest
		logger := zerolog.Ctx(ctx)

		if err := e.Bind(&requestData); err != nil {
			logger.Error().Err(err).Msg("error processing data")
			return e.NoContent(http.StatusInternalServerError)
		}

		if err := e.Validate(requestData); err != nil {
			logger.Warn().Err(err).Msg("request data malformed")
			return e.NoContent(http.StatusBadRequest)
		}

		switch requestData.Command {
		case "PLACE":
			err := toyCar.Place(requestData.PositionX, requestData.PositionY, toy.Direction(requestData.Direction))
			if err != nil {
				return err
				// logger.Error().Err(err)
				// return e.NoContent(http.StatusBadRequest)
			}
		case "LEFT":
			err := toyCar.Left()
			if err != nil {
				logger.Error().Err(err)
				return e.NoContent(http.StatusBadRequest)
			}
		case "RIGHT":
			err := toyCar.Right()
			if err != nil {
				logger.Error().Err(err)
				return e.NoContent(http.StatusBadRequest)
			}
		case "MOVE":
			err := toyCar.Move()
			if err != nil {
				logger.Error().Err(err)
				return e.NoContent(http.StatusBadRequest)
			}
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

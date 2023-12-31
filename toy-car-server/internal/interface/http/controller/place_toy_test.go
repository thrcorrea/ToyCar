package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/toy-simulator/internal/app/toy"
	middlewares "github.com/toy-simulator/internal/interface/http/middleware"
	toyMock "github.com/toy-simulator/internal/mocks/app/toy"
)

func TestPlaceToyCar(t *testing.T) {
	toyCarMock := toyMock.NewToy(t)
	ctx := context.Background()

	type args struct {
		ctx         context.Context
		toyCar      toy.Toy
		requestBody string
	}
	tests := []struct {
		name          string
		args          args
		mockBehaviour func()
		wantErr       bool
		wantStatus    int
	}{
		{
			name: "Should return error if toycar returns error",
			args: args{
				ctx:         ctx,
				toyCar:      toyCarMock,
				requestBody: `{"x": 10, "y": 2, "direction": "SOUTH"}`,
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Place(10, 2, toy.SOUTH).Return(toy.ErrCantPlaceToyCarOffBoard).Once()
			},
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return bad request when request is invalid",
			args: args{
				ctx:         ctx,
				toyCar:      toyCarMock,
				requestBody: "",
			},
			mockBehaviour: func() {},
			wantErr:       false,
			wantStatus:    http.StatusBadRequest,
		},
		{
			name: "Should return OK and command is processed correctly",
			args: args{
				ctx:         ctx,
				toyCar:      toyCarMock,
				requestBody: `{"x":1,"y":2,"direction":"SOUTH"}`,
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Place(1, 2, toy.SOUTH).Return(nil).Once()
				toyCarMock.EXPECT().GetPosition().Return(&toy.ToyPosition{PositionX: 1, PositionY: 2, Direction: toy.SOUTH}, nil).Once()
			},
			wantErr:    false,
			wantStatus: http.StatusAccepted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e := echo.New()
			e.Validator = middlewares.NewCustomValidator(validator.New())
			e.HTTPErrorHandler = middlewares.NewHttpErrorHandler(ctx, middlewares.NewErrorStatusCodeMaps()).Handler
			req := httptest.NewRequest(http.MethodPost, "/toy/place", strings.NewReader(tt.args.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mockBehaviour()
			if err := PlaceToyCar(tt.args.ctx, tt.args.toyCar)(c); (err != nil) != tt.wantErr {
				t.Errorf("PlaceToyCar() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.Equal(t, tt.wantStatus, rec.Code)
				assert.Equal(t, tt.args.requestBody, strings.ReplaceAll(rec.Body.String(), "\n", ""))
			}
			toyCarMock.Mock.AssertExpectations(t)
		})
	}
}

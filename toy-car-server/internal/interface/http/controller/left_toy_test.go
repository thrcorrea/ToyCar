package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
	middlewares "github.com/toy-simulator/internal/interface/http/middleware"
	toyMock "github.com/toy-simulator/internal/mocks/app/toy"
)

func TestLeftToyCar(t *testing.T) {
	toyCarMock := toyMock.NewToy(t)
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		toyCar toy.Toy
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
				ctx:    ctx,
				toyCar: toyCarMock,
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Left().Return(toy.ErrCantMoveToyCarOffBoard).Once()
			},
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Should return OK and command is processed correctly",
			args: args{
				ctx:    ctx,
				toyCar: toyCarMock,
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Left().Return(nil).Once()
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
			req := httptest.NewRequest(http.MethodPost, "/toy/left", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mockBehaviour()

			if err := LeftToyCar(tt.args.ctx, tt.args.toyCar)(c); (err != nil) != tt.wantErr {
				t.Errorf("LeftToyCar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package controller

import (
	"context"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/toy-simulator/internal/app/toy"
)

func TestPlaceToyCar(t *testing.T) {
	type args struct {
		ctx    context.Context
		toyCar toy.Toy
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlaceToyCar(tt.args.ctx, tt.args.toyCar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaceToyCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

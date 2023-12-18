package middlewares

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator interface {
	Validate(i interface{}) error
}

type customValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) CustomValidator {
	return &customValidator{
		validator: validator,
	}
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

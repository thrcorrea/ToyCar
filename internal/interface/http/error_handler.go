package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type httpErrorHandler struct {
	statusCodes map[error]int
	logger      *zerolog.Logger
}

func NewHttpErrorHandler(ctx context.Context, errorStatusCodeMaps map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		statusCodes: errorStatusCodeMaps,
		logger:      zerolog.Ctx(ctx),
	}
}

func (self *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range self.statusCodes {
		if errors.Is(err, key) {
			return value
		}
	}

	return http.StatusInternalServerError
}

func unwrapRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}

		originalErr = internalErr
	}

	return originalErr
}

func (self *httpErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: unwrapRecursive(err).Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = map[string]interface{}{"message": err.Error()}
	}

	if code >= 500 {
		self.logger.Error().Err(err).Msg("error")
	} else {
		self.logger.Warn().Err(err).Msg("warning")
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			self.logger.Error().Err(err).Msg("error")
		}
	}
}

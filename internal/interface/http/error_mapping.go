package http

import (
	"net/http"

	"github.com/toy-simulator/internal/app/toy"
)

func NewErrorStatusCodeMaps() map[error]int {

	var errorStatusCodeMaps = make(map[error]int)
	errorStatusCodeMaps[toy.ErrCantMoveToyCarOffBoard] = http.StatusBadRequest
	errorStatusCodeMaps[toy.ErrCantPlaceToyCarOffBoard] = http.StatusBadRequest
	errorStatusCodeMaps[toy.ErrToyCarNotPlaced] = http.StatusBadRequest
	return errorStatusCodeMaps
}

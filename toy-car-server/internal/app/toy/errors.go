package toy

import "errors"

var ErrToyCarNotPlaced = errors.New("toy car is not placed")
var ErrCantMoveToyCarOffBoard = errors.New("cant move toy car off board")
var ErrCantPlaceToyCarOffBoard = errors.New("cant place toy car off board")

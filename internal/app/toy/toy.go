package toy

import (
	"errors"
	"fmt"
)

type Direction string

const (
	NORTH Direction = "NORTH"
	SOUTH Direction = "SOUTH"
	EAST  Direction = "EAST"
	WEST  Direction = "WEST"
)

type Toy interface {
	Place(x, y int, direction Direction) error
	Move() error
	Left() error
	Right() error
	Report() (string, error)
}

type toy struct {
	tableMaxLength int
	tableMaxHeight int
	Facing         Direction
	Position       []int
}

func NewToy(
	tableMaxlength int,
	tableMaxHeigth int,
) Toy {
	return &toy{
		tableMaxLength: tableMaxlength,
		tableMaxHeight: tableMaxHeigth,
	}
}

func CheckIfToyOnTable(position []int) error {
	if len(position) <= 0 {
		return errors.New("Toy car is not placed yet")
	}
	return nil
}

// Left implements Toy.
func (t *toy) Left() error {

	err := CheckIfToyOnTable(t.Position)
	if err != nil {
		return err
	}

	switch t.Facing {
	case NORTH:
		t.Facing = WEST
	case WEST:
		t.Facing = SOUTH
	case SOUTH:
		t.Facing = EAST
	case EAST:
		t.Facing = NORTH
	}

	return nil
}

// Move implements Toy.
func (t *toy) Move() error {

	err := CheckIfToyOnTable(t.Position)
	if err != nil {
		return err
	}

	switch t.Facing {
	case NORTH:
		if t.Position[1] <= 0 {
			return errors.New("Cant move toy off the board")
		}
		t.Position[1]--
	case WEST:
		if t.Position[0] <= 0 {
			return errors.New("Cant move toy off the board")
		}
		t.Position[0]--
	case SOUTH:
		if t.Position[1] == (t.tableMaxHeight - 1) {
			return errors.New("Cant move toy off the board")
		}
		t.Position[1]++
	case EAST:
		if t.Position[0] == (t.tableMaxLength - 1) {
			return errors.New("Cant move toy off the board")
		}
		t.Position[0]++
	}

	return nil
}

// Place implements Toy.
func (t *toy) Place(x int, y int, direction Direction) error {

	t.Facing = direction

	if x < 0 || x >= (t.tableMaxLength-1) || y < 0 || y >= (t.tableMaxHeight-1) {
		return errors.New("Cant place toy off the board")
	}

	t.Position = []int{x, y}

	return nil
}

// Report implements Toy.
func (t *toy) Report() (string, error) {

	err := CheckIfToyOnTable(t.Position)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d,%d,%s", t.Position[0], t.Position[1], t.Facing), nil
}

// Right implements Toy.
func (t *toy) Right() error {

	err := CheckIfToyOnTable(t.Position)
	if err != nil {
		return err
	}

	switch t.Facing {
	case NORTH:
		t.Facing = EAST
	case EAST:
		t.Facing = SOUTH
	case SOUTH:
		t.Facing = WEST
	case WEST:
		t.Facing = NORTH
	}

	return nil
}

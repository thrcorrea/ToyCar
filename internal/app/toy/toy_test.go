package toy

import (
	"fmt"
	"reflect"
	"testing"
)

const TABLE_MAX_HEIGTH = 5
const TABLE_MAX_LENGTH = 5
const DIRECTION Direction = SOUTH

func Test_toy_Left(t *testing.T) {
	type fields struct {
		tableMaxLength int
		tableMaxHeight int
		direction      Direction
		position       []int
	}
	tests := []struct {
		name         string
		fields       fields
		wantFacing   Direction
		wantPosition []int
		wantErr      bool
	}{
		{
			name: "Should return error when toy isnt placed for Left function",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      DIRECTION,
			},
			wantErr: true,
		},
		{
			name: "Should return SOUTH when previous position was WEST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      WEST,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   SOUTH,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return EAST when previous position was SOUTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      SOUTH,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   EAST,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return NORTH when previous position was EAST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      EAST,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   NORTH,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return WEST when previous position was NORTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      NORTH,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   WEST,
			wantPosition: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &toy{
				tableMaxLength: tt.fields.tableMaxLength,
				tableMaxHeight: tt.fields.tableMaxHeight,
				Facing:         tt.fields.direction,
				Position:       tt.fields.position,
			}
			err := tr.Left()
			if (err != nil) != tt.wantErr {
				t.Errorf("toy.Left() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if tt.wantFacing != tr.Facing {
					t.Errorf("toy.Left() facing = %v, wantFacing %v", tr.Facing, tt.wantFacing)
					return
				}

				if !reflect.DeepEqual(tt.wantPosition, tr.Position) {
					t.Errorf("toy.Left() Position = %v, wantPosition %v", tr.Position, tt.wantPosition)
					return
				}

			}
		})
	}
}

func Test_toy_Right(t *testing.T) {
	type fields struct {
		tableMaxLength int
		tableMaxHeight int
		direction      Direction
		position       []int
	}
	tests := []struct {
		name         string
		fields       fields
		wantFacing   Direction
		wantPosition []int
		wantErr      bool
	}{
		{
			name: "Should return error when toy isnt placed for Right function",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      DIRECTION,
			},
			wantErr: true,
		},
		{
			name: "Should return SOUTH when previous position was EAST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      EAST,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   SOUTH,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return EAST when previous position was NORTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      NORTH,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   EAST,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return NORTH when previous position was WEST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      WEST,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   NORTH,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return WEST when previous position was SOUTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				direction:      SOUTH,
				position:       []int{0, 0},
			},
			wantErr:      false,
			wantFacing:   WEST,
			wantPosition: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &toy{
				tableMaxLength: tt.fields.tableMaxLength,
				tableMaxHeight: tt.fields.tableMaxHeight,
				Facing:         tt.fields.direction,
				Position:       tt.fields.position,
			}
			if err := tr.Right(); (err != nil) != tt.wantErr {
				t.Errorf("toy.Right() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if tt.wantFacing != tr.Facing {
					t.Errorf("toy.Right() facing = %v, wantFacing %v", tr.Facing, tt.wantFacing)
					return
				}

				if !reflect.DeepEqual(tt.wantPosition, tr.Position) {
					t.Errorf("toy.Right() Position = %v, wantPosition %v", tr.Position, tt.wantPosition)
					return
				}

			}
		})
	}
}

func Test_toy_Report(t *testing.T) {
	type fields struct {
		tableMaxLength int
		tableMaxHeight int
		Facing         Direction
		Position       []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Should return error when toy isnt placed for Right function",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         DIRECTION,
			},
			wantErr: true,
		},
		{
			name: "Should return position correctly",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         DIRECTION,
				Position:       []int{1, 2},
			},
			wantErr: false,
			want:    fmt.Sprintf("%d,%d,%s", 1, 2, DIRECTION),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &toy{
				tableMaxLength: tt.fields.tableMaxLength,
				tableMaxHeight: tt.fields.tableMaxHeight,
				Facing:         tt.fields.Facing,
				Position:       tt.fields.Position,
			}
			got, err := tr.Report()
			if (err != nil) != tt.wantErr {
				t.Errorf("toy.Report() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toy.Report() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toy_Place(t *testing.T) {
	type fields struct {
		tableMaxLength int
		tableMaxHeight int
		Facing         Direction
		Position       []int
	}
	type args struct {
		x         int
		y         int
		direction Direction
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFacing   Direction
		wantPosition []int
		wantErr      bool
	}{
		{
			name: "Should place at the top of the board",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
			},
			args: args{
				x:         0,
				y:         0,
				direction: SOUTH,
			},
			wantFacing:   SOUTH,
			wantPosition: []int{0, 0},
		},
		{
			name: "Should return error when place off the board",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
			},
			args: args{
				x:         10,
				y:         0,
				direction: SOUTH,
			},
			wantErr: true,
		},
		{
			name: "Should return error when place off the board",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
			},
			args: args{
				x:         -1,
				y:         -1,
				direction: SOUTH,
			},
			wantErr: true,
		},
		{
			name: "Should return error when place off the board",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
			},
			args: args{
				x:         -1,
				y:         10,
				direction: SOUTH,
			},
			wantErr: true,
		},
		{
			name: "Should return error when place off the board",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
			},
			args: args{
				x:         -1,
				y:         0,
				direction: SOUTH,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &toy{
				tableMaxLength: tt.fields.tableMaxLength,
				tableMaxHeight: tt.fields.tableMaxHeight,
				Facing:         tt.fields.Facing,
				Position:       tt.fields.Position,
			}
			if err := tr.Place(tt.args.x, tt.args.y, tt.args.direction); (err != nil) != tt.wantErr {
				t.Errorf("toy.Place() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if tt.wantFacing != tr.Facing {
					t.Errorf("toy.Place() facing = %v, wantFacing %v", tr.Facing, tt.wantFacing)
					return
				}

				if !reflect.DeepEqual(tt.wantPosition, tr.Position) {
					t.Errorf("toy.Place() Position = %v, wantPosition %v", tr.Position, tt.wantPosition)
					return
				}

			}
		})
	}
}

func Test_toy_Move(t *testing.T) {
	type fields struct {
		tableMaxLength int
		tableMaxHeight int
		Facing         Direction
		Position       []int
	}
	tests := []struct {
		name         string
		fields       fields
		wantPosition []int
		wantErr      bool
	}{
		{
			name: "Should return error when toy isnt placed for Move function",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         DIRECTION,
			},
			wantErr: true,
		},
		{
			name: "Should move to SOUTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         SOUTH,
				Position:       []int{0, 0},
			},
			wantPosition: []int{0, 1},
			wantErr:      false,
		},
		{
			name: "Should move to NORTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         NORTH,
				Position:       []int{0, 2},
			},
			wantPosition: []int{0, 1},
			wantErr:      false,
		},
		{
			name: "Should move to WEST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         WEST,
				Position:       []int{1, 0},
			},
			wantPosition: []int{0, 0},
			wantErr:      false,
		},
		{
			name: "Should move to EAST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         EAST,
				Position:       []int{0, 0},
			},
			wantPosition: []int{1, 0},
			wantErr:      false,
		},
		{
			name: "Should return error when move lands of the board facing NORTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         NORTH,
				Position:       []int{0, 0},
			},
			wantErr: true,
		},
		{
			name: "Should return error when move lands of the board facing SOUTH",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         SOUTH,
				Position:       []int{0, 4},
			},
			wantErr: true,
		},
		{
			name: "Should return error when move lands of the board facing EAST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         EAST,
				Position:       []int{4, 0},
			},
			wantErr: true,
		},
		{
			name: "Should return error when move lands of the board facing WEST",
			fields: fields{
				tableMaxLength: TABLE_MAX_LENGTH,
				tableMaxHeight: TABLE_MAX_HEIGTH,
				Facing:         WEST,
				Position:       []int{0, 0},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &toy{
				tableMaxLength: tt.fields.tableMaxLength,
				tableMaxHeight: tt.fields.tableMaxHeight,
				Facing:         tt.fields.Facing,
				Position:       tt.fields.Position,
			}
			if err := tr.Move(); (err != nil) != tt.wantErr {
				t.Errorf("toy.Move() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {

				if !reflect.DeepEqual(tt.wantPosition, tr.Position) {
					t.Errorf("toy.Move() Position = %v, wantPosition %v", tr.Position, tt.wantPosition)
					return
				}

			}
		})
	}
}

package cli

import (
	"errors"
	"reflect"
	"testing"

	"github.com/toy-simulator/internal/app/toy"
	toyMock "github.com/toy-simulator/internal/mocks/app/toy"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    command
		wantErr bool
	}{
		{
			name: "Should return command struct correctly",
			args: args{
				input: "PLACE 1,1,SOUTH",
			},
			want: command{
				action:    "PLACE",
				positionX: 1,
				positionY: 1,
				direction: "SOUTH",
			},
			wantErr: false,
		},
		{
			name: "Should return error when only spaces is provided",
			args: args{
				input: " ",
			},
			wantErr: true,
		},
		{
			name: "Should return error when it fails to parse the arguments",
			args: args{
				input: "PLACE 1,2",
			},
			wantErr: true,
		},
		{
			name: "Should return error when position x is invalid",
			args: args{
				input: "PLACE a,1,SOUTH",
			},
			wantErr: true,
		},
		{
			name: "Should return error when position y is invalid",
			args: args{
				input: "PLACE 1,b,SOUTH",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleCommands(t *testing.T) {
	toyCarMock := toyMock.NewToy(t)

	type args struct {
		c          cliInstance
		cliCommand command
	}

	type test struct {
		name          string
		args          args
		mockBehaviour func()
		wantErr       bool
	}

	tests := []test{
		{
			name: "Should execute place command correctly",
			args: args{
				cliCommand: command{
					action:    "PLACE",
					positionX: 1,
					positionY: 2,
					direction: toy.SOUTH,
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Place(1, 2, toy.SOUTH).Return(nil).Once()
			},
		},
		{
			name: "Should return error when place command returns error",
			args: args{
				cliCommand: command{
					action:    "PLACE",
					positionX: 10,
					positionY: 2,
					direction: toy.SOUTH,
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Place(10, 2, toy.SOUTH).Return(errors.New("Invalid place")).Once()
			},
			wantErr: true,
		},
		{
			name: "Should call function report correctly for command report",
			args: args{
				cliCommand: command{
					action: "REPORT",
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Report().Return("Valid Report", nil).Once()
			},
			wantErr: false,
		},
		{
			name: "Should return error when report function returns error",
			args: args{
				cliCommand: command{
					action: "REPORT",
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Report().Return("", errors.New("ToyCar not placed")).Once()
			},
			wantErr: true,
		},
		{
			name: "Should return error when left command trhows error",
			args: args{
				cliCommand: command{
					action: "LEFT",
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Left().Return(errors.New("ToyCar not placed")).Once()
			},
			wantErr: true,
		},
		{
			name: "Should return error when right command trhows error",
			args: args{
				cliCommand: command{
					action: "RIGHT",
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Right().Return(errors.New("ToyCar not placed")).Once()
			},
			wantErr: true,
		},
		{
			name: "Should return error when move command trhows error",
			args: args{
				cliCommand: command{
					action: "MOVE",
				},
				c: cliInstance{
					toyCar: toyCarMock,
				},
			},
			mockBehaviour: func() {
				toyCarMock.EXPECT().Move().Return(errors.New("ToyCar not placed")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour()
			if err := handleCommands(tt.args.c, tt.args.cliCommand); (err != nil) != tt.wantErr {
				t.Errorf("handleCommands() error = %v, wantErr %v", err, tt.wantErr)
			}
			toyCarMock.Mock.AssertExpectations(t)
		})
	}
}

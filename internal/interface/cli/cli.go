package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toy-simulator/internal/app/toy"
)

type Cli interface {
	Start()
}

type cliInstance struct {
	toyCar toy.Toy
}

type command struct {
	action    string
	positionX int
	positionY int
	direction toy.Direction
}

func NewCli(
	toyCar toy.Toy,
) Cli {
	return &cliInstance{
		toyCar: toyCar,
	}
}

func parseInput(input string) (command, error) {
	var actions []string
	var positionX, positionY int
	var direction string
	var err error
	var result command

	if strings.TrimLeft(input, " ") == "" {
		return result, errors.New("No command found")
	}

	raw := strings.Split(input, " ")
	for _, action := range raw {
		if strings.TrimSpace(action) != "" {
			actions = append(actions, strings.ToUpper(action))
		}
	}

	if actions[0] == "PLACE" {
		arguments := strings.Split(actions[1], ",")
		if len(arguments) == 3 {
			positionX, err = strconv.Atoi(arguments[0])
			if err != nil {
				return result, errors.New("Invalid position x")
			}
			positionY, err = strconv.Atoi(arguments[1])
			if err != nil {
				return result, errors.New("Invalid position y")
			}
			direction = arguments[2]
		} else {
			return result, errors.New("Failed to parse arguments")
		}
	}

	result.action = actions[0]
	result.positionX = positionX
	result.positionY = positionY
	result.direction = toy.Direction(direction)

	return result, nil
}

func handleCommands(c cliInstance, cliCommand command) error {
	switch cliCommand.action {
	case "PLACE":
		err := c.toyCar.Place(cliCommand.positionX, cliCommand.positionY, cliCommand.direction)
		if err != nil {
			return err
		}
	case "REPORT":
		result, err := c.toyCar.Report()
		if err != nil {
			return err
		} else {
			fmt.Println(result)
		}
	case "LEFT":
		err := c.toyCar.Left()
		if err != nil {
			return err
		}
	case "RIGHT":
		err := c.toyCar.Right()
		if err != nil {
			return err
		}
	case "MOVE":
		err := c.toyCar.Move()
		if err != nil {
			return err
		}
	default:
		fmt.Println("Command not found")
	}

	return nil
}

func (c cliInstance) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Commandline Started")
	for scanner.Scan() {
		input := scanner.Text()

		cliCommand, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = handleCommands(c, cliCommand)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toy-simulator/internal/app/toy"
)

type Cli interface {
	Start()
}

type cli struct {
	toyCar toy.Toy
}

func NewCli(
	toyCar toy.Toy,
) Cli {
	return &cli{
		toyCar: toyCar,
	}
}

func (c cli) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CommandLine Started")
	for scanner.Scan() {
		comando := scanner.Text()
		// Processar o comando
		var actions []string
		var positionX, positionY int
		var direction string

		if comando != "" {
			raw := strings.Split(comando, " ")
			for _, action := range raw {
				if strings.TrimSpace(action) != "" {
					actions = append(actions, action)
				}
			}

			if len(actions) > 1 {
				arguments := strings.Split(actions[1], ",")
				if len(arguments) > 0 {
					positionX, _ = strconv.Atoi(arguments[0])
					positionY, _ = strconv.Atoi(arguments[1])
					direction = arguments[2]
				}
			}
		}

		switch actions[0] {
		case "PLACE":
			err := c.toyCar.Place(positionX, positionY, toy.Direction(direction))
			if err != nil {
				fmt.Println(err)
			}
		case "REPORT":
			result, err := c.toyCar.Report()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case "LEFT":
			err := c.toyCar.Left()
			if err != nil {
				fmt.Println(err)
			}
		case "RIGHT":
			err := c.toyCar.Right()
			if err != nil {
				fmt.Println(err)
			}
		case "MOVE":
			err := c.toyCar.Move()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

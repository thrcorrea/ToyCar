package main

import (
	"github.com/toy-simulator/internal/app/toy"
	"github.com/toy-simulator/internal/interface/cli"
)

func main() {

	const TABLE_TOP_SIZE = 5
	toyCar := toy.NewToy(TABLE_TOP_SIZE, TABLE_TOP_SIZE)

	cliInterace := cli.NewCli(toyCar)

	go func() {
		cliInterace.Start()
	}()

	select {}
}

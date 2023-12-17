package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/toy-simulator/internal/app/toy"
	"github.com/toy-simulator/internal/interface/cli"
	"github.com/toy-simulator/internal/interface/http"
)

func main() {
	const TABLE_TOP_SIZE = 5
	toyCar := toy.NewToy(TABLE_TOP_SIZE, TABLE_TOP_SIZE)

	logger := zerolog.New(os.Stdout)
	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	cliInterace := cli.NewCli(toyCar)

	httpServer := http.NewHttpServer(ctx, ":3000", toyCar)

	go func() {
		cliInterace.Start()
	}()

	go func() {
		if err := httpServer.StartServer(); err != nil {
			logger.Error().Err(err)
		}
	}()

	select {}
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/toy-simulator/internal/app/toy"
	"github.com/toy-simulator/internal/interface/cli"
	"github.com/toy-simulator/internal/interface/http"
)

func main() {
	logger := zerolog.New(os.Stdout)
	err := godotenv.Load()
	if err != nil {
		logger.Info().Msg("env file not found")
	}
	const TABLE_TOP_SIZE = 5
	toyCar := toy.NewToy(TABLE_TOP_SIZE, TABLE_TOP_SIZE)

	ctx := context.Background()
	ctx = logger.WithContext(ctx)

	cliInterace := cli.NewCli(toyCar)

	httpServer := http.NewHttpServer(ctx, fmt.Sprintf(":%s", os.Getenv("PORT")), toyCar)

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

package main

import (
	"context"
	"os"

	"example.com/bizday-api/internal/api"
	_ "example.com/bizday-api/internal/logger" // configures global settings of zerolog
	"example.com/bizday-api/internal/settings"
	"github.com/rs/zerolog/log"
)

func main() {
	appContext := context.Background()

	log.Info().Msg("Starting server")
	err := api.RunServer(appContext, settings.Port)

	if err != nil {
		log.Fatal().Err(err).Msg("")
		os.Exit(2)
	}

	log.Info().Msg("Stopping server")
}

package main

import (
	"context"
	"os"

	"example.com/calendar-api/api"
	_ "example.com/calendar-api/logger"
	"github.com/rs/zerolog/log"
)

const PORT uint16 = 3333

func main() {
	appContext := context.Background()

	log.Info().Msg("Starting server")
	err := api.RunServer(appContext, api.Router, PORT)

	if err != nil {
		log.Fatal().Err(err).Msg("")
		os.Exit(2)
	}
	log.Info().Msg("Stopping server")
}

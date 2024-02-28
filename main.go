package main

import (
	"context"
	"time"

	"example.com/calendar-api/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

const PORT uint16 = 3333

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	appContext := context.Background()

	log.Info().Msg("Starting server")
	api.RunServer(appContext, api.Router, PORT)
	log.Info().Msg("Stopping server")
}

package settings

import (
	"time"

	"github.com/rs/zerolog"
)

var Debug bool
var Port uint16
var LogLevel zerolog.Level
var Timezone *time.Location

func init() {
	Debug = boolFromEnv("DEBUG", false)
	Port = uint16FromEnv("PORT", 3333)
	LogLevel = logLevelFromEnv("LOG_LEVEL", zerolog.InfoLevel)

	tz, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		panic("invalid timezone")
	}

	Timezone = tz
}

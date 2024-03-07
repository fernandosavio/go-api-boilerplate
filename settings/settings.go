package settings

import (
	"github.com/rs/zerolog"
)

var Debug bool
var Port uint16
var LogLevel zerolog.Level

func init() {
	Debug = boolFromEnv("DEBUG", false)
	Port = uint16FromEnv("PORT", 3333)
	LogLevel = logLevelFromEnv("LOG_LEVEL", zerolog.InfoLevel)
}

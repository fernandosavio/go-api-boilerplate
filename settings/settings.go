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

	parsedLevel, error := zerolog.ParseLevel(stringFromEnv("LOG_LEVEL", "info"))

	LogLevel = parsedLevel
	if error != nil {
		LogLevel = zerolog.InfoLevel
	}
}

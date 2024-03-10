package settings

import (
	"math"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

// Parse an environment variable as a boolean.
//
// Returns defaultValue if not found on environment or if value found is invalid.
//
// Valid truthy values are: "1", "t", "T", "true", "TRUE", "True"
//
// Valid falsy values are: "0", "f", "F", "false", "FALSE", "False"
func boolFromEnv(key string, defaultValue bool) bool {
	result, error := strconv.ParseBool(os.Getenv(key))

	if error != nil {
		return defaultValue
	}

	return result
}

// Parse an environment variable as a uint64.
//
// Returns defaultValue if not found on environment or if value found could not be converted.
func uint64FromEnv(key string, defaultValue uint64) uint64 {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	number, error := strconv.ParseUint(value, 10, 64)

	if error != nil {
		return defaultValue
	}

	return number
}

// Parse an environment variable as a uint16.
//
// Returns defaultValue if not found on environment or if value found could not be converted.
func uint16FromEnv(key string, defaultValue uint16) uint16 {
	default64 := uint64(defaultValue)
	value := uint64FromEnv(key, default64)

	if value == default64 || value > math.MaxUint16 {
		return defaultValue
	}

	return uint16(value)
}

// Parse an environment variable as a string.
//
// Returns defaultValue if not found on environment.
func stringFromEnv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	return value
}

// Parse an environment variable as a log level.
//
// Returns defaultValue if not found on environment or if value found is invalid.
func logLevelFromEnv(key string, defaultValue zerolog.Level) zerolog.Level {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	parsedLevel, error := zerolog.ParseLevel(value)

	if error != nil {
		return defaultValue
	}

	return parsedLevel
}

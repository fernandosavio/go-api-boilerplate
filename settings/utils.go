package settings

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var truthyValues = map[string]bool{
	"yes":  true,
	"true": true,
	"1":    true,
	"on":   true,
}

func boolFromEnv(key string, defaultValue bool) bool {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	_, found = truthyValues[strings.ToLower(strings.TrimSpace(value))]
	return found
}

func uint64FromEnv(key string, defaultValue uint64) uint64 {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	number, error := strconv.ParseUint(value, 10, 32)

	if error != nil {
		return defaultValue
	}

	return number
}

func uint16FromEnv(key string, defaultValue uint16) uint16 {
	default64 := uint64(defaultValue)
	value := uint64FromEnv(key, default64)

	if value == default64 || value > math.MaxUint16 {
		return defaultValue
	}

	return uint16(value)
}

func stringFromEnv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	return value
}

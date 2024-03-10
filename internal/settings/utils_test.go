package settings

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestBoolFromEnv(t *testing.T) {
	tests := []struct {
		inputValue   string
		defaultValue bool
		expected     bool
	}{
		// True values
		{inputValue: "true", defaultValue: false, expected: true},
		{inputValue: "TRUE", defaultValue: false, expected: true},
		{inputValue: "t", defaultValue: false, expected: true},
		{inputValue: "T", defaultValue: false, expected: true},
		{inputValue: "1", defaultValue: false, expected: true},
		// False values
		{inputValue: "false", defaultValue: true, expected: false},
		{inputValue: "FALSE", defaultValue: true, expected: false},
		{inputValue: "f", defaultValue: true, expected: false},
		{inputValue: "F", defaultValue: true, expected: false},
		{inputValue: "0", defaultValue: true, expected: false},
		// Invalid should fallback to defaultValue
		{inputValue: "invalid", defaultValue: true, expected: true},
		{inputValue: "invalid", defaultValue: false, expected: false},
		// Absent values should fallback to defaultValue
		{inputValue: "", defaultValue: true, expected: true},
		{inputValue: "", defaultValue: false, expected: false},
	}

	const envKey = "TEST_BOOL"
	for _, tt := range tests {
		if tt.inputValue == "" {
			os.Unsetenv(envKey)
		} else {
			t.Setenv(envKey, tt.inputValue)
		}

		got := boolFromEnv(envKey, tt.defaultValue)

		if got != tt.expected {
			t.Errorf("boolFromEnv(%s, %v) = %v, expected %v (%s=%q)", envKey, tt.defaultValue, got, tt.expected, envKey, tt.inputValue)
		}
	}
}

func TestUint64FromEnv(t *testing.T) {
	tests := []struct {
		inputValue   string
		defaultValue uint64
		expected     uint64
	}{
		// Correct parsing
		{inputValue: "0", defaultValue: 9999, expected: 0},
		{inputValue: "12345", defaultValue: 9999, expected: 12345},
		{inputValue: "18446744073709551615", defaultValue: 9999, expected: 18446744073709551615},
		// Invalid should fallback to defaultValue
		{inputValue: "invalid", defaultValue: 9999, expected: 9999},
		{inputValue: "18446744073709551616", defaultValue: 9999, expected: 9999},
		// Absent values should fallback to defaultValue
		{inputValue: "", defaultValue: 9999, expected: 9999},
	}

	const envKey = "TEST_UINT64"
	for _, tt := range tests {
		if tt.inputValue == "" {
			os.Unsetenv(envKey)
		} else {
			t.Setenv(envKey, tt.inputValue)
		}

		got := uint64FromEnv(envKey, tt.defaultValue)

		if got != tt.expected {
			t.Errorf("uint64FromEnv(%s, %v) = %v, expected %v (%s=%q)", envKey, tt.defaultValue, got, tt.expected, envKey, tt.inputValue)
		}
	}
}

func TestUint16FromEnv(t *testing.T) {
	tests := []struct {
		inputValue   string
		defaultValue uint16
		expected     uint16
	}{
		// Correct parsing
		{inputValue: "0", defaultValue: 9999, expected: 0},
		{inputValue: "12345", defaultValue: 9999, expected: 12345},
		{inputValue: "65535", defaultValue: 9999, expected: 65535},
		// Invalid should fallback to defaultValue
		{inputValue: "invalid", defaultValue: 9999, expected: 9999},
		{inputValue: "65536", defaultValue: 9999, expected: 9999},
		// Absent values should fallback to defaultValue
		{inputValue: "", defaultValue: 9999, expected: 9999},
	}

	const envKey = "TEST_UINT16"
	for _, tt := range tests {
		if tt.inputValue == "" {
			os.Unsetenv(envKey)
		} else {
			t.Setenv(envKey, tt.inputValue)
		}

		got := uint16FromEnv(envKey, tt.defaultValue)

		if got != tt.expected {
			t.Errorf("uint16FromEnv(%s, %v) = %v, expected %v (%s=%q)", envKey, tt.defaultValue, got, tt.expected, envKey, tt.inputValue)
		}
	}
}

func TestStringFromEnv(t *testing.T) {
	tests := []struct {
		inputValue   string
		defaultValue string
		expected     string
	}{
		// Correct parsing
		{inputValue: "abc", defaultValue: "default", expected: "abc"},
		{inputValue: "12345", defaultValue: "default", expected: "12345"},
		{inputValue: "  abc  ", defaultValue: "default", expected: "  abc  "},
		// Absent values should fallback to defaultValue
		{inputValue: "", defaultValue: "default", expected: "default"},
	}

	const envKey = "TEST_STRING"
	for _, tt := range tests {
		if tt.inputValue == "" {
			os.Unsetenv(envKey)
		} else {
			t.Setenv(envKey, tt.inputValue)
		}

		got := stringFromEnv(envKey, tt.defaultValue)

		if got != tt.expected {
			t.Errorf("stringFromEnv(%s, %v) = %v, expected %v (%s=%q)", envKey, tt.defaultValue, got, tt.expected, envKey, tt.inputValue)
		}
	}
}

func TestLogLevelFromEnv(t *testing.T) {
	tests := []struct {
		inputValue   string
		defaultValue zerolog.Level
		expected     zerolog.Level
	}{
		// Trace
		{inputValue: "trace", defaultValue: zerolog.Disabled, expected: zerolog.TraceLevel},
		{inputValue: "TRACE", defaultValue: zerolog.Disabled, expected: zerolog.TraceLevel},
		{inputValue: "-1", defaultValue: zerolog.Disabled, expected: zerolog.TraceLevel},
		// Debug
		{inputValue: "debug", defaultValue: zerolog.Disabled, expected: zerolog.DebugLevel},
		{inputValue: "DEBUG", defaultValue: zerolog.Disabled, expected: zerolog.DebugLevel},
		{inputValue: "0", defaultValue: zerolog.Disabled, expected: zerolog.DebugLevel},
		// Info
		{inputValue: "info", defaultValue: zerolog.Disabled, expected: zerolog.InfoLevel},
		{inputValue: "INFO", defaultValue: zerolog.Disabled, expected: zerolog.InfoLevel},
		{inputValue: "1", defaultValue: zerolog.Disabled, expected: zerolog.InfoLevel},
		// Warning
		{inputValue: "warn", defaultValue: zerolog.Disabled, expected: zerolog.WarnLevel},
		{inputValue: "WARN", defaultValue: zerolog.Disabled, expected: zerolog.WarnLevel},
		{inputValue: "2", defaultValue: zerolog.Disabled, expected: zerolog.WarnLevel},
		// Error
		{inputValue: "error", defaultValue: zerolog.Disabled, expected: zerolog.ErrorLevel},
		{inputValue: "ERROR", defaultValue: zerolog.Disabled, expected: zerolog.ErrorLevel},
		{inputValue: "3", defaultValue: zerolog.Disabled, expected: zerolog.ErrorLevel},
		// Fatal
		{inputValue: "fatal", defaultValue: zerolog.Disabled, expected: zerolog.FatalLevel},
		{inputValue: "FATAL", defaultValue: zerolog.Disabled, expected: zerolog.FatalLevel},
		{inputValue: "4", defaultValue: zerolog.Disabled, expected: zerolog.FatalLevel},
		// Panic
		{inputValue: "panic", defaultValue: zerolog.Disabled, expected: zerolog.PanicLevel},
		{inputValue: "PANIC", defaultValue: zerolog.Disabled, expected: zerolog.PanicLevel},
		{inputValue: "5", defaultValue: zerolog.Disabled, expected: zerolog.PanicLevel},
		// Disabled
		{inputValue: "disabled", defaultValue: zerolog.NoLevel, expected: zerolog.Disabled},
		{inputValue: "DISABLED", defaultValue: zerolog.NoLevel, expected: zerolog.Disabled},
		{inputValue: "7", defaultValue: zerolog.NoLevel, expected: zerolog.Disabled},
		// Absent values should fallback to defaultValue
		{inputValue: "", defaultValue: zerolog.Disabled, expected: zerolog.Disabled},
		// Invalid values should fallback to defaultValue
		{inputValue: "INVALID", defaultValue: zerolog.Disabled, expected: zerolog.Disabled},
		{inputValue: "99999999", defaultValue: zerolog.Disabled, expected: zerolog.Disabled},
	}

	const envKey = "TEST_LOGLEVEL"
	for _, tt := range tests {
		if tt.inputValue == "" {
			os.Unsetenv(envKey)
		} else {
			t.Setenv(envKey, tt.inputValue)
		}

		got := logLevelFromEnv(envKey, tt.defaultValue)

		if got != tt.expected {
			t.Errorf("logLevelFromEnv(%s, %v) = %v, expected %v (%s=%q)", envKey, tt.defaultValue, got, tt.expected, envKey, tt.inputValue)
		}
	}
}

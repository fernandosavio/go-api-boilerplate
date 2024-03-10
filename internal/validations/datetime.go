package validations

import (
	"time"

	"example.com/calendar-api/internal/settings"
)

var timeZeroValue = time.Time{}
var formatDateWithTZ = "2006-01-02Z07:00"
var tzSuffix = time.Date(2000, 1, 1, 0, 0, 0, 0, settings.Timezone).Format("Z07:00")

func ParseDate(value string) (time.Time, error) {
	if value == "" {
		return timeZeroValue, InvalidDateError
	}

	parsedDate, err := time.Parse(formatDateWithTZ, value+tzSuffix)

	if err != nil {
		return timeZeroValue, InvalidDateError
	}

	return parsedDate, nil
}

// Checks if a string is a valid ISO8601 date, like "2000-12-31"
func IsValidDate(value string) bool {
	parseDatetime, err := ParseDate(value)
	return err == nil && !parseDatetime.Equal(timeZeroValue)
}

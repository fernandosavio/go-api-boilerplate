package bizday

import "time"

type Holiday string

func validateDateString(value string) error {
	if value == "" {
		return InvalidDateError
	}

	parseDatetime, err := time.Parse(time.DateOnly, value)

	if err != nil || parseDatetime.IsZero() {
		return InvalidDateError
	}

	return nil
}

func NewHolidayFromTime(date time.Time) (*Holiday, error) {
	if date.IsZero() {
		return nil, InvalidDateError
	}

	result := Holiday(date.Format(time.DateOnly))
	return &result, nil
}

func NewHoliday(date string) (*Holiday, error) {
	if err := validateDateString(date); err != nil {
		return nil, err
	}

	result := Holiday(date)
	return &result, nil
}

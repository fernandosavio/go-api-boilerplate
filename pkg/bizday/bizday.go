package bizday

import "time"

type HolidayRepository interface {
	HasHoliday(time.Time) bool
}

type HolidayRegistry struct {
	repository HolidayRepository
}

func NewHolidayRegistry(holidayService HolidayRepository) (*HolidayRegistry, error) {
	if holidayService == nil {
		return nil, InvalidServiceError
	}

	return &HolidayRegistry{
		repository: holidayService,
	}, nil
}

func (hr *HolidayRegistry) IsBusinessDay(date time.Time) bool {
	if weekday := date.Weekday(); weekday == time.Saturday || weekday == time.Sunday {
		return false
	}

	return !hr.repository.HasHoliday(date)
}

func getNextWeekday(current time.Time) time.Time {
	switch current.Weekday() {
	case time.Friday:
		return current.AddDate(0, 0, 3)
	case time.Saturday:
		return current.AddDate(0, 0, 2)
	default:
		return current.AddDate(0, 0, 1)
	}
}

func (hr *HolidayRegistry) GetNextBusinessDayFrom(initialDate time.Time) time.Time {
	currentDate := initialDate

	for {
		currentDate = getNextWeekday(currentDate)

		if hr.repository.HasHoliday(currentDate) {
			continue
		}

		return currentDate
	}
}

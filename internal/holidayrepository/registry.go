package holidayrepository

import "time"

type NoHolidayRepo struct{}

func (hr *NoHolidayRepo) HasHoliday(time.Time) bool {
	return false
}

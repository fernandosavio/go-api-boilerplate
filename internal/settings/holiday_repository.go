package settings

import (
	"example.com/calendar-api/internal/bizday"
	"example.com/calendar-api/internal/holidayrepository"
)

var holidayRepository = holidayrepository.NoHolidayRepo{}

var HolidayService *bizday.HolidayRegistry

func init() {
	hs, err := bizday.NewHolidayRegistry(&holidayRepository)

	if err != nil {
		panic("could not create a new HolidayRegistry")
	}

	HolidayService = hs
}

package businessday

import (
	"errors"
	"net/http"
	"time"

	"example.com/calendar-api/internal/api/response"
	"example.com/calendar-api/internal/logger"
	"example.com/calendar-api/internal/settings"
	"example.com/calendar-api/internal/validations"
)

type NextBizdayResponse struct {
	Date string `json:"date"`
}

func validateNextBizdayRequest(r *http.Request) (time.Time, error) {
	qs := r.URL.Query()
	field := "from"

	if !qs.Has(field) {
		today := time.Now().In(settings.Timezone)
		return today, nil
	}

	input := qs.Get(field)

	parsedDate, err := validations.ParseDate(input)

	if err != nil {
		return time.Time{}, errors.New("invalid date")
	}

	return parsedDate, nil
}

func NextBizdayHandler(w http.ResponseWriter, r *http.Request) {
	from, err := validateNextBizdayRequest(r)

	if err != nil {
		correlationID, _ := logger.CorrelationIDFromRequest(r)
		response.JSON(w, r, http.StatusUnprocessableEntity, map[string]interface{}{
			"request_id": correlationID,
			"errors": map[string]string{
				"from": err.Error(),
			},
		})
		return
	}

	respBody := NextBizdayResponse{
		Date: settings.HolidayService.GetNextBusinessDayFrom(from).Format(time.DateOnly),
	}

	response.JSON(w, r, http.StatusOK, respBody)
}

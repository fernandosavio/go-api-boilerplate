package businessday

import (
	"net/http"
	"time"

	"example.com/bizday-api/internal/api/response"
	"example.com/bizday-api/internal/logger"
	"example.com/bizday-api/internal/settings"
	"example.com/bizday-api/internal/validation"
)

type NextBizdayResponse struct {
	Date string `json:"date"`
}

func validateNextBizdayRequest(r *http.Request) (time.Time, *response.ValidationErrorResponse) {
	fieldName := "from"
	qs := r.URL.Query()
	errs := response.ValidationError(logger.CorrelationIDFromRequest(r), nil)

	if !qs.Has(fieldName) {
		today := time.Now().In(settings.Timezone)
		return today, nil
	}

	input := qs.Get(fieldName)

	parsedDate, err := validation.ParseDate(input)

	if err != nil {
		errs.Add(fieldName, err.FieldError(response.LocQuerystring, response.ErrorContext{
			"value": input,
		}))
		return time.Time{}, errs
	}

	return parsedDate, nil
}

func NextBizdayHandler(w http.ResponseWriter, r *http.Request) {
	from, err := validateNextBizdayRequest(r)

	if err != nil {
		response.JSON(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	respBody := NextBizdayResponse{
		Date: settings.HolidayService.GetNextBusinessDayFrom(from).Format(time.DateOnly),
	}

	response.JSON(w, r, http.StatusOK, respBody)
}

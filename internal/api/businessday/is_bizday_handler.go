package businessday

import (
	"net/http"
	"time"

	"example.com/bizday-api/internal/api/response"
	"example.com/bizday-api/internal/logger"
	"example.com/bizday-api/internal/settings"
	"example.com/bizday-api/internal/validation"
	"github.com/go-chi/chi/v5"
)

type IsBizdayResponse struct {
	Date     string `json:"date"`
	IsBizDay bool   `json:"is_business_day"`
}

func validateIsBizdayRequest(r *http.Request) (time.Time, *response.ValidationErrorResponse) {
	fieldName := "date"
	input := chi.URLParam(r, fieldName)
	errs := response.ValidationError(logger.CorrelationIDFromRequest(r), nil)

	if input == "" {
		fieldError := validation.RequiredError.FieldError(response.LocPath, response.ErrorContext{})
		errs.Add(fieldName, fieldError)
		// RequiredError doesn't need to perform other validations so we return
		return time.Time{}, errs
	}

	parsedDate, err := validation.ParseDate(input)

	if err != nil {
		fieldError := validation.InvalidDateError.FieldError(response.LocPath, response.ErrorContext{})
		errs.Add(fieldName, fieldError)
	}

	if errs.HasErrors() {
		return time.Time{}, errs
	}

	return parsedDate, nil
}

func IsBizdayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := validateIsBizdayRequest(r)

	if err != nil {
		response.JSON(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	respBody := IsBizdayResponse{
		Date:     date.Format(time.DateOnly),
		IsBizDay: settings.HolidayService.IsBusinessDay(date),
	}
	response.JSON(w, r, http.StatusOK, respBody)
}

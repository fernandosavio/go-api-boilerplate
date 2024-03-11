package businessday

import (
	"errors"
	"net/http"
	"time"

	"example.com/bizday-api/internal/api/response"
	"example.com/bizday-api/internal/logger"
	"example.com/bizday-api/internal/settings"
	"example.com/bizday-api/internal/validations"
	"github.com/go-chi/chi/v5"
)

type IsBizdayResponse struct {
	Date     string `json:"date"`
	IsBizDay bool   `json:"is_business_day"`
}

func validateIsBizdayRequest(r *http.Request) (time.Time, error) {
	input := chi.URLParam(r, "date")

	if input == "" {
		return time.Time{}, errors.New("invalid date")
	}

	parsedDate, err := validations.ParseDate(input)

	if err != nil {
		return time.Time{}, errors.New("invalid date")
	}

	return parsedDate, nil
}

func IsBizdayHandler(w http.ResponseWriter, r *http.Request) {
	date, err := validateIsBizdayRequest(r)

	if err != nil {
		correlationID, _ := logger.CorrelationIDFromRequest(r)
		response.JSON(w, r, http.StatusUnprocessableEntity, map[string]interface{}{
			"request_id": correlationID,
			"errors": map[string]string{
				"date": err.Error(),
			},
		})
		return
	}

	respBody := IsBizdayResponse{
		Date:     date.Format(time.DateOnly),
		IsBizDay: settings.HolidayService.IsBusinessDay(date),
	}
	response.JSON(w, r, http.StatusOK, respBody)
}

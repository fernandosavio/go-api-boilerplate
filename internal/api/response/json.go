package response

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

const JSONMimeType = "application/json"

func JSONWithHeaders(w http.ResponseWriter, r *http.Request, status int, body any, extraHeaders map[string]string) {
	headers := w.Header()

	bytes := []byte{}
	if body != nil {
		responseBody, err := json.Marshal(body)
		// If body cannot be serialized, return error 500
		if err != nil {
			headers.Set("Content-Type", JSONMimeType)
			w.WriteHeader(http.StatusInternalServerError)

			hlog.FromRequest(r).
				Error().
				Msg("could not serialize response body to JSON")

			return
		}
		bytes = responseBody
	}

	for key, value := range extraHeaders {
		headers.Set(key, value)
	}
	headers.Set("Content-Type", JSONMimeType)
	w.WriteHeader(status)

	w.Write(bytes)
}

func JSON(w http.ResponseWriter, r *http.Request, status int, body any) {
	JSONWithHeaders(w, r, status, body, nil)
}

package logger

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/hlog"
)

var LoggerMiddleware func(next http.Handler) http.Handler = hlog.AccessHandler(accessMiddleware)

func accessMiddleware(r *http.Request, status, size int, duration time.Duration) {
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Int("status_code", status).
		Int("response_size_bytes", size).
		Dur("elapsed_ms", duration).
		Msg("incoming request")
}

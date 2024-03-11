package logger

import (
	"context"
	"net/http"
	"time"

	"example.com/bizday-api/internal/settings"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	zerolog.SetGlobalLevel(settings.LogLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

const (
	// Field name correlation ID will appear in logs
	CorrelationIDFieldName string = "correlation_id"
	// Field name correlation ID will appear in response headers (empty string to omit)
	CorrelationIDHeaderName string = "Correlation-ID"
)

// Inject logger in each request
var InjectLoggerMiddleware func(next http.Handler) http.Handler = hlog.NewHandler(log.Logger)

// Create a Correlation ID for each request and add it to the request/log context
var RequestIDMiddleware func(next http.Handler) http.Handler = hlog.RequestIDHandler(CorrelationIDFieldName, CorrelationIDHeaderName)

// Log request information after each request
var LogRequestMiddleware func(next http.Handler) http.Handler = hlog.AccessHandler(accessMiddleware)

func accessMiddleware(r *http.Request, status, size int, duration time.Duration) {
	hlog.FromRequest(r).Info().
		Str("method", r.Method).
		Stringer("url", r.URL).
		Int("status_code", status).
		Int("response_size", size).
		Dur("elapsed_ms", duration).
		Msg("request access log")
}

func CorrelationIDFromRequest(r *http.Request) (correlationID string, ok bool) {
	id, ok := hlog.IDFromRequest(r)
	correlationID = id.String()
	return
}

func CorrelationIDFromContext(ctx context.Context) (correlationID string, ok bool) {
	id, ok := hlog.IDFromCtx(ctx)
	correlationID = id.String()
	return
}

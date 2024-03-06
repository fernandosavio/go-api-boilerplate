package api

import (
	"fmt"
	"net/http"

	"example.com/calendar-api/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var Router *chi.Mux

func init() {
	r := chi.NewRouter()

	r.Use(
		middleware.Heartbeat("/healthcheck"),
		middleware.AllowContentType("application/json"),
		middleware.AllowContentEncoding("gzip", "deflate"),
		middleware.Compress(5, "application/json"),
		logger.InjectLoggerMiddleware,
		logger.RequestIDMiddleware,
		logger.LogRequestMiddleware,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		correlationId, _ := logger.CorrelationIDFromRequest(r)
		responseTxt := fmt.Sprintf(`{"request_id":"%s","message":"OK"}`, correlationId)
		w.Write([]byte(responseTxt))
	})

	Router = r
}

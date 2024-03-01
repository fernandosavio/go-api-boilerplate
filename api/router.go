package api

import (
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
		w.Write([]byte(`{"message":"OK"}`))
	})

	Router = r
}

package api

import (
	"net/http"

	"example.com/bizday-api/internal/api/businessday"
	"example.com/bizday-api/internal/api/response"
	"example.com/bizday-api/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Heartbeat("/healthcheck"),
		middleware.AllowContentType("application/json"),
		middleware.AllowContentEncoding("gzip", "deflate"),
		middleware.Compress(5, "application/json"),
		logger.InjectLoggerMiddleware,
		logger.RequestIDMiddleware,
		logger.LogRequestMiddleware,
	)

	// Default 404 handler
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		correlationId := logger.CorrelationIDFromRequest(r)
		response.JSON(w, r, http.StatusNotFound, response.Error404(correlationId))
	})

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/businessday", businessday.Router())
	})

	return router
}

func RegisterHandlers(server *http.Server) {
	server.Handler = Router()
}

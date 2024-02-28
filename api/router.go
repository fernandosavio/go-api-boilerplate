package api

import (
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
	)

	Router = r
}

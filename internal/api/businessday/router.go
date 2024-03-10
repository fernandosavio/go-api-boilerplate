package businessday

import "github.com/go-chi/chi/v5"

func Router() *chi.Mux {
	subrouter := chi.NewRouter()

	subrouter.Get("/next", NextBizdayHandler)
	subrouter.Get("/{date}", IsBizdayHandler)

	return subrouter
}

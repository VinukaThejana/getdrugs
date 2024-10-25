package router

import (
	env "github.com/VinukaThejana/getdrugs/internal/config"
	"github.com/go-chi/chi/v5"
	middlewares "github.com/go-chi/chi/v5/middleware"
)

// Init initializes the router
func Init(
	env *env.Env,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.Logger)
	r.Use(middlewares.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/drugs", nil)
	})

	return r
}
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/krispekla/pro-profile-ai-api/config"
)

func routes(app *config.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(CorsMiddleware(app))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(AuthMiddleware(app))
	r.Route("/api", func(r chi.Router) {
		r.Get("/user", userDetails(app))
		r.Get("/characters", getCharacters(app))
		r.Post("/characters", createCharacter(app))
		r.Get("/packages", getPackages(app))
		r.Get("/packages/:id", getPackageDetails(app))
		r.Post("/packages/buy", buyPackage(app))
		r.Post("/checkout/sessions", createCheckoutSession(app))
		r.Get("/checkout/sessions", retrieveCheckoutSession(app))
		r.Get("/buckets/all", getAllBuckets(app))
	})
	return r
}

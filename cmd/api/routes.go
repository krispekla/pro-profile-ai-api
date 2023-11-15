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
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(AuthMiddleware(app))
	r.Post("/login", login(app))
	r.Post("/register", register(app))
	r.Route("/app", func(r chi.Router) {
		r.Get("/user-details", userDetails(app))
	})
	r.Route("/api/v1/", func(r chi.Router) {
		// Apply the middleware to this router

		// Define your routes
		r.Post("/login", login(app))
		// Add more routes as needed
	})
	return r
}

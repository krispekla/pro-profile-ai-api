package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/krispekla/pro-profile-ai-api/config"
)

func routes(app *config.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Post("/login", login(app))
	r.Post("/register", register(app))
	r.Route("/app", func(r chi.Router) {
		r.Get("/user-details", userDetails(app))
	})
	return r
}

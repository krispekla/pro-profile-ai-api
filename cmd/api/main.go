package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/krispekla/pro-profile-ai-api/config"
)

func routes(app *config.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/ping", ping(app))
	r.Post("/login", login(app))
	r.Post("/register", register(app))
	r.Route("/app", func(r chi.Router) {
		r.Get("/user-details", userDetails(app))
	})
	return r
}

func main() {
	app := &config.Application{}
	app.CreateLoggers()
	flag.StringVar(&app.Addr, "addr", ":3002", "Port to run this service on")
	flag.Parse()

	srv := &http.Server{
		Addr:     app.Addr,
		Handler:  routes(app),
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

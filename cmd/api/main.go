package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", ping)
	srv := &http.Server{
		Addr:    ":3002",
		Handler: r,
	}
	log.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

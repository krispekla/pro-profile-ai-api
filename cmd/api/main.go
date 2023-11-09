package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Live"))
	})
	srv := &http.Server{
		Addr:    ":3002",
		Handler: r,
	}
	log.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	Addr string
}

func main() {
	cfg := &Config{}
	flag.StringVar(&cfg.Addr, "addr", ":3002", "Port to run this service on")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/ping", ping)
	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}
	log.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	Addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	cfg := &Config{}
	flag.StringVar(&cfg.Addr, "addr", ":3002", "Port to run this service on")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	r := chi.NewRouter()
	r.Get("/ping", app.ping)
	r.Post("/login", app.login)
	r.Post("/register", app.register)
	r.Route("/app", func(r chi.Router) {
		r.Get("/user-details", app.userDetails)
	})

	srv := &http.Server{
		Addr:     cfg.Addr,
		Handler:  r,
		ErrorLog: errorLog,
	}
	infoLog.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"flag"
	"net/http"

	"github.com/krispekla/pro-profile-ai-api/config"
)

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

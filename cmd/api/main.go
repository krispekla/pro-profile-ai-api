package main

import (
	"net/http"

	"github.com/krispekla/pro-profile-ai-api/config"
)

func main() {
	app := &config.Application{}
	app.LoadConfig()
	app.CreateLoggers()

	srv := &http.Server{
		Addr:     app.Addr,
		Handler:  routes(app),
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

package main

import (
	"net/http"

	"github.com/krispekla/pro-profile-ai-api/config"
	"github.com/krispekla/pro-profile-ai-api/storage"
)

func main() {
	app := &config.Application{}
	app.CreateLoggers()
	app.LoadConfig()
	app.SetR2Config()
	storage.OpenDb(app)

	srv := &http.Server{
		Addr:     app.Addr,
		Handler:  routes(app),
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Println("Starting server on port ", srv.Addr)
	err := srv.ListenAndServe()
	defer app.Db.Close()
	app.ErrorLog.Fatal(err)
}

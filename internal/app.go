package internal

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/krispekla/pro-profile-ai-api/internal/database"
	"github.com/krispekla/pro-profile-ai-api/internal/services"
)

type Application struct {
	Config    *AppConfig
	Db        *sql.DB
	R2Service *services.R2Service
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
}

type AppConfig struct {
	Addr          string
	JwtSecret     string
	StorageConfig *database.DbConn
	R2StorageCfg  *services.R2Config
}

func NewApplication(config *AppConfig, db *sql.DB, r2Service *services.R2Service, infoLog *log.Logger, errorLog *log.Logger) *Application {
	return &Application{
		Config:    config,
		Db:        db,
		R2Service: r2Service,
		ErrorLog:  errorLog,
		InfoLog:   infoLog,
	}
}

func (app *Application) Run() {
	srv := &http.Server{
		Addr:     app.Config.Addr,
		Handler:  routes(app),
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Println("Starting server on:", srv.Addr)
	err := srv.ListenAndServe()
	defer app.Db.Close()
	app.ErrorLog.Fatal(err)
}

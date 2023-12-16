package storage

import (
	"database/sql"
	"fmt"

	"github.com/krispekla/pro-profile-ai-api/config"
)

func OpenDb(app *config.Application) {
	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		app.StorageConfig.DbHost, app.StorageConfig.DbPort, app.StorageConfig.DbUser, app.StorageConfig.DbPassword, app.StorageConfig.DbName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		app.ErrorLog.Panicf("Error opening database connection: %v", err)
	}
	defer db.Close()
	app.Db = db
	app.InfoLog.Println("Database connection established")
}

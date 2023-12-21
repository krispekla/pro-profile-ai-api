package main

import (
	"log"

	"github.com/krispekla/pro-profile-ai-api/internal"
	"github.com/krispekla/pro-profile-ai-api/internal/database"
	"github.com/krispekla/pro-profile-ai-api/internal/services"
)

func main() {
	config := internal.LoadConfig()
	infoLog, errorLog := internal.CreateLoggers()
	db, err := database.SetupDatabase(config.StorageConfig)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	r2Service := services.NewR2Service(config.R2StorageCfg)
	app := internal.NewApplication(config, db, r2Service, infoLog, errorLog)
	app.Run()
}

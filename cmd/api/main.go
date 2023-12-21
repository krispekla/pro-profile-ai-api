package main

import (
	"github.com/krispekla/pro-profile-ai-api/internal"

	_ "github.com/lib/pq"
)

func main() {
	app := internal.NewApplication()
	app.Run()
}

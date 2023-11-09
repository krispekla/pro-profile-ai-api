package config

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Application struct {
	Addr     string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (app *Application) CreateLoggers() {
	app.InfoLog = log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
}

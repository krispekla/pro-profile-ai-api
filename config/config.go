package config

import (
	"log"
	"os"
)

type Application struct {
	addr     string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (app *Application) CreateLoggers() {
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func (app *Application) GetAddr() *string {
	return &app.addr
}

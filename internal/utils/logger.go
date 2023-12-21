package utils

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func CreateLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

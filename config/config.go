package config

import (
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/stripe/stripe-go/v76"
)

type Application struct {
	Addr       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
}

func (app *Application) CreateLoggers() {
	app.InfoLog = log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
}

func (app *Application) LoadConfig() {
	// TODO: Replace with viper
	flag.StringVar(&app.Addr, "addr", ":3002", "Port to run this service on")
	flag.StringVar(&app.DbHost, "db-host", "localhost", "Database host")
	flag.StringVar(&app.DbPort, "db-port", "5432", "Database port")
	flag.StringVar(&app.DbName, "db-name", "", "Database name")
	flag.StringVar(&app.DbUser, "db-user", "", "Database user")
	flag.StringVar(&app.DbPassword, "db-password", "", "Database password")
	flag.StringVar(&app.JwtSecret, "jwt-secret", "", "JWT secret for checking token validity")
	flag.StringVar(&stripe.Key, "stripe-secret", "", "Stripe secret key")
	flag.Parse()
}

package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/fatih/color"
	"github.com/stripe/stripe-go/v76"
)

type Application struct {
	Addr              string
	DbHost            string
	DbPort            string
	DbName            string
	DbUser            string
	DbPassword        string
	JwtSecret         string
	R2AccountId       string
	R2AccessKeyId     string
	R2AccessKeySecret string
	R2BucketName      string
	ErrorLog          *log.Logger
	InfoLog           *log.Logger
	R2Config          *aws.Config
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
	flag.StringVar(&app.R2AccountId, "r2-account-id", "", "R2 account id")
	flag.StringVar(&app.R2AccessKeyId, "r2-access-key-id", "", "R2 access key id")
	flag.StringVar(&app.R2AccessKeySecret, "r2-access-key-secret", "", "R2 access key secret")
	flag.StringVar(&app.R2BucketName, "r2-bucket-name", "", "R2 bucket name")
	flag.Parse()
}

func (app *Application) SetR2Config() {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", app.R2AccountId),
		}, nil
	})

	r2Config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(app.R2AccessKeyId, app.R2AccessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}
	app.R2Config = &r2Config
}

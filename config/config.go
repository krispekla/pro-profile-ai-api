package config

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func (app *Application) LoadConfig() {
	var envFilePath string
	pflag.StringVar(&envFilePath, "env", "./.env", "Path to .env file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigType("env")
	viper.SetConfigFile(envFilePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	// Load the values from the config file into the application struct
	app.Addr = viper.GetString("ppai_api_addr")
	app.DbHost = viper.GetString("ppai_api_db_host")
	app.DbPort = viper.GetString("ppai_api_db_port")
	app.DbName = viper.GetString("ppai_api_db_name")
	app.DbUser = viper.GetString("ppai_api_db_user")
	app.DbPassword = viper.GetString("ppai_api_db_password")
	app.JwtSecret = viper.GetString("ppai_api_supabase_secret")
	stripe.Key = viper.GetString("ppai_api_stripe_secret")
	app.R2AccountId = viper.GetString("ppai_api_r2_account_id")
	app.R2AccessKeyId = viper.GetString("ppai_api_r2_access_key_id")
	app.R2AccessKeySecret = viper.GetString("ppai_api_r2_access_key_secret")
	app.R2BucketName = viper.GetString("ppai_api_r2_bucket_name")
	app.InfoLog.Printf("Loaded config: %+v\n", envFilePath)
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
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	app.R2Config = &r2Config
}

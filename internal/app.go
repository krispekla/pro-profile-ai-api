package internal

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/krispekla/pro-profile-ai-api/internal/database"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v76"
)

type Application struct {
	Config   *AppConfig
	Db       *sql.DB
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

type AppConfig struct {
	Addr          string
	JwtSecret     string
	StorageConfig *database.DbConn
	R2StorageCfg  *R2StorageCfg
	R2Config      *aws.Config
}
type R2StorageCfg struct {
	R2AccountId       string
	R2AccessKeyId     string
	R2AccessKeySecret string
	R2BucketName      string
}

func NewApplication() *Application {
	app := &Application{}
	app.CreateLoggers()
	app.Config = loadConfig(app)
	app.SetR2Config()
	return app
}

func (app *Application) Run() {
	db, err := database.SetupDatabase(app.Config.StorageConfig) // Assume this function sets up your database
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	app.Db = db
	r := routes(app)
	srv := &http.Server{
		Addr:     app.Config.Addr,
		Handler:  r,
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Println("Starting server on:", srv.Addr)
	err = srv.ListenAndServe()
	defer app.Db.Close()
	app.ErrorLog.Fatal(err)
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) CreateLoggers() {
	app.InfoLog = log.New(os.Stdout, color.GreenString("INFO\t"), log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, color.RedString("ERROR\t"), log.Ldate|log.Ltime|log.Lshortfile)
}

func (app *Application) SetR2Config() {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", app.Config.R2StorageCfg.R2AccountId),
		}, nil
	})

	r2Config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(app.Config.R2StorageCfg.R2AccessKeyId, app.Config.R2StorageCfg.R2AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	app.Config.R2Config = &r2Config
}

func loadConfig(app *Application) *AppConfig {
	var envFilePath string
	cfg := &AppConfig{}
	storCfg := &database.DbConn{}
	r2Config := &R2StorageCfg{}

	pflag.StringVar(&envFilePath, "env", "../../.env", "Path to .env file")

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
	cfg.Addr = viper.GetString("ppai_api_addr")
	cfg.JwtSecret = viper.GetString("ppai_api_supabase_secret")

	storCfg.DbHost = viper.GetString("ppai_api_db_host")
	storCfg.DbPort = viper.GetString("ppai_api_db_port")
	storCfg.DbName = viper.GetString("ppai_api_db_name")
	storCfg.DbUser = viper.GetString("ppai_api_db_user")
	storCfg.DbPassword = viper.GetString("ppai_api_db_password")
	stripe.Key = viper.GetString("ppai_api_stripe_secret")

	r2Config.R2AccountId = viper.GetString("ppai_api_r2_account_id")
	r2Config.R2AccessKeyId = viper.GetString("ppai_api_r2_access_key_id")
	r2Config.R2AccessKeySecret = viper.GetString("ppai_api_r2_access_key_secret")
	r2Config.R2BucketName = viper.GetString("ppai_api_r2_bucket_name")

	cfg.StorageConfig = storCfg
	cfg.R2StorageCfg = r2Config
	app.InfoLog.Printf("Loaded config: %+v\n", envFilePath)
	return cfg
}

package config

import (
	"flag"
	"log"

	"github.com/krispekla/pro-profile-ai-api/internal"
	"github.com/krispekla/pro-profile-ai-api/internal/database"
	"github.com/krispekla/pro-profile-ai-api/internal/services"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v76"
)

func Load() *internal.AppConfig {
	var envFilePath string
	cfg := &internal.AppConfig{}
	storCfg := &database.DbConn{}
	r2Config := &services.R2Config{}

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
	return cfg
}

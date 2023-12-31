package services

import (
	"log"
	"os"

	configs "github.com/gulizay91/go-rest-api/config"
	"github.com/spf13/viper"
)

var config *configs.Config

func InitConfig() {
	v := viper.New()

	log.Printf(os.Getenv("SERVICE_ENVIRONMENT"))

	var environment = os.Getenv("SERVICE_ENVIRONMENT")
	configName := ".env"
	if environment != "" {
		configName = ".env." + environment
	}
	log.Printf(configName)
	//v.SetDefault("SERVER_PORT", "8080")
	//v.SetConfigType("dotenv")
	v.SetConfigType("yaml")
	v.SetConfigName(configName)
	v.AddConfigPath("../")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// override config for validation check
	// log.Printf(v.GetString("service.logLevel"))
	// v.Set("service.logLevel", "")
	// log.Printf(v.GetString("service.logLevel"))

	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}

	if err := config.Validate(); err != nil {
		panic(err)
	}
}
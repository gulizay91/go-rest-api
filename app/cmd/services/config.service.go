package services

import (
	"log"
	"os"
	"path/filepath"

	configs "github.com/gulizay91/go-rest-api/config"
	"github.com/spf13/viper"
)

var config *configs.Config

func InitConfig() {
	v := viper.New()

	log.Printf("env>SERVICE_ENVIRONMENT: %s", os.Getenv("SERVICE_ENVIRONMENT"))
	log.Printf("env>TEST_KEY: %s", os.Getenv("TEST_KEY"))
	log.Printf("env>cluster: %s", os.Getenv("cluster"))
	log.Printf("env>service.logLevel: %s", os.Getenv("service.logLevel"))
	log.Printf("env>service__name: %s", os.Getenv("service__name"))

	var environment = os.Getenv("SERVICE_ENVIRONMENT")
	configName := "env"
	if environment != "" {
		configName = "env." + environment
	}
	log.Printf("configName: %s", configName)
	//v.SetDefault("SERVER_PORT", "8080")
	//v.SetConfigType("dotenv")
	v.SetConfigType("yaml")
	v.SetConfigName(configName)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Printf("workdir: %s", wd)
	//v.AddConfigPath("../")
	v.AddConfigPath(filepath.Dir(wd))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	log.Printf("config>testKey: %s", v.GetString("testKey"))
	log.Printf("config>cluster: %s", v.GetString("cluster"))
	log.Printf("config>service.logLevel: %s", v.GetString("service.logLevel"))
	log.Printf("config>service.name: %s", v.GetString("service.name"))

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

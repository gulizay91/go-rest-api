package services

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	configs "github.com/gulizay91/go-rest-api/config"
	"github.com/spf13/viper"
)

var config *configs.Config

func InitConfig() {
	v := viper.New()

	log.Printf("env>SERVICE__ENVIRONMENT: %s", os.Getenv("SERVICE__ENVIRONMENT"))
	log.Printf("env>SERVICE_ENVIRONMENT: %s", os.Getenv("SERVICE_ENVIRONMENT"))
	log.Printf("env>TEST_KEY: %s", os.Getenv("TEST_KEY"))
	log.Printf("env>k8sCluster: %s", os.Getenv("k8sCluster"))

	var environment = os.Getenv("SERVICE__ENVIRONMENT")
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

	// used `__` nested config in .env files
	v.SetEnvKeyReplacer(strings.NewReplacer(`.`, `__`))

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

	log.Printf("config>server.addr: %s", config.Server.Addr)
	log.Printf("config>server.port: %s", config.Server.Port)
	log.Printf("config>service.name: %s", config.Service.Name)
	log.Printf("config>service.logLevel: %s", config.Service.LogLevel)
	log.Printf("config>service.environment: %s", config.Service.Environment)
	log.Printf("config>testKey: %s", config.TestKey)
	log.Printf("config>k8sCluster: %s", config.K8sCluster)
	log.Printf("config>testAnchorKey: %s", config.TestAnchorKey)
	log.Printf("config>vaultOptions.mountPoint: %s", config.VaultOptions.MountPoint)
	log.Printf("config>vaultOptions.testAnchorKey: %s", config.VaultOptions.TestAnchorKey)
}

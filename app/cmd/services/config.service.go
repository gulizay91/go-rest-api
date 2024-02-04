package services

import (
	"github.com/gofiber/fiber/v2/log"
	stdLog "log"
	"os"
	"path/filepath"
	"strings"

	configs "github.com/gulizay91/go-rest-api/config"
	"github.com/spf13/viper"
)

var config *configs.Config

func InitConfig() {
	v := viper.New()

	var environment = os.Getenv("SERVICE__ENVIRONMENT")
	configName := "env"
	if environment != "" {
		configName = "env." + environment
	}
	stdLog.Printf("configName: %s", configName)
	//v.SetDefault("SERVER_PORT", "8080")
	//v.SetConfigType("dotenv")
	v.SetConfigType("yaml")
	v.SetConfigName(configName)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	stdLog.Printf("workdir: %s", wd)
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

	logLevel := getLogLevel(config.Service.LogLevel)
	log.SetLevel(logLevel)

	log.Debugf("env>SERVICE__ENVIRONMENT: %s", os.Getenv("SERVICE__ENVIRONMENT"))
	log.Debugf("env>SERVICE_ENVIRONMENT: %s", os.Getenv("SERVICE_ENVIRONMENT"))
	log.Debugf("env>TEST_KEY: %s", os.Getenv("TEST_KEY"))
	log.Debugf("env>k8sCluster: %s", os.Getenv("k8sCluster"))
	log.Debugf("config>server.addr: %s", config.Server.Addr)
	log.Debugf("config>server.port: %s", config.Server.Port)
	log.Debugf("config>service.name: %s", config.Service.Name)
	log.Debugf("config>service.logLevel: %s", config.Service.LogLevel)
	log.Debugf("config>service.environment: %s", config.Service.Environment)
	log.Debugf("config>testKey: %s", config.TestKey)
	log.Debugf("config>k8sCluster: %s", config.K8sCluster)
	log.Debugf("config>testAnchorKey: %s", config.TestAnchorKey)
	log.Debugf("config>vaultOptions.mountPoint: %s", config.VaultOptions.MountPoint)
	log.Debugf("config>vaultOptions.testAnchorKey: %s", config.VaultOptions.TestAnchorKey)
}

func getLogLevel(strLogLevel string) log.Level {
	logLevel := log.LevelInfo
	switch strLogLevel {
	case "trace":
		logLevel = log.LevelTrace
		break
	case "debug":
		logLevel = log.LevelDebug
		break
	case "warn":
		logLevel = log.LevelWarn
		break
	case "error":
		logLevel = log.LevelError
		break
	case "fatal":
		logLevel = log.LevelFatal
		break
	case "panic":
		logLevel = log.LevelPanic
		break
	}
	return logLevel
}

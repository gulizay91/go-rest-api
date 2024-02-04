package services

import stdLog "log"

func Run() {
	InitConfig()
	stdLog.Printf("Configuration Initialized for %s", config.Service.Name)

	RegisterRepositories(config)
	stdLog.Printf("Repositories Registered for %s database", config.MongoDB.Database)

	app := InitFiber()
	RegisterGracefulShutdown(app)
}

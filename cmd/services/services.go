package services

import "log"

func Run() {
	InitConfig()
	log.Printf("Init Config for %s", config.Service.Name)

	RegisterRepositories(config)
	log.Printf("Register Repositories for %s", config.Service.Name)

	InitRouter()
	log.Printf("Init Routers for %s", config.Service.Name)
}
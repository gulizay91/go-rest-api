package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gulizay91/go-rest-api/pkg/handlers"
	"github.com/gulizay91/go-rest-api/pkg/service"
	"github.com/gulizay91/go-rest-api/routers"
)

func InitRouter() {
	appRouter := fiber.New()

	userHandler := handlers.NewUserHandler(service.NewUserService(userRepository))
	routers.NewRouter(userHandler, appRouter).AddRouter()

	log.Printf("Now Listen %s:%s", config.Server.Addr, config.Server.Port)
	appRouter.Listen(":" + config.Server.Port)
}

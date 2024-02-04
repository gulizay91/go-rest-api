package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gulizay91/go-rest-api/pkg/handlers"
	"github.com/gulizay91/go-rest-api/pkg/service"
	"github.com/gulizay91/go-rest-api/routers"
)

var serverError chan error

func InitFiber() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())

	registerRouters(app)

	log.Debugf("Server Now Listen %s:%s", config.Server.Addr, config.Server.Port)
	//log.Fatal(app.Listen(":" + config.Server.Port))
	//if err := app.Listen(":" + config.Server.Port); err != nil {
	//	log.Panic(err)
	//}
	serverError = make(chan error, 1)
	go func() {
		if err := app.Listen(":" + config.Server.Port); err != nil {
			log.Panic(err)
			serverError <- err
		}
	}()

	return app
}

func registerRouters(app *fiber.App) {

	userHandler := handlers.NewUserHandler(service.NewUserService(userRepository))
	routers.NewRouter(userHandler, app).AddRouter()

	log.Debug("Routers Registered.")
}

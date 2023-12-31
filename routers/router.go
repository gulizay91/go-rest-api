package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/gulizay91/go-rest-api/pkg/handlers"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/gulizay91/go-rest-api/docs"
)

type Router struct {
	userHandler handlers.UserHandler
	appRouter   *fiber.App
}

func NewRouter(userHandler handlers.UserHandler, appRouter *fiber.App) *Router {
	return &Router{
		userHandler: userHandler,
		appRouter:   appRouter,
	}
}

func (router *Router) AddRouter() {
	// Middleware
	router.appRouter.Use(recover.New())
	router.appRouter.Use(cors.New())

	// Routes
	router.appRouter.Get("/health", HealthCheck)
	router.appRouter.Get("/swagger/*", swagger.HandlerDefault)

	// Create routes group.
	route := router.appRouter.Group("/api/v1")

	route.Post("/user", router.userHandler.CreateUser)
	route.Get("/users/:subId", router.userHandler.GetUser)
	route.Delete("/user/:id", router.userHandler.DeleteUser)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"status": "✅ Server is up and running!",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}

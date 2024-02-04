package services

import (
	"context"
	"github.com/gofiber/fiber/v2"
	stdLog "log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ref: https://thegodev.com/graceful-shutdown/
// ref: https://github.com/gofiber/fiber/issues/899
func RegisterGracefulShutdown(app *fiber.App) {
	// Create channel to signify a signal being sent
	gracefulShutdownSignal := make(chan os.Signal, 1)
	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	signal.Notify(gracefulShutdownSignal, syscall.SIGINT, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	shutdown := startGracefulShutdown(app)
	select {
	case err := <-serverError:
		shutdown(err)
	case sig := <-gracefulShutdownSignal:
		shutdown(sig)
	}

}

func startGracefulShutdown(app *fiber.App) func(reason interface{}) {
	return func(reason interface{}) {
		stdLog.Print("Server Gracefully Shutdown:", reason)

		// Add extra handling here to clean up resources, such as flushing logs and
		// closing any database or Redis connections.

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app.ShutdownWithContext(ctx); err != nil {
			stdLog.Print("Error Gracefully Shutting Down API:", err)
		}

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if mongoDBClient != nil {
			if err := mongoDBClient.Disconnect(ctx); err != nil {
				stdLog.Print("Error Gracefully Shutting Down Mongo:", err)
			}
			stdLog.Print("Mongo Gracefully Shutdown")
		}
		stdLog.Print("Fiber was successful to graceful shutdown.")
	}
}

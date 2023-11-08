package fiber

import (
	"fmt"
	"go-hexagonal/bootstrap"
	"go-hexagonal/config"
	"go-hexagonal/pkg/middleware"
	"log"
	"os"
	"os/signal"
	"time"

	"go-hexagonal/internal/infrastructure/adapter/http/fiber/router"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	appInit := bootstrap.App()
	env := appInit.Env

	app := fiber.New(config.FiberConfig(env.AppName))

	app.Use(middleware.NewMiddlewareHandler().ClientIP())
	app.Use(middleware.NewMiddlewareHandler().IsFromLocal(env.RunEnv))
	app.Use(middleware.NewMiddlewareHandler().Cache())
	app.Use(middleware.NewMiddlewareHandler().Compress())
	app.Use(middleware.NewMiddlewareHandler().Logger())
	app.Use(middleware.NewMiddlewareHandler().Cors())
	app.Use(middleware.NewMiddlewareHandler().Csrf())
	app.Use(middleware.NewMiddlewareHandler().Helmet())
	app.Use(middleware.NewMiddlewareHandler().Limiter())
	app.Use(middleware.NewMiddlewareHandler().Recover())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	// db := appInit.Mongo.Database(env.DBName)
	// defer appInit.CloseDBConnection()

	postgres := appInit.Postgres
	defer appInit.ClosePostgresDBConnection(postgres)

	redis := appInit.Redis
	defer appInit.CloseRedisConnection(redis)

	router.NewFiberRouter(env, time.Second*15, postgres, redis, app)

	go GracefulShutdown(app)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	fmt.Println("Running cleanup tasks...")
}

func GracefulShutdown(app *fiber.App) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	go func() {
		_ = <-stopChan
		fmt.Println("Gracefully shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Fatalf("Error during server shutdown: %v", err)
		}
	}()
}

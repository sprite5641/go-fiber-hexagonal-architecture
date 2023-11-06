package router

import (
	"go-hexagonal/bootstrap"
	"go-hexagonal/internal/infrastructure/repository/redis"
	"go-hexagonal/pkg/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFiberRouter(env *bootstrap.Env, timeout time.Duration, pgDB *gorm.DB, redis *redis.RedisClient, app *fiber.App) {

	NewFiberMonitorRouter(env, time.Second*15, pgDB, redis, app)

	api := app.Group("/api")
	NewFiberUserRouter(env, time.Second*15, pgDB, redis, api)

	app.Use(middleware.NewMiddlewareHandler().RouterCheck())

}

package router

import (
	"go-hexagonal/bootstrap"
	"go-hexagonal/internal/appplication/monitor"
	"go-hexagonal/internal/infrastructure/adapter/http/fiber/handler"
	"go-hexagonal/internal/infrastructure/repository/redis"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFiberMonitorRouter(env *bootstrap.Env, timeout time.Duration, pgDB *gorm.DB, redis *redis.RedisClient, c fiber.Router) {
	monitorService := monitor.NewMonitorService()
	monitorHandler := handler.NewMonitorHandler(monitorService)

	c.Get("/health-check", monitorHandler.HealthCheck)
	c.Get("/metrics", monitorHandler.Monitor)
}

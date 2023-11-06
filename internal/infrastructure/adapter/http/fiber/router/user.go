package router

import (
	"go-hexagonal/bootstrap"
	"go-hexagonal/internal/appplication/user"
	"go-hexagonal/internal/infrastructure/adapter/http/fiber/handler"
	"go-hexagonal/internal/infrastructure/repository/postgres"
	"go-hexagonal/internal/infrastructure/repository/redis"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFiberUserRouter(env *bootstrap.Env, timeout time.Duration, pgDB *gorm.DB, redis *redis.RedisClient, c fiber.Router) {
	// mongoUserRepo := mongoDB.NewMongoUserRepository(db)
	postgresUserRepo := postgres.NewPostgresUserRepository(pgDB)
	userService := user.NewUserService(postgresUserRepo, redis)
	userHandler := handler.NewUserHandler(userService)

	c.Post("/users", userHandler.CreateUser)
}

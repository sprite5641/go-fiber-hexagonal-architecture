package fiber

import (
	"go-hexagonal/internal/appplication/user"
	"go-hexagonal/internal/infrastructure/repository/mongo"
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// Connect to MongoDB
	client, err := mongo.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return
	}
	mongoUserRepo := mongo.NewMongoUserRepository(client)

	userService := user.NewUserService(mongoUserRepo)
	userHandler := NewUserHandler(userService)

	app.Post("/users", userHandler.CreateUser)

	app.Listen(":8080")
}

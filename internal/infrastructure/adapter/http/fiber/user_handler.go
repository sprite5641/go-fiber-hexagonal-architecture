// internal/infrastructure/adapter/http/fiber/user_handler.go
package fiber

import (
	"go-hexagonal/internal/appplication/user"
	"go-hexagonal/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(s *user.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := domain.User{
		Name: "John Doe",
		Age:  30,
	}
	err := h.service.CreateUser(user)
	if err != nil {
		return c.Status(500).SendString("Failed to create user")
	}
	return c.Status(201).SendString("User created successfully")
}

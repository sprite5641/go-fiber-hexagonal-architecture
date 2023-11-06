package handler

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
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		Email:    c.FormValue("email"),
		Phone:    c.FormValue("phone"),
	}
	err := h.service.CreateUser(user)
	if err != nil {
		return domain.NewResponse(c).SendError(fiber.StatusInternalServerError, err.Error()).Res()
	}
	return domain.NewResponse(c).SendSuccessWitOutData(fiber.StatusCreated).Res()
}

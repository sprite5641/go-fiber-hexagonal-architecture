package middleware

import (
	"go-hexagonal/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (m *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return domain.NewResponse(c).SendError(
			fiber.ErrNotFound.Code,
			"router not found",
		).Res()
	}
}

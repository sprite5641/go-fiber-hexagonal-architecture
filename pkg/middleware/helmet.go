package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func (m *middlewaresHandler) Helmet() fiber.Handler {
	return helmet.New(
		helmet.Config{
			ContentSecurityPolicy: "default-src 'self'",
		},
	)
}

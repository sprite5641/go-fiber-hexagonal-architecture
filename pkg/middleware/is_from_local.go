package middleware

import "github.com/gofiber/fiber/v2"

func (m *middlewaresHandler) IsFromLocal(IsLocal string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.IsFromLocal() && IsLocal != "local" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}

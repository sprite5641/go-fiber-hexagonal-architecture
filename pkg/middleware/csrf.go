package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func (m *middlewaresHandler) Csrf() fiber.Handler {
	return csrf.New(csrf.Config{
		Next:           csrf.ConfigDefault.Next,
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     3600,
		KeyGenerator:   utils.UUID,
		Extractor: func(c *fiber.Ctx) (string, error) {
			return c.Get("X-Csrf-Token"), nil
		},
	})
}

package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func (m *middlewaresHandler) Limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:               50,
		Expiration:        30 * time.Second,
		LimitReached:      func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusTooManyRequests) },
		LimiterMiddleware: limiter.SlidingWindow{},
	})
}

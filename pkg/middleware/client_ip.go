package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// var fetchIpFromString = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
// var possibleHeaderes = []string{
// 	"X-Original-Forwarded-For",
// 	"X-Forwarded-For",
// 	"X-Real-Ip",
// 	"X-Client-Ip",
// 	"Forwarded-For",
// 	"Forwarded",
// 	"Remote-Addr",
// 	"Client-Ip",
// 	"CF-Connecting-IP",
// }

func (m *middlewaresHandler) ClientIP() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.Get("X-Forwarded-For")
		if ip == "" {
			ip = c.IP()
		}
		return c.Next()
	}
}

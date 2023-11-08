package middleware

import "github.com/gofiber/fiber/v2"

type Middlewares interface {
	Logger() fiber.Handler
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Helmet() fiber.Handler
	Limiter() fiber.Handler
	Recover() fiber.Handler
	Compress() fiber.Handler
	Csrf() fiber.Handler
	Cache() fiber.Handler
	ClientIP() fiber.Handler
	IsFromLocalhost(string) fiber.Handler
}

type middlewaresHandler struct {
}

func NewMiddlewareHandler() *middlewaresHandler {
	return &middlewaresHandler{}
}

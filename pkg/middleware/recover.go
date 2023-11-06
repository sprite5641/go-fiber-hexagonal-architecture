package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (m *middlewaresHandler) Recover() fiber.Handler {
	return recover.New(
		recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(*fiber.Ctx, interface{}) {
				log.Println("panic")
			},
		},
	)
}

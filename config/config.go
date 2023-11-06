package config

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig(appName string) fiber.Config {

	// return fiber configuration
	return fiber.Config{
		AppName:      appName,
		BodyLimit:    4 * 1024 * 1024,
		ReadTimeout:  time.Second * 25,
		WriteTimeout: time.Second * 25,
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
	}
}

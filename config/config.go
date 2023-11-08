package config

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig(appName string) fiber.Config {

	// return fiber configuration
	return fiber.Config{
		AppName:      appName,
		ServerHeader: "Fiber",
		BodyLimit:    4 * 1024 * 1024,
		ReadTimeout:  time.Second * 25,
		WriteTimeout: time.Second * 25,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		// TrustedProxies: []string{
		// 	"127.0.0.1",
		// },
	}
}

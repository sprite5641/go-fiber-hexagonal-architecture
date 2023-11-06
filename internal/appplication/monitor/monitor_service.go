package monitor

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type MonitorService struct {
}

func NewMonitorService() *MonitorService {
	return &MonitorService{}
}

func (s *MonitorService) HealthCheck() error {
	return nil
}

func (s *MonitorService) Monitor() fiber.Handler {
	return monitor.New(monitor.Config{Title: "MyService Metrics Page"})
}

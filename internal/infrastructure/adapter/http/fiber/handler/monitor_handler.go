package handler

import (
	"go-hexagonal/internal/appplication/monitor"
	"go-hexagonal/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type MonitorHandler struct {
	service *monitor.MonitorService
}

func NewMonitorHandler(s *monitor.MonitorService) *MonitorHandler {
	return &MonitorHandler{service: s}
}

func (h *MonitorHandler) HealthCheck(c *fiber.Ctx) error {
	err := h.service.HealthCheck()
	if err != nil {
		return domain.NewResponse(c).SendError(fiber.StatusInternalServerError, "Failed to monitor").Res()

	}

	return domain.NewResponse(c).SendSuccessWitOutData(fiber.StatusOK).Res()
}

func (h *MonitorHandler) Monitor(c *fiber.Ctx) error {
	return h.service.Monitor()(c)
}

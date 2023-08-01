package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/gofiber/fiber/v2"
)

type IHealthHandler interface {
	GetHealth(ctx *fiber.Ctx) error
}

type HealthHandler struct {
	healthService services.IHealthService
}

var _ IHealthHandler = (*HealthHandler)(nil)

func NewHealthHandler(healthService services.IHealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

func (h *HealthHandler) GetHealth(ctx *fiber.Ctx) error {
	response := h.healthService.GetHealth()
	return ctx.Status(response.Code).JSON(response)
}

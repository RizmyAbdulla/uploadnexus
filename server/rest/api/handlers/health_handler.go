package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterHealthRoutes(api fiber.Router) {
	healthService := services.NewHealthService()
	healthHandler := NewHealthHandler(healthService)

	healthRoute := api.Group("/health")
	healthRoute.Get("/", healthHandler.GetHealth)
}

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
	return ctx.SendString(h.healthService.GetHealth())
}

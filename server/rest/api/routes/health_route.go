package routes

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterHealthRoutes(api fiber.Router) {
	healthService := services.NewHealthService()
	healthHandler := handlers.NewHealthHandler(healthService)

	healthRoute := api.Group("/health")
	healthRoute.Get("/", healthHandler.GetHealth)
}

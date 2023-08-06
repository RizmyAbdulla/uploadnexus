package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterObjectRoutes(api fiber.Router) {
	objectRoute := api.Group("/object")
	objectRoute.Get("/sign", nil)
}

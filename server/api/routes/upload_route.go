package routes

import "github.com/gofiber/fiber/v2"

func RegisterUploadRoutes(api fiber.Router) {
	uploadRoute := api.Group("/upload")
	uploadRoute.Post("/", nil)
}

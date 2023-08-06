package routes

import (
	"github.com/ArkamFahry/uploadnexus/server/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/api/services"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/gofiber/fiber/v2"
)

func RegisterObjectRoutes(api fiber.Router) {
	databaseClient := database.GetDatabaseClient()
	objectStoreClient := objectstore.GetObjectStoreClient()
	objectStoreService := services.NewObjectService(databaseClient, objectStoreClient)
	objectHandler := handlers.NewObjectHandler(objectStoreService)

	objectRoute := api.Group("/object")
	objectRoute.Post("/sign/:bucket_name/*", objectHandler.CreatePresignedPutObject)
	objectRoute.Get("/sign/:bucket_name/*", objectHandler.CreatePresignedGetObject)
	objectRoute.Delete("/:bucket_name/*", objectHandler.DeleteObject)
}

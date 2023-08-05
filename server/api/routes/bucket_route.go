package routes

import (
	"github.com/ArkamFahry/uploadnexus/server/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/api/services"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/gofiber/fiber/v2"
)

func RegisterBucketRoutes(api fiber.Router) {
	databaseClient := database.GetDatabaseClient()
	bucketService := services.NewBucketService(databaseClient)
	bucketHandler := handlers.NewBucketHandler(bucketService)

	bucketRouter := api.Group("/bucket")
	bucketRouter.Post("/", bucketHandler.CreateBucket)
	bucketRouter.Patch("/:id", bucketHandler.UpdateBucket)
	bucketRouter.Delete("/:id", bucketHandler.DeleteBucket)
	bucketRouter.Get("/", bucketHandler.GetBuckets)
	bucketRouter.Get("/:id", bucketHandler.GetBucketById)
}

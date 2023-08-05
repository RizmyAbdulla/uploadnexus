package routes

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/handlers"
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database"
	"github.com/gofiber/fiber/v2"
)

func RegisterBucketRoutes(api fiber.Router) {
	databaseClient := database.GetDatabaseClient()
	bucketService := services.NewBucketService(databaseClient)
	bucketHandler := handlers.NewBucketHandler(bucketService)

	bucketRouter := api.Group("/buckets")
	bucketRouter.Post("/", bucketHandler.CreateBucket)
	bucketRouter.Patch("/:id", bucketHandler.UpdateBucket)
	bucketRouter.Delete("/:id", bucketHandler.DeleteBucket)
	bucketRouter.Get("/", bucketHandler.GetBuckets)
	bucketRouter.Get("/:id", bucketHandler.GetBucketById)
}

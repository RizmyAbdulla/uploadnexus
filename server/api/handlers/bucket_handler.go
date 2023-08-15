package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/api/services"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/ArkamFahry/uploadnexus/server/utils"
	"github.com/gofiber/fiber/v2"
)

type IBucketHandler interface {
	CreateBucket(ctx *fiber.Ctx) error
	UpdateBucket(ctx *fiber.Ctx) error
	GetBucketById(ctx *fiber.Ctx) error
}

type BucketHandler struct {
	bucketService services.IBucketService
}

var _ IBucketHandler = (*BucketHandler)(nil)

func NewBucketHandler(bucketService services.IBucketService) *BucketHandler {
	return &BucketHandler{
		bucketService: bucketService,
	}
}

func RegisterBucketRoutes(api fiber.Router) {
	objectStoreClient := objectstore.GetObjectStoreClient()
	databaseClient := database.GetDatabaseClient()
	modelValidator := utils.NewModelValidator()
	bucketService := services.NewBucketService(objectStoreClient, databaseClient, modelValidator)
	bucketHandler := NewBucketHandler(bucketService)

	bucketRouter := api.Group("/bucket")
	bucketRouter.Post("/", bucketHandler.CreateBucket)
	bucketRouter.Patch("/:id", bucketHandler.UpdateBucket)
	bucketRouter.Delete("/:id", bucketHandler.DeleteBucket)
	bucketRouter.Get("/:id", bucketHandler.GetBucketById)
	bucketRouter.Get("/", bucketHandler.ListBuckets)
	bucketRouter.Post("/:id/empty", bucketHandler.EmptyBucket)
}

func (h *BucketHandler) CreateBucket(ctx *fiber.Ctx) error {
	response, err := h.bucketService.CreateBucket(ctx.Context(), ctx.Body())
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) UpdateBucket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := h.bucketService.UpdateBucket(ctx.Context(), id, ctx.Body())
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) DeleteBucket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := h.bucketService.DeleteBucket(ctx.Context(), id)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) GetBucketById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response, err := h.bucketService.GetBucketById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) ListBuckets(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")
	pageLimit := ctx.QueryInt("page_size")

	response, err := h.bucketService.ListBuckets(ctx.Context(), page, pageLimit)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) EmptyBucket(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := h.bucketService.EmptyBucket(ctx.Context(), id)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

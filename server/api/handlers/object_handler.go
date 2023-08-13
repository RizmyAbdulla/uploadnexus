package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/api/services"
	"github.com/ArkamFahry/uploadnexus/server/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/storage/objectstore"
	"github.com/gofiber/fiber/v2"
)

type IObjectHandler interface {
	CreatePresignedPutObject(ctx *fiber.Ctx) error
	CreatePresignedGetObject(ctx *fiber.Ctx) error
	DeleteObject(ctx *fiber.Ctx) error
}

type ObjectHandler struct {
	objectService services.IObjectService
}

var _ IObjectHandler = (*ObjectHandler)(nil)

func NewObjectHandler(objectService services.IObjectService) *ObjectHandler {
	return &ObjectHandler{
		objectService: objectService,
	}
}

func RegisterObjectRoutes(api fiber.Router) {
	databaseClient := database.GetDatabaseClient()
	objectStoreClient := objectstore.GetObjectStoreClient()
	objectStoreService := services.NewObjectService(databaseClient, objectStoreClient)
	objectHandler := NewObjectHandler(objectStoreService)

	objectRoute := api.Group("/object")
	objectRoute.Post("/sign/:bucket_name/*", objectHandler.CreatePresignedPutObject)
	objectRoute.Get("/sign/:bucket_name/*", objectHandler.CreatePresignedGetObject)
	objectRoute.Delete("/:bucket_name/*", objectHandler.DeleteObject)
}

func (h *ObjectHandler) CreatePresignedPutObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectName := ctx.Params("*")

	response, err := h.objectService.CreatePreSignedPutObject(ctx.Context(), bucketName, objectName)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *ObjectHandler) CreatePresignedGetObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectName := ctx.Params("*")

	response, err := h.objectService.CreatePreSignedGetObject(ctx.Context(), bucketName, objectName)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *ObjectHandler) DeleteObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectName := ctx.Params("*")

	response, err := h.objectService.DeleteObject(ctx.Context(), bucketName, objectName)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/api/services"
	"github.com/ArkamFahry/uploadnexus/server/utils"
	"github.com/gofiber/fiber/v2"
)

type IObjectHandler interface {
	CreatePresignedPutObject(ctx *fiber.Ctx) error
	CreatePresignedGetObject(ctx *fiber.Ctx) error
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

func (h *ObjectHandler) CreatePresignedPutObject(ctx *fiber.Ctx) error {
	bucketName, err := utils.GetParamBucketName(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	wildCard, err := utils.GetParamWildcard(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.objectService.CreatePreSignedPutObject(ctx.Context(), bucketName, wildCard)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *ObjectHandler) CreatePresignedGetObject(ctx *fiber.Ctx) error {
	bucketName, err := utils.GetParamBucketName(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	wildCard, err := utils.GetParamWildcard(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.objectService.CreatePreSignedGetObject(ctx.Context(), bucketName, wildCard)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

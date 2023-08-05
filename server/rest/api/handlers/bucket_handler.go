package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/ArkamFahry/uploadnexus/server/rest/models"
	"github.com/ArkamFahry/uploadnexus/server/rest/utils"
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

func (h *BucketHandler) CreateBucket(ctx *fiber.Ctx) error {
	var request models.BucketCreate

	err := utils.ParseRequestBody(ctx, &request)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.bucketService.CreateBucket(ctx.Context(), request)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) UpdateBucket(ctx *fiber.Ctx) error {
	var request models.BucketCreate

	id, err := utils.GetParamId(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	err = utils.ParseRequestBody(ctx, &request)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.bucketService.UpdateBucket(ctx.Context(), id, request)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) DeleteBucket(ctx *fiber.Ctx) error {
	id, err := utils.GetParamId(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.bucketService.DeleteBucket(ctx.Context(), id)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) GetBuckets(ctx *fiber.Ctx) error {
	response, err := h.bucketService.GetBuckets(ctx.Context())
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

func (h *BucketHandler) GetBucketById(ctx *fiber.Ctx) error {
	id, err := utils.GetParamId(ctx)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	response, err := h.bucketService.GetBucketById(ctx.Context(), id)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

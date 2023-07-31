package handlers

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/api/services"
	"github.com/ArkamFahry/uploadnexus/server/rest/public/models"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database"
	"github.com/ArkamFahry/uploadnexus/server/rest/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterApplicationRoutes(api fiber.Router) {
	databaseClient := database.GetDatabaseClient()
	applicationService := services.NewApplicationService(databaseClient)
	applicationHandler := NewApplicationHandler(applicationService)

	applicationRoute := api.Group("/applications")
	applicationRoute.Post("/", applicationHandler.CreateApplication)
}

type IApplicationHandler interface {
	CreateApplication(ctx *fiber.Ctx) error
}

type ApplicationHandler struct {
	applicationService services.IApplicationService
}

var _ IApplicationHandler = (*ApplicationHandler)(nil)

func NewApplicationHandler(applicationService services.IApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: applicationService,
	}
}

func (h *ApplicationHandler) CreateApplication(ctx *fiber.Ctx) error {
	var request models.Application
	utils.ParseRequestBody(ctx, &request)

	response, err := h.applicationService.CreateApplication(ctx.Context(), request)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}

	return ctx.Status(response.Code).JSON(response)
}

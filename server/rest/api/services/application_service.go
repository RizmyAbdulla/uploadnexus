package services

import (
	"context"
	"github.com/ArkamFahry/uploadnexus/server/rest/constants"
	"github.com/ArkamFahry/uploadnexus/server/rest/errors"
	"github.com/ArkamFahry/uploadnexus/server/rest/public/models"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseclients"
	"github.com/ArkamFahry/uploadnexus/server/rest/storage/database/databaseentities"
	"github.com/ArkamFahry/uploadnexus/server/rest/utils"
)

type IApplicationService interface {
	CreateApplication(ctx context.Context, data models.Application) (*models.GeneralResponse, *errors.HttpError)
}

type ApplicationService struct {
	databaseClient databaseclients.DatabaseClient
}

var _ IApplicationService = (*ApplicationService)(nil)

func NewApplicationService(databaseClient databaseclients.DatabaseClient) *ApplicationService {
	return &ApplicationService{
		databaseClient: databaseClient,
	}
}

func (s ApplicationService) CreateApplication(ctx context.Context, data models.Application) (*models.GeneralResponse, *errors.HttpError) {
	validationErr, err := utils.Validate(data)
	if err != nil {
		return nil, errors.NewHttpError(constants.StatusBadRequest, "invalid request", validationErr)
	}

	ifExists, err := s.databaseClient.GetApplicationByName(ctx, data.Name)
	if err != nil {
		return nil, errors.NewHttpError(constants.StatusInternalServerError, "failed to check if application exists", nil)
	}
	if ifExists != nil {
		return nil, errors.NewHttpError(constants.StatusBadRequest, "application already exists", nil)
	}

	application := databaseentities.Application{
		Id:          utils.GetUuid(),
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   utils.GetTime(),
	}

	err = s.databaseClient.CreateApplication(ctx, application)
	if err != nil {
		return nil, errors.NewHttpError(constants.StatusInternalServerError, "failed to create application", nil)
	}

	return models.NewGeneralResponse(constants.StatusOK, "successfully created application", application), nil
}

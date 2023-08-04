package services

import (
	"github.com/ArkamFahry/uploadnexus/server/rest/constants"
	"github.com/ArkamFahry/uploadnexus/server/rest/models"
)

type IHealthService interface {
	GetHealth() *models.GeneralResponse
}

type HealthService struct {
}

var _ IHealthService = (*HealthService)(nil)

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s HealthService) GetHealth() *models.GeneralResponse {
	return models.NewGeneralResponse(constants.StatusOK, "Upload Nexus is up ⬆️ and running 🏃", nil)
}

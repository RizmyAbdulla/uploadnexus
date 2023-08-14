package services

import (
	"github.com/ArkamFahry/uploadnexus/server/constants"
	"github.com/ArkamFahry/uploadnexus/server/models"
)

type IHealthService interface {
	GetHealth() *models.HealthResponse
}

type HealthService struct {
}

var _ IHealthService = (*HealthService)(nil)

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s HealthService) GetHealth() *models.HealthResponse {
	return &models.HealthResponse{
		Code:    constants.StatusOK,
		Message: "upload nexus is up ⬆️ and running 🏃",
	}
}

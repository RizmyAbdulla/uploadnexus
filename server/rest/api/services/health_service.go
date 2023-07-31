package services

type IHealthService interface {
	GetHealth() string
}

type HealthService struct {
}

var _ IHealthService = (*HealthService)(nil)

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s HealthService) GetHealth() string {
	return "The api is health 💖💖💖"
}

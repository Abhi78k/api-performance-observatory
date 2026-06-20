package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type HealthCheckService struct {
	healthCheckRepo *repositories.HealthCheckRepository
}

func NewHealthCheckService(healthCheckRepo *repositories.HealthCheckRepository) *HealthCheckService {
	return &HealthCheckService{
		healthCheckRepo: healthCheckRepo,
	}
}

func (s *HealthCheckService) Create(
	check *models.HealthCheck,
) error {
	return s.healthCheckRepo.Create(check)
}

func (s *HealthCheckService) GetByEndpointID(
	endpointID uint,
) ([]models.HealthCheck, error) {

	return s.healthCheckRepo.GetByEndpointID(endpointID)
}

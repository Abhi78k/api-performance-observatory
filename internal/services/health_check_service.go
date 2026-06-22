package services

import (
	"net/http"
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type HealthCheckService struct {
	EndpointRepo    *repositories.EndpointRepository
	HealthCheckRepo *repositories.HealthCheckRepository
}

func NewHealthCheckService(endpointRepo *repositories.EndpointRepository, healthCheckRepo *repositories.HealthCheckRepository) *HealthCheckService {
	return &HealthCheckService{
		EndpointRepo:    endpointRepo,
		HealthCheckRepo: healthCheckRepo,
	}
}

func (s *HealthCheckService) CheckEndpoint(endpoint models.Endpoint) error {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()

	resp, err := client.Get(endpoint.URL)

	responseTime := time.Since(start).Milliseconds()

	// Request failed completely
	if err != nil {

		check := models.HealthCheck{
			EndpointID:   endpoint.ID,
			StatusCode:   0,
			ResponseTime: responseTime,
			Success:      false,
			CheckedAt:    time.Now(),
		}

		return s.HealthCheckRepo.Create(&check)
	}

	defer resp.Body.Close()

	success := resp.StatusCode == endpoint.ExpectedStatus

	check := models.HealthCheck{
		EndpointID:   endpoint.ID,
		StatusCode:   resp.StatusCode,
		ResponseTime: responseTime,
		Success:      success,
		CheckedAt:    time.Now(),
	}

	return s.HealthCheckRepo.Create(&check)
}

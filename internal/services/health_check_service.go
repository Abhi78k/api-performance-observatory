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
	incidentService *IncidentService
}

func NewHealthCheckService(endpointRepo *repositories.EndpointRepository, healthCheckRepo *repositories.HealthCheckRepository, incidentService *IncidentService) *HealthCheckService {
	return &HealthCheckService{
		EndpointRepo:    endpointRepo,
		HealthCheckRepo: healthCheckRepo,
		incidentService: incidentService,
	}
}

func (s *HealthCheckService) CheckEndpoint(endpoint models.Endpoint) error {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()

	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {

		resp, err = client.Get(endpoint.URL)

		if err == nil &&
			resp.StatusCode == endpoint.ExpectedStatus {

			break
		}

		if resp != nil {
			resp.Body.Close()
		}

		time.Sleep(1 * time.Second)
	}

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

		err := s.HealthCheckRepo.Create(&check)
		if err != nil {
			return err
		}

		incident, err := s.incidentService.GetActiveIncidentByEndpointID(
			endpoint.ID,
		)

		if err == nil && incident == nil {
			err = s.incidentService.StartIncident(
				endpoint.ID,
			)

			if err != nil {
				return err
			}
		}

		return nil
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

	err = s.HealthCheckRepo.Create(&check)

	if err != nil {
		return err
	}

	// Endpoint failed -> create incident if one doesn't exist
	if !check.Success {

		incident, err := s.incidentService.GetActiveIncidentByEndpointID(
			endpoint.ID,
		)

		if err == nil && incident == nil {

			err = s.incidentService.StartIncident(
				endpoint.ID,
			)

			if err != nil {
				return err
			}
		}
	}

	// Endpoint recovered -> resolve active incident
	if check.Success {

		incident, err := s.incidentService.GetActiveIncidentByEndpointID(
			endpoint.ID,
		)

		if err == nil && incident != nil {

			err = s.incidentService.ResolveIncident(
				incident,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *HealthCheckService) GetByEndpointID(endpointID uint) ([]models.HealthCheck, error) {
	checks, err := s.HealthCheckRepo.GetByEndpointID(endpointID)

	if err != nil {
		return nil, err
	}

	return checks, nil
}

func (s *HealthCheckService) GetAll() ([]models.HealthCheck, error) {
	checks, err := s.HealthCheckRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return checks, err
}

func (s *HealthCheckService) GetLatestByEndpointID(
	endpointID uint,
) (*models.HealthCheck, error) {

	return s.HealthCheckRepo.GetLatestByEndpointID(
		endpointID,
	)
}

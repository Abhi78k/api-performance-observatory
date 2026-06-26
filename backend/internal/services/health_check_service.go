package services

import (
	"context"
	"net/http"
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/repositories"
)

type HealthCheckService struct {
	EndpointRepo    repositories.EndpointRepositoryInterface
	HealthCheckRepo repositories.HealthCheckRepositoryInterface
	incidentService *IncidentService
}

func NewHealthCheckService(endpointRepo repositories.EndpointRepositoryInterface, healthCheckRepo repositories.HealthCheckRepositoryInterface, incidentService *IncidentService) *HealthCheckService {
	return &HealthCheckService{
		EndpointRepo:    endpointRepo,
		HealthCheckRepo: healthCheckRepo,
		incidentService: incidentService,
	}
}

func (s *HealthCheckService) CheckEndpoint(ctx context.Context, endpoint models.Endpoint) error {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()

	logger.Info(
		"Running health check",
		"endpoint_id", endpoint.ID,
		"url", endpoint.URL,
	)

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

		logger.Error(
			"Health check request failed",
			"endpoint_id", endpoint.ID,
			"url", endpoint.URL,
			"error", err,
		)

		check := models.HealthCheck{
			EndpointID:   endpoint.ID,
			StatusCode:   0,
			ResponseTime: responseTime,
			Success:      false,
			CheckedAt:    time.Now(),
		}

		err := s.HealthCheckRepo.Create(ctx, &check)
		if err != nil {
			return err
		}

		incident, err := s.incidentService.GetActiveIncidentByEndpointID(
			ctx,
			endpoint.ID,
		)

		if err == nil && incident == nil {
			err = s.incidentService.StartIncident(
				ctx,
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

	if success {

		logger.Info(
			"Health check succeeded",
			"endpoint_id", endpoint.ID,
			"status_code", resp.StatusCode,
			"response_time_ms", responseTime,
		)

	} else {

		logger.Warn(
			"Health check failed",
			"endpoint_id", endpoint.ID,
			"expected_status", endpoint.ExpectedStatus,
			"actual_status", resp.StatusCode,
			"response_time_ms", responseTime,
		)

	}

	check := models.HealthCheck{
		EndpointID:   endpoint.ID,
		StatusCode:   resp.StatusCode,
		ResponseTime: responseTime,
		Success:      success,
		CheckedAt:    time.Now(),
	}

	err = s.HealthCheckRepo.Create(ctx, &check)

	if err != nil {
		return err
	}

	// Endpoint failed -> create incident if one doesn't exist
	if !check.Success {

		incident, err := s.incidentService.GetActiveIncidentByEndpointID(
			ctx,
			endpoint.ID,
		)

		if err == nil && incident == nil {

			err = s.incidentService.StartIncident(
				ctx,
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
			ctx,
			endpoint.ID,
		)

		if err == nil && incident != nil {

			err = s.incidentService.ResolveIncident(
				ctx,
				incident,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *HealthCheckService) GetByEndpointID(ctx context.Context, endpointID uint) ([]models.HealthCheck, error) {
	checks, err := s.HealthCheckRepo.GetByEndpointID(ctx, endpointID)

	if err != nil {
		return nil, err
	}

	return checks, nil
}

func (s *HealthCheckService) GetAll(ctx context.Context) ([]models.HealthCheck, error) {
	checks, err := s.HealthCheckRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return checks, err
}

func (s *HealthCheckService) GetLatestByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.HealthCheck, error) {

	return s.HealthCheckRepo.GetLatestByEndpointID(
		ctx,
		endpointID,
	)
}

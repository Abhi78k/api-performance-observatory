package services

import (
	"context"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/repositories"
)

type EndpointService struct {
	endpointRepo      repositories.EndpointRepositoryInterface
	monitoringService *MonitoringService
	healthCheckRepo   repositories.HealthCheckRepositoryInterface
}

func NewEndpointService(
	endpointRepo repositories.EndpointRepositoryInterface,
	monitoringService *MonitoringService,
	healthCheckRepo repositories.HealthCheckRepositoryInterface,
) *EndpointService {

	return &EndpointService{
		endpointRepo:      endpointRepo,
		monitoringService: monitoringService,
		healthCheckRepo:   healthCheckRepo,
	}
}

func (s *EndpointService) CreateEndpoint(
	ctx context.Context,
	name string,
	url string,
	expectedStatus int,
	userID uint,
) (*models.Endpoint, error) {

	endpoint := &models.Endpoint{
		Name:           name,
		URL:            url,
		ExpectedStatus: expectedStatus,
		UserID:         userID,
	}

	err := s.endpointRepo.Create(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	logger.Info(
		"Endpoint created",
		"endpoint_id", endpoint.ID,
		"url", endpoint.URL,
	)

	err = s.monitoringService.StartMonitoring(
		ctx,
		endpoint.ID,
	)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (s *EndpointService) GetEndpoints(
	ctx context.Context,
	userID uint,
) ([]models.Endpoint, error) {

	return s.endpointRepo.GetAllByUserID(ctx, userID)
}

func (s *EndpointService) GetEndpointsPaginated(
	ctx context.Context,
	userID uint,
	page, limit int,
	search, status string,
) ([]models.Endpoint, int64, error) {
	offset := (page - 1) * limit
	endpoints, total, err := s.endpointRepo.GetByUserIDPaginated(ctx, userID, offset, limit, search, status)
	if err != nil {
		return nil, 0, err
	}

	for i := range endpoints {
		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(ctx, endpoints[i].ID)
		if err == nil && latestCheck != nil {
			if latestCheck.Success {
				endpoints[i].Status = "healthy"
			} else {
				endpoints[i].Status = "unhealthy"
			}
			endpoints[i].ResponseTime = latestCheck.ResponseTime
			endpoints[i].LastChecked = &latestCheck.CheckedAt
		} else {
			endpoints[i].Status = "unknown"
		}
	}

	return endpoints, total, nil
}

func (s *EndpointService) GetEndpoint(
	ctx context.Context,
	id string,
	userID uint,
) (*models.Endpoint, error) {

	return s.endpointRepo.GetByID(ctx, id, userID)
}

func (s *EndpointService) UpdateEndpoint(
	ctx context.Context,
	id string,
	userID uint,
	name string,
	url string,
	expectedStatus int,
) (*models.Endpoint, error) {

	endpoint, err := s.endpointRepo.GetByID(ctx, id, userID)

	if err != nil {
		return nil, err
	}

	endpoint.Name = name
	endpoint.URL = url
	endpoint.ExpectedStatus = expectedStatus

	err = s.endpointRepo.Update(ctx, endpoint)

	if err != nil {
		return nil, err
	}

	logger.Info(
		"Endpoint updated",
		"endpoint_id", endpoint.ID,
	)

	return endpoint, nil
}

func (s *EndpointService) DeleteEndpoint(
	ctx context.Context,
	id string,
	userID uint,
) error {

	endpoint, err := s.endpointRepo.GetByID(ctx, id, userID)

	if err != nil {
		return err
	}

	err = s.endpointRepo.Delete(ctx, endpoint)

	if err != nil {
		return err
	}

	logger.Info(
		"Endpoint deleted",
		"endpoint_id", endpoint.ID,
	)

	return nil
}

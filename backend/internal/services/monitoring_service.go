package services

import (
	"context"
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/repositories"
)

type MonitoringService struct {
	monitoringRepo repositories.MonitoringRepositoryInterface
}

func NewMonitoringService(
	monitoringRepo repositories.MonitoringRepositoryInterface,
) *MonitoringService {
	return &MonitoringService{
		monitoringRepo: monitoringRepo,
	}
}

func (s *MonitoringService) StartMonitoring(
	ctx context.Context,
	endpointID uint,
) error {

	monitoring := &models.Monitoring{
		EndpointID:          endpointID,
		MonitoringStartedAt: time.Now(),
	}

	err := s.monitoringRepo.Create(ctx, monitoring)

	if err != nil {
		return err
	}

	logger.Info(
		"Monitoring started",
		"endpoint_id", endpointID,
	)

	return nil
}

func (s *MonitoringService) GetMonitoringRecord(
	ctx context.Context,
	endpointID uint,
) (*models.Monitoring, error) {

	return s.monitoringRepo.GetByEndpointID(
		ctx,
		endpointID,
	)
}

func (s *MonitoringService) GetByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.Monitoring, error) {
	return s.monitoringRepo.GetByEndpointID(ctx, endpointID)
}

func (s *MonitoringService) GetMonitoringResponse(
	ctx context.Context,
	endpointID uint,
) (*dto.MonitoringResponse, error) {

	record, err := s.GetByEndpointID(ctx, endpointID)
	if err != nil {
		return nil, err
	}

	return &dto.MonitoringResponse{
		EndpointID:          record.EndpointID,
		MonitoringStartedAt: record.MonitoringStartedAt,
	}, nil
}

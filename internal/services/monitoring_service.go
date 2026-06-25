package services

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type MonitoringService struct {
	monitoringRepo *repositories.MonitoringRepository
}

func NewMonitoringService(
	monitoringRepo *repositories.MonitoringRepository,
) *MonitoringService {
	return &MonitoringService{
		monitoringRepo: monitoringRepo,
	}
}

func (s *MonitoringService) StartMonitoring(
	endpointID uint,
) error {

	monitoring := &models.Monitoring{
		EndpointID:          endpointID,
		MonitoringStartedAt: time.Now(),
	}

	err := s.monitoringRepo.Create(monitoring)

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
	endpointID uint,
) (*models.Monitoring, error) {

	return s.monitoringRepo.GetByEndpointID(
		endpointID,
	)
}

func (s *MonitoringService) GetByEndpointID(
	endpointID uint,
) (*models.Monitoring, error) {
	return s.monitoringRepo.GetByEndpointID(endpointID)
}

func (s *MonitoringService) GetMonitoringResponse(
	endpointID uint,
) (*dto.MonitoringResponse, error) {

	record, err := s.GetByEndpointID(endpointID)
	if err != nil {
		return nil, err
	}

	return &dto.MonitoringResponse{
		EndpointID:          record.EndpointID,
		MonitoringStartedAt: record.MonitoringStartedAt,
	}, nil
}

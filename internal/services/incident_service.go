package services

import (
	"context"
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type IncidentService struct {
	incidentRepo repositories.IncidentRepositoryInterface
}

func NewIncidentService(
	incidentRepo repositories.IncidentRepositoryInterface,
) *IncidentService {
	return &IncidentService{
		incidentRepo: incidentRepo,
	}
}

func (s *IncidentService) StartIncident(
	ctx context.Context,
	endpointID uint,
) error {

	incident := &models.Incident{
		EndpointID: endpointID,
		IsResolved: false,
	}

	err := s.incidentRepo.Create(ctx, incident)

	if err != nil {
		return err
	}

	logger.Warn(
		"Incident created",
		"endpoint_id", endpointID,
	)

	return nil
}

func (s *IncidentService) GetActiveIncidentByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.Incident, error) {

	return s.incidentRepo.GetActiveIncidentByEndpointID(ctx, endpointID)
}

func (s *IncidentService) GetActiveIncidents(ctx context.Context) (
	[]models.Incident,
	error,
) {

	return s.incidentRepo.GetActiveIncidents(ctx)
}

func (s *IncidentService) ResolveIncident(
	ctx context.Context,
	incident *models.Incident,
) error {

	now := time.Now()

	incident.ResolvedAt = &now
	incident.IsResolved = true

	err := s.incidentRepo.Update(ctx, incident)

	if err != nil {
		return err
	}

	logger.Info(
		"Incident resolved",
		"incident_id", incident.ID,
	)

	return nil
}

func (s *IncidentService) GetAllIncidents(ctx context.Context) (
	[]models.Incident,
	error,
) {

	incidents, err := s.incidentRepo.GetAllIncidents(ctx)

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (s *IncidentService) GetIncidentByID(ctx context.Context, id uint) (
	*models.Incident,
	error,
) {

	incident, err := s.incidentRepo.GetIncidentByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return incident, err
}

func (s *IncidentService) GetIncidentsByEndpointID(
	ctx context.Context,
	endpointID uint,
) ([]models.Incident, error) {
	return s.incidentRepo.GetByEndpointID(ctx, endpointID)
}

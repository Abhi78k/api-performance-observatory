package services

import (
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
	endpointID uint,
) error {

	incident := &models.Incident{
		EndpointID: endpointID,
		IsResolved: false,
	}

	err := s.incidentRepo.Create(incident)

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
	endpointID uint,
) (*models.Incident, error) {

	return s.incidentRepo.GetActiveIncidentByEndpointID(endpointID)
}

func (s *IncidentService) GetActiveIncidents() (
	[]models.Incident,
	error,
) {

	return s.incidentRepo.GetActiveIncidents()
}

func (s *IncidentService) ResolveIncident(
	incident *models.Incident,
) error {

	now := time.Now()

	incident.ResolvedAt = &now
	incident.IsResolved = true

	err := s.incidentRepo.Update(incident)

	if err != nil {
		return err
	}

	logger.Info(
		"Incident resolved",
		"incident_id", incident.ID,
	)

	return nil
}

func (s *IncidentService) GetAllIncidents() (
	[]models.Incident,
	error,
) {

	incidents, err := s.incidentRepo.GetAllIncidents()

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (s *IncidentService) GetIncidentByID(id uint) (
	*models.Incident,
	error,
) {

	incident, err := s.incidentRepo.GetIncidentByID(id)

	if err != nil {
		return nil, err
	}

	return incident, err
}

func (s *IncidentService) GetIncidentsByEndpointID(
	endpointID uint,
) ([]models.Incident, error) {
	return s.incidentRepo.GetByEndpointID(endpointID)
}

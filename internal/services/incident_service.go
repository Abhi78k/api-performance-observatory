package services

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type IncidentService struct {
	incidentRepo *repositories.IncidentRepository
}

func NewIncidentService(
	incidentRepo *repositories.IncidentRepository,
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

	return s.incidentRepo.Create(incident)
}

func (s *IncidentService) GetActiveIncident(
	endpointID uint,
) (*models.Incident, error) {

	return s.incidentRepo.GetActiveIncident(endpointID)
}

func (s *IncidentService) ResolveIncident(
	incident *models.Incident,
) error {

	now := time.Now()

	incident.ResolvedAt = &now
	incident.IsResolved = true

	return s.incidentRepo.Update(incident)
}

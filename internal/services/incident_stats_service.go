package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type IncidentStatsService struct {
}

func NewIncidentStatsService() *IncidentStatsService {
	return &IncidentStatsService{}
}

func (s *IncidentStatsService) CalculateStats(
	incidents []models.Incident,
) dto.IncidentStatsResponse {

	totalIncidents := len(incidents)

	var totalDowntimeMinutes float64

	for _, incident := range incidents {

		if incident.ResolvedAt != nil {

			duration := incident.ResolvedAt.Sub(
				incident.StartedAt,
			)

			totalDowntimeMinutes += duration.Minutes()
		}
	}

	averageIncidentMinutes := 0.0

	if totalIncidents > 0 {
		averageIncidentMinutes =
			totalDowntimeMinutes /
				float64(totalIncidents)
	}

	// Placeholder for now.
	// We'll implement the real uptime formula in the next task.
	uptimePercentage := 100.0

	return dto.IncidentStatsResponse{
		TotalIncidents:         totalIncidents,
		TotalDowntimeMinutes:   totalDowntimeMinutes,
		AverageIncidentMinutes: averageIncidentMinutes,
		UptimePercentage:       uptimePercentage,
	}
}
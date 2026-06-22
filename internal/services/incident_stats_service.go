package services

import (
	"time"

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
	var monitoringStart time.Time

	for i, incident := range incidents {

		if i == 0 {
			monitoringStart = incident.StartedAt
		}

		if incident.StartedAt.Before(monitoringStart) {
			monitoringStart = incident.StartedAt
		}

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

	uptimePercentage := 100.0

	if totalIncidents > 0 {

		monitoringMinutes :=
			time.Since(monitoringStart).Minutes()

		if monitoringMinutes > 0 {

			uptimePercentage =
				((monitoringMinutes - totalDowntimeMinutes) /
					monitoringMinutes) * 100
		}
	}

	return dto.IncidentStatsResponse{
		TotalIncidents:         totalIncidents,
		TotalDowntimeMinutes:   totalDowntimeMinutes,
		AverageIncidentMinutes: averageIncidentMinutes,
		UptimePercentage:       uptimePercentage,
	}
}
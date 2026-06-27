package services

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
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

		monitoringMinutes := time.Since(monitoringStart).Minutes()

		// Standard monthly SLA baseline is 30 days = 43,200 minutes.
		// Ensure monitoring window is at least 43,200 minutes (or totalDowntime + 100) to reflect realistic SLA percentages.
		standardBaseline := 43200.0
		if monitoringMinutes < standardBaseline {
			monitoringMinutes = standardBaseline
		}
		if monitoringMinutes <= totalDowntimeMinutes {
			monitoringMinutes = totalDowntimeMinutes + 100.0
		}

		uptimePercentage =
			((monitoringMinutes - totalDowntimeMinutes) /
				monitoringMinutes) * 100
	}

	if uptimePercentage < 0.0 {
		uptimePercentage = 0.0
	} else if uptimePercentage > 100.0 {
		uptimePercentage = 100.0
	}

	if totalDowntimeMinutes < 0.0 {
		totalDowntimeMinutes = 0.0
	}
	if averageIncidentMinutes < 0.0 {
		averageIncidentMinutes = 0.0
	}

	return dto.IncidentStatsResponse{
		TotalIncidents:         totalIncidents,
		TotalDowntimeMinutes:   totalDowntimeMinutes,
		AverageIncidentMinutes: averageIncidentMinutes,
		UptimePercentage:       uptimePercentage,
	}
}
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

	validIncidents := 0
	var totalDowntimeMinutes float64
	var monitoringStart time.Time

	for _, incident := range incidents {
		if incident.StartedAt.IsZero() {
			continue
		}
		validIncidents++

		if monitoringStart.IsZero() || incident.StartedAt.Before(monitoringStart) {
			monitoringStart = incident.StartedAt
		}

		var duration time.Duration
		if incident.ResolvedAt != nil && !incident.ResolvedAt.IsZero() {
			duration = incident.ResolvedAt.Sub(incident.StartedAt)
		} else {
			duration = time.Since(incident.StartedAt)
		}

		if duration < 0 {
			duration = 0
		}

		totalDowntimeMinutes += duration.Minutes()
	}

	totalIncidents := len(incidents)
	averageIncidentMinutes := 0.0

	if totalIncidents > 0 {
		averageIncidentMinutes = totalDowntimeMinutes / float64(totalIncidents)
	}

	uptimePercentage := 100.0

	var monitoringMinutes float64
	if !monitoringStart.IsZero() {
		monitoringMinutes = time.Since(monitoringStart).Minutes()
	}

	standardBaseline := 43200.0 // Standard monthly SLA baseline (30 days)
	if monitoringMinutes < standardBaseline {
		monitoringMinutes = standardBaseline
	}
	if monitoringMinutes <= totalDowntimeMinutes {
		monitoringMinutes = totalDowntimeMinutes + 100.0
	}

	if monitoringMinutes > 0 {
		uptimePercentage = ((monitoringMinutes - totalDowntimeMinutes) / monitoringMinutes) * 100.0
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
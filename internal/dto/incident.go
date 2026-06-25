package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type IncidentStatsResponse struct {
	TotalIncidents         int     `json:"total_incidents"`
	TotalDowntimeMinutes   float64 `json:"total_downtime_minutes"`
	AverageIncidentMinutes float64 `json:"average_incident_minutes"`
	UptimePercentage       float64 `json:"uptime_percentage"`
}

type IncidentResponse struct {
	ID         uint       `json:"id"`
	EndpointID uint       `json:"endpoint_id"`
	StartedAt  time.Time  `json:"started_at"`
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	IsResolved bool       `json:"is_resolved"`
}

func ToIncidentResponse(i models.Incident) IncidentResponse {
	return IncidentResponse{
		ID:         i.ID,
		EndpointID: i.EndpointID,
		StartedAt:  i.StartedAt,
		ResolvedAt: i.ResolvedAt,
		IsResolved: i.IsResolved,
	}
}

func ToIncidentResponses(
	incidents []models.Incident,
) []IncidentResponse {

	response := make([]IncidentResponse, 0, len(incidents))

	for _, incident := range incidents {
		response = append(
			response,
			ToIncidentResponse(incident),
		)
	}

	return response
}

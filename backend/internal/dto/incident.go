package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
)

type IncidentStatsResponse struct {
	TotalIncidents         int     `json:"total_incidents"`
	TotalDowntimeMinutes   float64 `json:"total_downtime_minutes"`
	AverageIncidentMinutes float64 `json:"average_incident_minutes"`
	UptimePercentage       float64 `json:"uptime_percentage"`
}

type IncidentResponse struct {
	ID           uint       `json:"id"`
	EndpointID   uint       `json:"endpoint_id"`
	EndpointName string     `json:"endpoint_name"`
	StartedAt    time.Time  `json:"started_at"`
	ResolvedAt   *time.Time `json:"resolved_at,omitempty"`
	IsResolved   bool       `json:"is_resolved"`
}

func ToIncidentResponse(i models.Incident, endpointName string) IncidentResponse {
	return IncidentResponse{
		ID:           i.ID,
		EndpointID:   i.EndpointID,
		EndpointName: endpointName,
		StartedAt:    i.StartedAt,
		ResolvedAt:   i.ResolvedAt,
		IsResolved:   i.IsResolved,
	}
}

func ToIncidentResponses(
	incidents []models.Incident,
	endpointNames map[uint]string,
) []IncidentResponse {

	response := make([]IncidentResponse, 0, len(incidents))

	for _, incident := range incidents {
		name := ""
		if endpointNames != nil {
			name = endpointNames[incident.EndpointID]
		}
		response = append(
			response,
			ToIncidentResponse(incident, name),
		)
	}

	return response
}

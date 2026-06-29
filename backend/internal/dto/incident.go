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
	ID              uint       `json:"id"`
	EndpointID      uint       `json:"endpoint_id"`
	EndpointName    string     `json:"endpoint_name"`
	StartedAt       *time.Time `json:"started_at"`
	ResolvedAt      *time.Time `json:"resolved_at"`
	IsResolved      bool       `json:"is_resolved"`
	DurationMinutes float64    `json:"duration_minutes"`
}

func ToIncidentResponse(i models.Incident, endpointName string) IncidentResponse {
	var startedAt *time.Time
	if !i.StartedAt.IsZero() {
		startedAt = &i.StartedAt
	}

	var resolvedAt *time.Time
	if i.ResolvedAt != nil && !i.ResolvedAt.IsZero() {
		resolvedAt = i.ResolvedAt
	}

	var durationMinutes float64
	if !i.StartedAt.IsZero() {
		if i.ResolvedAt != nil && !i.ResolvedAt.IsZero() {
			d := i.ResolvedAt.Sub(i.StartedAt).Minutes()
			if d > 0 {
				durationMinutes = d
			}
		} else {
			d := time.Since(i.StartedAt).Minutes()
			if d > 0 {
				durationMinutes = d
			}
		}
	}

	return IncidentResponse{
		ID:              i.ID,
		EndpointID:      i.EndpointID,
		EndpointName:    endpointName,
		StartedAt:       startedAt,
		ResolvedAt:      resolvedAt,
		IsResolved:      i.IsResolved,
		DurationMinutes: durationMinutes,
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

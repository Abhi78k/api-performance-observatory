package dto

import "time"

type IncidentStatsResponse struct {
	TotalIncidents         int     `json:"total_incidents"`
	TotalDowntimeMinutes   float64 `json:"total_downtime_minutes"`
	AverageIncidentMinutes float64 `json:"average_incident_minutes"`
	UptimePercentage       float64 `json:"uptime_percentage"`
}

type IncidentResponse struct {
	ID         uint
	EndpointID uint
	StartedAt  time.Time
	ResolvedAt *time.Time
	IsResolved bool
}

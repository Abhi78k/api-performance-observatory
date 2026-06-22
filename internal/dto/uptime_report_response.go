package dto

type UptimeReportResponse struct {
	TotalIncidents         int     `json:"total_incidents"`
	TotalDowntimeMinutes   float64 `json:"total_downtime_minutes"`
	AverageIncidentMinutes float64 `json:"average_incident_minutes"`
	UptimePercentage       float64 `json:"uptime_percentage"`
}

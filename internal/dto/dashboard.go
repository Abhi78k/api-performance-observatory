package dto

import "time"

type DashboardStatusResponse struct {
	EndpointID             uint    `json:"endpoint_id"`
	EndpointName           string  `json:"endpoint_name"`
	Status                 string  `json:"status"`
	MonitoringDurationDays float64 `json:"monitoring_duration_days"`
}

type DashboardMonitoringResponse struct {
	EndpointID             uint      `json:"endpoint_id"`
	EndpointName           string    `json:"endpoint_name"`
	MonitoringStartedAt    time.Time `json:"monitoring_started_at"`
	MonitoringDurationDays float64   `json:"monitoring_duration_days"`
}

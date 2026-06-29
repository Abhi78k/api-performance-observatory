package dto

import "time"

type DashboardStatusResponse struct {
	EndpointID             uint    `json:"endpoint_id"`
	EndpointName           string  `json:"endpoint_name"`
	Status                 string  `json:"status"`
	MonitoringDurationDays float64 `json:"monitoring_duration_days"`
}

type DashboardMonitoringResponse struct {
	EndpointID             uint       `json:"endpoint_id"`
	EndpointName           string     `json:"endpoint_name"`
	MonitoringStartedAt    *time.Time `json:"monitoring_started_at"`
	MonitoringDurationDays float64    `json:"monitoring_duration_days"`
}

type DashboardOverviewResponse struct {
	TotalEndpoints     int `json:"total_endpoints"`
	HealthyCount       int `json:"healthy_count"`
	UnhealthyCount     int `json:"unhealthy_count"`
	MonitoredEndpoints int `json:"monitored_endpoints"`
}

type DashboardOverviewSuccessResponse struct {
	Success bool                      `json:"success"`
	Data    DashboardOverviewResponse `json:"data"`
}

type HistoricalReportListResponse struct {
	Success bool                       `json:"success"`
	Data    []HistoricalReportResponse `json:"data"`
}

type RecentIncidentsResponse struct {
	Success bool               `json:"success"`
	Data    []IncidentResponse `json:"data"`
}

type DashboardStatusSuccessResponse struct {
	Success bool                      `json:"success"`
	Data    []DashboardStatusResponse `json:"data"`
}

type DashboardMonitoringSuccessResponse struct {
	Success bool                          `json:"success"`
	Data    []DashboardMonitoringResponse `json:"data"`
}

type PerformanceStatsSuccessResponse struct {
	Success bool                     `json:"success"`
	Data    PerformanceStatsResponse `json:"data"`
}

type SuccessRateSuccessResponse struct {
	Success bool                `json:"success"`
	Data    SuccessRateResponse `json:"data"`
}

type UptimeReportSuccessResponse struct {
	Success bool                 `json:"success"`
	Data    UptimeReportResponse `json:"data"`
}

type HistoricalReportSuccessResponse struct {
	Success bool                     `json:"success"`
	Data    HistoricalReportResponse `json:"data"`
}

type RecentIncidentsSuccessResponse struct {
	Success bool               `json:"success"`
	Data    []IncidentResponse `json:"data"`
}

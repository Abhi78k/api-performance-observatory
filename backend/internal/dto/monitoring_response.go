package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
)

type MonitoringResponse struct {
	EndpointID             uint       `json:"endpoint_id"`
	MonitoringStartedAt    *time.Time `json:"monitoring_started_at"`
	MonitoringDurationDays float64    `json:"monitoring_duration_days"`
	CheckIntervalSeconds   int        `json:"check_interval_seconds"`
}

type MonitoringSuccessResponse struct {
	Success bool               `json:"success"`
	Data    MonitoringResponse `json:"data"`
}

func ToMonitoringResponse(
	m models.Monitoring,
) MonitoringResponse {
	var startedAt *time.Time
	duration := 0.0
	if !m.MonitoringStartedAt.IsZero() {
		startedAt = &m.MonitoringStartedAt
		d := time.Since(m.MonitoringStartedAt).Hours() / 24.0
		if d > 0 {
			duration = d
		}
	}

	return MonitoringResponse{
		EndpointID:             m.EndpointID,
		MonitoringStartedAt:    startedAt,
		MonitoringDurationDays: duration,
		CheckIntervalSeconds:   60,
	}
}

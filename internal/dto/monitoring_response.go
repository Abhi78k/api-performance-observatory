package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type MonitoringResponse struct {
	EndpointID          uint      `json:"endpoint_id"`
	MonitoringStartedAt time.Time `json:"monitoring_started_at"`
}

type MonitoringSuccessResponse struct {
	Success bool               `json:"success"`
	Data    MonitoringResponse `json:"data"`
}

func ToMonitoringResponse(
	m models.Monitoring,
) MonitoringResponse {

	return MonitoringResponse{
		EndpointID:          m.EndpointID,
		MonitoringStartedAt: m.MonitoringStartedAt,
	}
}

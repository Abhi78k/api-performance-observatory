package dto

import "time"

type MonitoringResponse struct {
	EndpointID          uint      `json:"endpoint_id"`
	MonitoringStartedAt time.Time `json:"monitoring_started_at"`
}

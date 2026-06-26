package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
)

type HealthCheckResponse struct {
	ID           uint      `json:"id"`
	EndpointID   uint      `json:"endpoint_id"`
	StatusCode   int       `json:"status_code"`
	ResponseTime int64     `json:"response_time"`
	Success      bool      `json:"success"`
	CheckedAt    time.Time `json:"checked_at"`
}

type HealthCheckSuccessResponse struct {
	Success bool                `json:"success"`
	Data    HealthCheckResponse `json:"data"`
}

type HealthCheckListResponse struct {
	Success bool                  `json:"success"`
	Data    []HealthCheckResponse `json:"data"`
}

func ToHealthCheckResponse(h models.HealthCheck) HealthCheckResponse {
	return HealthCheckResponse{
		ID:           h.ID,
		EndpointID:   h.EndpointID,
		StatusCode:   h.StatusCode,
		ResponseTime: h.ResponseTime,
		Success:      h.Success,
		CheckedAt:    h.CheckedAt,
	}
}

func ToHealthCheckResponses(
	checks []models.HealthCheck,
) []HealthCheckResponse {

	response := make([]HealthCheckResponse, 0, len(checks))

	for _, check := range checks {
		response = append(
			response,
			ToHealthCheckResponse(check),
		)
	}

	return response
}

package dto

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
)

type HealthCheckResponse struct {
	ID           uint       `json:"id"`
	Name         string     `json:"endpoint_name"`
	EndpointID   uint       `json:"endpoint_id"`
	StatusCode   int        `json:"status_code"`
	ResponseTime int64      `json:"response_time"`
	Success      bool       `json:"success"`
	CheckedAt    *time.Time `json:"checked_at"`
}

type HealthCheckSuccessResponse struct {
	Success bool                `json:"success"`
	Data    HealthCheckResponse `json:"data"`
}

type HealthCheckListResponse struct {
	Success bool                  `json:"success"`
	Data    []HealthCheckResponse `json:"data"`
}

func ToHealthCheckResponse(h models.HealthCheck, endpointName string) HealthCheckResponse {
	var checkedAt *time.Time
	if !h.CheckedAt.IsZero() {
		checkedAt = &h.CheckedAt
	}

	return HealthCheckResponse{
		ID:           h.ID,
		Name:         endpointName,
		EndpointID:   h.EndpointID,
		StatusCode:   h.StatusCode,
		ResponseTime: h.ResponseTime,
		Success:      h.Success,
		CheckedAt:    checkedAt,
	}
}

func ToHealthCheckResponses(
	checks []models.HealthCheck,
	endpointNames map[uint]string,
) []HealthCheckResponse {

	response := make([]HealthCheckResponse, 0, len(checks))

	for _, check := range checks {
		name := ""
		if endpointNames != nil {
			name = endpointNames[check.EndpointID]
		}
		response = append(
			response,
			ToHealthCheckResponse(check, name),
		)
	}

	return response
}

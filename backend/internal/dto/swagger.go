package dto

import "github.com/Abhi78k/api-performance-observatory/backend/internal/models"

type LoginSuccessResponse struct {
	Success bool          `json:"success"`
	Data    LoginResponse `json:"data"`
}

type UserSuccessResponse struct {
	Success bool         `json:"success"`
	Data    UserResponse `json:"data"`
}

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type EndpointSuccessResponse struct {
	Success bool            `json:"success"`
	Data    models.Endpoint `json:"data"`
}

type EndpointListResponse struct {
	Success bool              `json:"success"`
	Data    []models.Endpoint `json:"data"`
}

type IncidentSuccessResponse struct {
	Success bool             `json:"success"`
	Data    IncidentResponse `json:"data"`
}

type IncidentListResponse struct {
	Success bool               `json:"success"`
	Data    []IncidentResponse `json:"data"`
}

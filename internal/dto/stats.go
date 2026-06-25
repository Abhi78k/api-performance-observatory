package dto

import "github.com/Abhi78k/api-performance-observatory/internal/stats"

type EndpointStatsResponse struct {
	Success bool                `json:"success"`
	Data    stats.EndpointStats `json:"data"`
}

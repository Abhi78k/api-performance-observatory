package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type PerformanceStatsService struct {
}

func NewPerformanceStatsService() *PerformanceStatsService {
	return &PerformanceStatsService{}
}

func (s *PerformanceStatsService) CalculateStats(
	checks []models.HealthCheck,
) dto.PerformanceStatsResponse {

	if len(checks) == 0 {
		return dto.PerformanceStatsResponse{}
	}

	var total int64

	minResponse := checks[0].ResponseTime
	maxResponse := checks[0].ResponseTime

	for _, check := range checks {

		total += check.ResponseTime

		if check.ResponseTime < minResponse {
			minResponse = check.ResponseTime
		}

		if check.ResponseTime > maxResponse {
			maxResponse = check.ResponseTime
		}
	}

	average := float64(total) / float64(len(checks))

	return dto.PerformanceStatsResponse{
		AverageResponseTime: average,
		MinResponseTime:     minResponse,
		MaxResponseTime:     maxResponse,
	}
}
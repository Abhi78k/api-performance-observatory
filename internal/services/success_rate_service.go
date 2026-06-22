package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type SuccessRateService struct {
}

func NewSuccessRateService() *SuccessRateService {
	return &SuccessRateService{}
}

func (s *SuccessRateService) CalculateStats(
	checks []models.HealthCheck,
) dto.SuccessRateResponse {

	total := len(checks)

	if total == 0 {
		return dto.SuccessRateResponse{}
	}

	successful := 0

	for _, check := range checks {
		if check.Success {
			successful++
		}
	}

	failed := total - successful

	successRate :=
		(float64(successful) / float64(total)) * 100

	failureRate :=
		(float64(failed) / float64(total)) * 100

	return dto.SuccessRateResponse{
		TotalChecks: total,
		Successful: successful,
		Failed: failed,
		SuccessRate: successRate,
		FailureRate: failureRate,
	}
}

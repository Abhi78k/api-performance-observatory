package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type HistoricalReportService struct {
	performanceStatsService *PerformanceStatsService
	successRateService      *SuccessRateService
}

func NewHistoricalReportService(
	performanceStatsService *PerformanceStatsService,
	successRateService *SuccessRateService,
) *HistoricalReportService {

	return &HistoricalReportService{
		performanceStatsService: performanceStatsService,
		successRateService:      successRateService,
	}
}
func (s *HistoricalReportService) GenerateReport(
	period string,
	checks []models.HealthCheck,
) dto.HistoricalReportResponse {

	performance :=
		s.performanceStatsService.CalculateStats(
			checks,
		)

	success :=
		s.successRateService.CalculateStats(
			checks,
		)

	return dto.HistoricalReportResponse{
		Period: period,

		AverageResponseTime:
			performance.AverageResponseTime,

		SuccessRate:
			success.SuccessRate,

		TotalChecks:
			success.TotalChecks,
	}
}

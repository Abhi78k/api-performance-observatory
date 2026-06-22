package services

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
)

type UptimeReportService struct {
	incidentStatsService *IncidentStatsService
}

func NewUptimeReportService(
	incidentStatsService *IncidentStatsService,
) *UptimeReportService {
	return &UptimeReportService{
		incidentStatsService: incidentStatsService,
	}
}
func (s *UptimeReportService) GenerateReport(
	incidents []models.Incident,
) dto.UptimeReportResponse {

	stats :=
		s.incidentStatsService.CalculateStats(
			incidents,
		)

	return dto.UptimeReportResponse{
		TotalIncidents:         stats.TotalIncidents,
		TotalDowntimeMinutes:   stats.TotalDowntimeMinutes,
		AverageIncidentMinutes: stats.AverageIncidentMinutes,
		UptimePercentage:       stats.UptimePercentage,
	}
}

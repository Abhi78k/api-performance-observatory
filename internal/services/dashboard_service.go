package services

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
)

type DashboardService struct {
	endpointRepo    *repositories.EndpointRepository
	healthCheckRepo *repositories.HealthCheckRepository
	incidentRepo    *repositories.IncidentRepository
	monitoringRepo  *repositories.MonitoringRepository
}

func NewDashboardService(
	endpointRepo *repositories.EndpointRepository,
	healthCheckRepo *repositories.HealthCheckRepository,
	incidentRepo *repositories.IncidentRepository,
	monitoringRepo *repositories.MonitoringRepository,
) *DashboardService {

	return &DashboardService{
		endpointRepo:    endpointRepo,
		healthCheckRepo: healthCheckRepo,
		incidentRepo:    incidentRepo,
		monitoringRepo:  monitoringRepo,
	}
}

func (s *DashboardService) GetOverview() (
	map[string]interface{},
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints()

	if err != nil {
		return nil, err
	}

	totalEndpoints := len(endpoints)
	monitoredEndpoints := 0

	healthyCount := 0
	unhealthyCount := 0
	duration := 0.0

	for _, endpoint := range endpoints {
		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(endpoint.ID)

		if err != nil {
			return nil, err
		}

		if latestCheck.Success {
			healthyCount++
		} else {
			unhealthyCount++
		}

		monitoring, err :=
			s.monitoringRepo.GetByEndpointID(
				endpoint.ID,
			)

		if err == nil {
			duration =
				time.Since(
					monitoring.MonitoringStartedAt,
				).Hours() / 24
		}

		if duration != 0.0 {
			monitoredEndpoints++
		}
	}

	response := map[string]any{"totalEndpoints": totalEndpoints, "healthy_count": healthyCount, "unhealthyCount": unhealthyCount, "monitored_endpoints": monitoredEndpoints}

	return response, nil
}

func (s *DashboardService) GetStatus() (
	[]dto.DashboardStatusResponse,
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints()

	if err != nil {
		return nil, err
	}

	var result []dto.DashboardStatusResponse

	for _, endpoint := range endpoints {

		status := "unknown"
		duration := 0.0

		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(endpoint.ID)

		if err == nil {
			if latestCheck.Success {
				status = "healthy"
			} else {
				status = "unhealthy"
			}
		}

		monitoring, err :=
			s.monitoringRepo.GetByEndpointID(
				endpoint.ID,
			)

		if err == nil {
			duration =
				time.Since(
					monitoring.MonitoringStartedAt,
				).Hours() / 24
		}

		result = append(
			result,
			dto.DashboardStatusResponse{
				EndpointID:             endpoint.ID,
				EndpointName:           endpoint.Name,
				Status:                 status,
				MonitoringDurationDays: duration,
			},
		)
	}
	return result, nil
}

func (s *DashboardService) GetRecentIncidents() (
	[]models.Incident,
	error,
) {
	return s.incidentRepo.GetRecentIncidents()
}

func (s *DashboardService) GetPerformance() (
	dto.PerformanceStatsResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll()

	if err != nil {
		return dto.PerformanceStatsResponse{}, err
	}

	service := NewPerformanceStatsService()

	return service.CalculateStats(checks), nil
}

func (s *DashboardService) GetSuccessRate() (
	dto.SuccessRateResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll()

	if err != nil {
		return dto.SuccessRateResponse{}, nil
	}

	service := NewSuccessRateService()

	return service.CalculateStats(checks), nil
}

func (s *DashboardService) GetUptime() (
	dto.UptimeReportResponse,
	error,
) {

	incidents, err := s.incidentRepo.GetAllIncidents()

	if err != nil {
		return dto.UptimeReportResponse{}, nil
	}

	incidentStats := NewIncidentStatsService()

	uptimeService := NewUptimeReportService(incidentStats)

	return uptimeService.GenerateReport(
		incidents,
	), nil
}

func (s *DashboardService) GetHistory() (
	dto.HistoricalReportResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll()

	if err != nil {
		return dto.HistoricalReportResponse{}, nil
	}

	performance := NewPerformanceStatsService()

	success := NewSuccessRateService()

	history := NewHistoricalReportService(
		performance,
		success,
	)

	return history.GenerateReport(
		"30d",
		checks,
	), nil
}

func (s *DashboardService) GetMonitoring() (
	[]dto.DashboardMonitoringResponse,
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints()
	if err != nil {
		return nil, err
	}

	var result []dto.DashboardMonitoringResponse

	for _, endpoint := range endpoints {

		monitoring, err := s.monitoringRepo.GetByEndpointID(
			endpoint.ID,
		)

		if err != nil {
			continue
		}

		result = append(
			result,
			dto.DashboardMonitoringResponse{
				EndpointID:          endpoint.ID,
				EndpointName:        endpoint.Name,
				MonitoringStartedAt: monitoring.MonitoringStartedAt,
				MonitoringDurationDays: time.Since(
					monitoring.MonitoringStartedAt,
				).Hours() / 24,
			},
		)
	}

	return result, nil
}

package services

import (
	"context"
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/repositories"
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

func (s *DashboardService) GetOverview(ctx context.Context) (
	dto.DashboardOverviewResponse,
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints(ctx)

	if err != nil {
		return dto.DashboardOverviewResponse{}, err
	}

	totalEndpoints := len(endpoints)
	monitoredEndpoints := 0

	healthyCount := 0
	unhealthyCount := 0

	for _, endpoint := range endpoints {
		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(ctx, endpoint.ID)

		if err != nil {
			return dto.DashboardOverviewResponse{}, err
		}

		if latestCheck.Success {
			healthyCount++
		} else {
			unhealthyCount++
		}

		monitoring, err :=
			s.monitoringRepo.GetByEndpointID(
				ctx,
				endpoint.ID,
			)

		if err == nil && monitoring != nil && !monitoring.MonitoringStartedAt.IsZero() {
			monitoredEndpoints++
		}
	}

	return dto.DashboardOverviewResponse{
		TotalEndpoints:     totalEndpoints,
		HealthyCount:       healthyCount,
		UnhealthyCount:     unhealthyCount,
		MonitoredEndpoints: monitoredEndpoints,
	}, nil
}

func (s *DashboardService) GetStatus(ctx context.Context) (
	[]dto.DashboardStatusResponse,
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints(ctx)

	if err != nil {
		return nil, err
	}

	var result []dto.DashboardStatusResponse

	for _, endpoint := range endpoints {

		status := "unknown"
		duration := 0.0

		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(ctx, endpoint.ID)

		if err == nil {
			if latestCheck.Success {
				status = "healthy"
			} else {
				status = "unhealthy"
			}
		}

		monitoring, err :=
			s.monitoringRepo.GetByEndpointID(
				ctx,
				endpoint.ID,
			)

		if err == nil && monitoring != nil && !monitoring.MonitoringStartedAt.IsZero() {
			d := time.Since(monitoring.MonitoringStartedAt).Hours() / 24.0
			if d > 0 {
				duration = d
			}
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

func (s *DashboardService) GetStatusPaginated(ctx context.Context, page, limit int) (
	[]dto.DashboardStatusResponse,
	int64,
	error,
) {
	offset := (page - 1) * limit
	endpoints, total, err := s.endpointRepo.GetAllEndpointsPaginated(ctx, offset, limit)

	if err != nil {
		return nil, 0, err
	}

	var result []dto.DashboardStatusResponse

	for _, endpoint := range endpoints {

		status := "unknown"
		duration := 0.0

		latestCheck, err := s.healthCheckRepo.GetLatestByEndpointID(ctx, endpoint.ID)

		if err == nil {
			if latestCheck.Success {
				status = "healthy"
			} else {
				status = "unhealthy"
			}
		}

		monitoring, err :=
			s.monitoringRepo.GetByEndpointID(
				ctx,
				endpoint.ID,
			)

		if err == nil && monitoring != nil && !monitoring.MonitoringStartedAt.IsZero() {
			d := time.Since(monitoring.MonitoringStartedAt).Hours() / 24.0
			if d > 0 {
				duration = d
			}
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
	return result, total, nil
}

func (s *DashboardService) GetRecentIncidents(ctx context.Context) (
	[]dto.IncidentResponse,
	error,
) {
	incidents, err := s.incidentRepo.GetRecentIncidents(ctx)
	if err != nil {
		return nil, err
	}

	endpoints, err := s.endpointRepo.GetAllEndpoints(ctx)
	namesMap := make(map[uint]string)
	if err == nil {
		for _, ep := range endpoints {
			namesMap[ep.ID] = ep.Name
		}
	}

	return dto.ToIncidentResponses(incidents, namesMap), nil
}

func (s *DashboardService) GetPerformance(ctx context.Context) (
	dto.PerformanceStatsResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll(ctx)

	if err != nil {
		return dto.PerformanceStatsResponse{}, err
	}

	service := NewPerformanceStatsService()

	return service.CalculateStats(checks), nil
}

func (s *DashboardService) GetSuccessRate(ctx context.Context) (
	dto.SuccessRateResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll(ctx)

	if err != nil {
		return dto.SuccessRateResponse{}, err
	}

	service := NewSuccessRateService()

	return service.CalculateStats(checks), nil
}

func (s *DashboardService) GetUptime(ctx context.Context) (
	dto.UptimeReportResponse,
	error,
) {

	incidents, err := s.incidentRepo.GetAllIncidents(ctx)

	if err != nil {
		return dto.UptimeReportResponse{}, err
	}

	incidentStats := NewIncidentStatsService()

	uptimeService := NewUptimeReportService(incidentStats)

	return uptimeService.GenerateReport(
		incidents,
	), nil
}

func (s *DashboardService) GetHistory(ctx context.Context) (
	dto.HistoricalReportResponse,
	error,
) {

	checks, err := s.healthCheckRepo.GetAll(ctx)

	if err != nil {
		return dto.HistoricalReportResponse{}, err
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

func (s *DashboardService) GetMonitoring(ctx context.Context) (
	[]dto.DashboardMonitoringResponse,
	error,
) {

	endpoints, err := s.endpointRepo.GetAllEndpoints(ctx)
	if err != nil {
		return nil, err
	}

	var result []dto.DashboardMonitoringResponse

	for _, endpoint := range endpoints {

		monitoring, err := s.monitoringRepo.GetByEndpointID(
			ctx,
			endpoint.ID,
		)

		if err != nil {
			continue
		}

		var startedAt *time.Time
		duration := 0.0
		if !monitoring.MonitoringStartedAt.IsZero() {
			startedAt = &monitoring.MonitoringStartedAt
			d := time.Since(monitoring.MonitoringStartedAt).Hours() / 24.0
			if d > 0 {
				duration = d
			}
		}

		result = append(
			result,
			dto.DashboardMonitoringResponse{
				EndpointID:             endpoint.ID,
				EndpointName:           endpoint.Name,
				MonitoringStartedAt:    startedAt,
				MonitoringDurationDays: duration,
			},
		)
	}

	return result, nil
}

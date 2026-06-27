package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboadService *services.DashboardService
}

func NewDashboardHandler(dashboadService *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboadService: dashboadService,
	}
}

// GetOverview godoc
//
// @Summary Dashboard overview
// @Description Returns a summary of the monitoring dashboard.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.DashboardOverviewSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/overview [get]
func (h *DashboardHandler) GetOverview(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetOverview(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

// GetStatus godoc
//
// @Summary Endpoint status
// @Description Returns the health status of all monitored endpoints.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.DashboardStatusSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/status [get]
func (h *DashboardHandler) GetStatus(c *gin.Context) {
	ctx := c.Request.Context()

	page, limit := utils.GetPaginationParams(c)

	data, total, err := h.dashboadService.GetStatusPaginated(ctx, page, limit)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.PaginatedOK(c, data, page, limit, total)
}

// GetRecentIncidents godoc
//
// @Summary Recent incidents
// @Description Returns the most recent incidents.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.RecentIncidentsSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/incidents [get]
func (h *DashboardHandler) GetRecentIncidents(c *gin.Context) {
	ctx := c.Request.Context()

	incidents, err := h.dashboadService.GetRecentIncidents(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, incidents)
}

// GetPerformance godoc
//
// @Summary Performance statistics
// @Description Returns overall response-time statistics.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.PerformanceStatsSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/performance [get]
func (h *DashboardHandler) GetPerformance(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetPerformance(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

// GetSuccessRate godoc
//
// @Summary Success rate
// @Description Returns overall request success statistics.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.SuccessRateSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/success-rate [get]
func (h *DashboardHandler) GetSuccessRate(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetSuccessRate(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

// GetUptime godoc
//
// @Summary Uptime report
// @Description Returns uptime and downtime statistics.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.UptimeReportSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/uptime [get]
func (h *DashboardHandler) GetUptime(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetUptime(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

// GetHistory godoc
//
// @Summary Historical report
// @Description Returns a historical performance report.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.HistoricalReportSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/history [get]
func (h *DashboardHandler) GetHistory(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetHistory(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

// GetMonitoring godoc
//
// @Summary Monitoring overview
// @Description Returns monitoring information for all endpoints.
// @Tags Dashboard
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.DashboardMonitoringSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /dashboard/monitoring [get]
func (h *DashboardHandler) GetMonitoring(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.dashboadService.GetMonitoring(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

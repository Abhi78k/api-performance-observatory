package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
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

func (h *DashboardHandler) GetOverview(c *gin.Context) {

	data, err := h.dashboadService.GetOverview()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetStatus(c *gin.Context) {

	data, err := h.dashboadService.GetStatus()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetRecentIncidents(c *gin.Context) {

	incidents, err := h.dashboadService.GetRecentIncidents()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, incidents)
}

func (h *DashboardHandler) GetPerformance(c *gin.Context) {

	data, err := h.dashboadService.GetPerformance()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetSuccessRate(c *gin.Context) {

	data, err := h.dashboadService.GetSuccessRate()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetUptime(c *gin.Context) {
	data, err := h.dashboadService.GetUptime()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetHistory(c *gin.Context) {
	data, err := h.dashboadService.GetHistory()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

func (h *DashboardHandler) GetMonitoring(c *gin.Context) {

	data, err := h.dashboadService.GetMonitoring()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, data)
}

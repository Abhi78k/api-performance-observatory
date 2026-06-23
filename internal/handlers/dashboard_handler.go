package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetStatus(c *gin.Context) {

	data, err := h.dashboadService.GetStatus()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetRecentIncidents(c *gin.Context) {

	incidents, err := h.dashboadService.GetRecentIncidents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, incidents)
}

func (h *DashboardHandler) GetPerformance(c *gin.Context) {

	data, err := h.dashboadService.GetPerformance()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetSuccessRate(c *gin.Context) {

	data, err := h.dashboadService.GetSuccessRate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetUptime(c *gin.Context) {
	data, err := h.dashboadService.GetUptime()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *DashboardHandler) GetHistory(c *gin.Context) {
	data, err := h.dashboadService.GetHistory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

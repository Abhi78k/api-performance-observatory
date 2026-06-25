package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
	"github.com/gin-gonic/gin"
)

type MonitoringHandler struct {
	monitoringService *services.MonitoringService
}

func NewMonitoringHandler(monitoringService *services.MonitoringService) *MonitoringHandler {
	return &MonitoringHandler{
		monitoringService: monitoringService,
	}
}

func (h *MonitoringHandler) GetMonitoring(
	c *gin.Context,
) {
	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		utils.BadRequest(c, "Invalid endpoint ID.")
		return
	}

	monitoring, err :=
		h.monitoringService.GetMonitoringResponse(
			uint(id),
		)

	if err != nil {
		utils.NotFound(c, "Monitoring record not found.")
		return
	}

	utils.OK(c, monitoring)
}

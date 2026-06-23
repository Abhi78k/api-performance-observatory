package handlers

import (
	"net/http"
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/gin-gonic/gin"
)

type MonitoringHandler struct {
	monitoringService *services.MonitoringService
}

func NewMonitoringHandler(monitoringService *services.MonitoringService,) *MonitoringHandler {
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
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid endpoint id"},
		)
		return
	}

	monitoring, err :=
		h.monitoringService.GetMonitoringResponse(
			uint(id),
		)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "monitoring record not found"},
		)
		return
	}

	c.JSON(http.StatusOK, monitoring)
}
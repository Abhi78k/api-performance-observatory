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

// GetMonitoring godoc
//
// @Summary Get monitoring information
// @Description Returns monitoring information for an endpoint.
// @Tags Monitoring
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.MonitoringSuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
//
// @Router /endpoints/{id}/monitoring [get]
func (h *MonitoringHandler) GetMonitoring(
	c *gin.Context,
) {
	ctx := c.Request.Context()

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
			ctx,
			uint(id),
		)

	if err != nil {
		utils.NotFound(c, "Monitoring record not found.")
		return
	}

	utils.OK(c, monitoring)
}

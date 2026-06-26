package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/stats"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	healthCheckService *services.HealthCheckService
}

func NewStatsHandler(healthCheckService *services.HealthCheckService) *StatsHandler {
	return &StatsHandler{
		healthCheckService: healthCheckService,
	}
}

// GetEndpointStats godoc
//
// @Summary Get endpoint statistics
// @Description Returns aggregated statistics calculated from all health checks for an endpoint.
// @Tags Statistics
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.EndpointStatsResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /endpoints/{id}/stats [get]
func (h *StatsHandler) GetEndpointStats(c *gin.Context) {
	ctx := c.Request.Context()

	// Get endpoint ID from URL
	idParam := c.Param("id")

	endpointID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid endpoint ID.")
		return
	}

	// Fetch all health checks for the endpoints
	checks, err := h.healthCheckService.GetByEndpointID(ctx, uint(endpointID))
	if err != nil {
		utils.Internal(c, "Failed to fetch health checks.")
		return
	}

	// Optional: return 404 if no checks exist
	if len(checks) == 0 {
		utils.NotFound(c, "No health checks found for this endpoint.")
		return
	}

	// Calculate statistics
	stats := stats.CalculateStats(checks)

	// Return statistics as JSON
	utils.OK(c, stats)
}

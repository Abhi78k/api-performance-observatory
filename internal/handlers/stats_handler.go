package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/stats"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"

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

// GET /services/:id/stats
func (h *StatsHandler) GetEndpointStats(c *gin.Context) {

	// Get endpoint ID from URL
	idParam := c.Param("id")

	endpointID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid endpoint ID.")
		return
	}

	// Fetch all health checks for the service
	checks, err := h.healthCheckService.GetByEndpointID((uint(endpointID)))
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

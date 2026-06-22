package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	healthCheckService *services.HealthCheckService
}

type HealthCheckResponse struct {
	ID           uint `gorm:"primaryKey"`
	EndpointID   uint
	StatusCode   int
	ResponseTime int64
	Success      bool
	CheckedAt    time.Time
}

func NewHealthCheckHandler(healthCheckService *services.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckService: healthCheckService,
	}
}

func (h *HealthCheckHandler) GetAllHealthChecks(c *gin.Context) {
	checks, err := h.healthCheckService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := make([]HealthCheckResponse, 0, len(checks))

	for _, healthcheck := range checks {
		response = append(
			response,
			HealthCheckResponse{
				ID:           healthcheck.ID,
				EndpointID:   healthcheck.EndpointID,
				StatusCode:   healthcheck.StatusCode,
				ResponseTime: healthcheck.ResponseTime,
				Success:      healthcheck.Success,
				CheckedAt:    healthcheck.CheckedAt,
			},
		)
	}
	c.JSON(http.StatusOK, response)
}

func (h *HealthCheckHandler) GetByEndpointID(c *gin.Context) {

	endpointID := c.Param("id")

	id, err := strconv.ParseUint(endpointID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid endpoint id",
		})
		return
	}

	checks, err := h.healthCheckService.GetByEndpointID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := make([]HealthCheckResponse, 0, len(checks))

	for _, healthcheck := range checks {
		response = append(
			response,
			HealthCheckResponse{
				ID:           healthcheck.ID,
				EndpointID:   healthcheck.EndpointID,
				StatusCode:   healthcheck.StatusCode,
				ResponseTime: healthcheck.ResponseTime,
				Success:      healthcheck.Success,
				CheckedAt:    healthcheck.CheckedAt,
			},
		)
	}
	c.JSON(http.StatusOK, response)
}

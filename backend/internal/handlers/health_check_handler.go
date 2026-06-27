package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	healthCheckService *services.HealthCheckService
}

func NewHealthCheckHandler(healthCheckService *services.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthCheckService: healthCheckService,
	}
}

// GetAllHealthChecks godoc
//
// @Summary List all health checks
// @Description Returns all recorded health checks.
// @Tags Health Checks
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.HealthCheckListResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /healthchecks [get]
func (h *HealthCheckHandler) GetAllHealthChecks(c *gin.Context) {
	ctx := c.Request.Context()

	page, limit := utils.GetPaginationParams(c)

	endpointIDStr := c.Query("endpoint_id")
	var endpointID uint
	if endpointIDStr != "" {
		if id, err := strconv.ParseUint(endpointIDStr, 10, 64); err == nil {
			endpointID = uint(id)
		}
	}

	success := c.Query("success")

	checks, total, err := h.healthCheckService.GetAllPaginated(ctx, page, limit, endpointID, success)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.PaginatedOK(c, dto.ToHealthCheckResponses(checks), page, limit, total)
}

// GetByEndpointID godoc
//
// @Summary Get health checks for an endpoint
// @Description Returns all health checks belonging to a specific endpoint.
// @Tags Health Checks
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.HealthCheckListResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /healthchecks/{id} [get]
func (h *HealthCheckHandler) GetByEndpointID(c *gin.Context) {
	ctx := c.Request.Context()

	endpointID := c.Param("id")

	id, err := strconv.ParseUint(endpointID, 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid endpoint ID.")
		return
	}

	checks, err := h.healthCheckService.GetByEndpointID(ctx, uint(id))

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, dto.ToHealthCheckResponses(checks))
}

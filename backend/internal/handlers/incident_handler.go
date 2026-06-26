package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type IncidentHandler struct {
	incidentService *services.IncidentService
}

func NewIncidentHandler(incidentService *services.IncidentService) *IncidentHandler {
	return &IncidentHandler{
		incidentService: incidentService,
	}
}

// ListIncidents godoc
//
// @Summary List incidents
// @Description Returns all recorded incidents.
// @Tags Incidents
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.IncidentListResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /incidents [get]
func (h *IncidentHandler) ListIncidents(c *gin.Context) {
	ctx := c.Request.Context()

	incidents, err := h.incidentService.GetAllIncidents(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, dto.ToIncidentResponses(incidents))
}

// GetIncidentByID godoc
//
// @Summary Get incident
// @Description Returns an incident by ID.
// @Tags Incidents
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Incident ID"
//
// @Success 200 {object} dto.IncidentSuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /incidents/{id} [get]
func (h *IncidentHandler) GetIncidentByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	incidentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		utils.BadRequest(c, "Invalid incident ID.")
		return
	}

	incident, err := h.incidentService.GetIncidentByID(ctx, uint(incidentID))

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, dto.ToIncidentResponse(*incident))
}

// GetActiveIncidents godoc
//
// @Summary List active incidents
// @Description Returns all unresolved incidents.
// @Tags Incidents
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.IncidentListResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /incidents/active [get]
func (h *IncidentHandler) GetActiveIncidents(c *gin.Context) {
	ctx := c.Request.Context()

	incidents, err := h.incidentService.GetActiveIncidents(ctx)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, dto.ToIncidentResponses(incidents))
}

// GetIncidentByEndpointID godoc
//
// @Summary Get active incident for endpoint
// @Description Returns the currently active incident for an endpoint.
// @Tags Incidents
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Param id path int true "Endpoint ID"
//
// @Success 200 {object} dto.IncidentSuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
//
// @Router /endpoints/{id}/incidents [get]
func (h *IncidentHandler) GetIncidentByEndpointID(c *gin.Context) {
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

	incident, err := h.incidentService.GetActiveIncidentByEndpointID(ctx, uint(id))

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, dto.ToIncidentResponse(*incident))
}

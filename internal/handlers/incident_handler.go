package handlers

import (
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
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

func (h *IncidentHandler) ListIncidents(c *gin.Context) {

	incidents, err := h.incidentService.GetAllIncidents()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	var response []dto.IncidentResponse

	for _, incident := range incidents {
		response = append(
			response,
			dto.IncidentResponse{
				ID:         incident.ID,
				EndpointID: incident.EndpointID,
				StartedAt:  incident.StartedAt,
				ResolvedAt: incident.ResolvedAt,
				IsResolved: incident.IsResolved,
			},
		)
	}

	utils.OK(c, response)
}

func (h *IncidentHandler) GetIncidentByID(c *gin.Context) {

	id := c.Param("id")

	incidentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		utils.BadRequest(c, "Invalid incident ID.")
		return
	}

	incident, err := h.incidentService.GetIncidentByID(uint(incidentID))

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	response := dto.IncidentResponse{
		ID:         incident.ID,
		EndpointID: incident.EndpointID,
		StartedAt:  incident.StartedAt,
		ResolvedAt: incident.ResolvedAt,
		IsResolved: incident.IsResolved,
	}

	utils.OK(c, response)
}

func (h *IncidentHandler) GetActiveIncidents(c *gin.Context) {

	incidents, err := h.incidentService.GetActiveIncidents()

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, incidents)
}

func (h *IncidentHandler) GetIncidentByEndpointID(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		utils.BadRequest(c, "Invalid endpoint ID.")
		return
	}

	incidents, err := h.incidentService.GetActiveIncidentByEndpointID(uint(id))

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, incidents)
}

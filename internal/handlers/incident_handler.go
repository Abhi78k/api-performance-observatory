package handlers

import (
	"net/http"
	"strconv"

	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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

	c.JSON(http.StatusOK, response)
}

func (h *IncidentHandler) GetIncidentByID(c *gin.Context) {

	id := c.Param("id")

	incidentID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid incident ID.",
		})
		return
	}

	incident, err := h.incidentService.GetIncidentByID(uint(incidentID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.IncidentResponse{
		ID:         incident.ID,
		EndpointID: incident.EndpointID,
		StartedAt:  incident.StartedAt,
		ResolvedAt: incident.ResolvedAt,
		IsResolved: incident.IsResolved,
	}

	c.JSON(http.StatusOK, response)
}

func (h *IncidentHandler) GetActiveIncidents(c *gin.Context) {

	incidents, err := h.incidentService.GetActiveIncidents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, incidents)
}

func (h *IncidentHandler) GetIncidentByEndpointID(c *gin.Context) {

	id, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid endpoint ID.",
		})
		return
	}

	incidents, err := h.incidentService.GetActiveIncidentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, incidents)
}

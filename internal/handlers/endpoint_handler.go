package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
	"github.com/gin-gonic/gin"
)

type EndpointHandler struct {
	endpointService *services.EndpointService
}

func NewEndpointHandler(
	endpointService *services.EndpointService,
) *EndpointHandler {

	return &EndpointHandler{
		endpointService: endpointService,
	}
}

func (h *EndpointHandler) CreateEndpoint(c *gin.Context) {

	var req dto.CreateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request.")
		return
	}

	userID := c.MustGet("UserID").(uint)

	endpoint, err := h.endpointService.CreateEndpoint(
		req.Name,
		req.URL,
		req.ExpectedStatus,
		userID,
	)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.Created(c, endpoint)
}

func (h *EndpointHandler) GetEndpoints(c *gin.Context) {

	userID := c.MustGet("UserID").(uint)

	endpoints, err := h.endpointService.GetEndpoints(userID)

	if err != nil {
		utils.Internal(c, err.Error())
		return
	}

	utils.OK(c, endpoints)
}

func (h *EndpointHandler) GetEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	endpoint, err := h.endpointService.GetEndpoint(
		id,
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.OK(c, endpoint)
}

func (h *EndpointHandler) UpdateEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	var req dto.UpdateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body.")
		return
	}

	endpoint, err := h.endpointService.UpdateEndpoint(
		id,
		userID,
		req.Name,
		req.URL,
		req.ExpectedStatus,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.OK(c, endpoint)
}

func (h *EndpointHandler) DeleteEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	err := h.endpointService.DeleteEndpoint(
		id,
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Endpoint not found.")
		return
	}

	utils.OK(c, "Endpoint deleted successfully.")
}

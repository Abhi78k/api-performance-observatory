package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/gin-gonic/gin"
)

type EndpointHandler struct {
	endpointService *services.EndpointService
}

type CreateEndpointRequest struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	ExpectedStatus int    `json:"expected_status"`
}

type UpdateEndpointRequest struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	ExpectedStatus int    `json:"expected_status"`
}

func NewEndpointHandler(
	endpointService *services.EndpointService,
) *EndpointHandler {

	return &EndpointHandler{
		endpointService: endpointService,
	}
}

func (h *EndpointHandler) CreateEndpoint(c *gin.Context) {

	var req CreateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, endpoint)
}

func (h *EndpointHandler) GetEndpoints(c *gin.Context) {

	userID := c.MustGet("UserID").(uint)

	endpoints, err := h.endpointService.GetEndpoints(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, endpoints)
}

func (h *EndpointHandler) GetEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	endpoint, err := h.endpointService.GetEndpoint(
		id,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func (h *EndpointHandler) UpdateEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	var req UpdateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
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
		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func (h *EndpointHandler) DeleteEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID := c.MustGet("UserID").(uint)

	err := h.endpointService.DeleteEndpoint(
		id,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "endpoint deleted successfully",
	})
}

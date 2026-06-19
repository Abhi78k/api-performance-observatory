package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

type EndpointHandler struct {
	DB *gorm.DB
}

func (h *EndpointHandler) CreateEndpoint(c *gin.Context) {

	var req CreateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}

	endpoint := models.Endpoint{
		Name:           req.Name,
		URL:            req.URL,
		ExpectedStatus: req.ExpectedStatus,
		UserID:         userID.(uint),
	}

	if err := h.DB.Create(&endpoint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create endpoint",
		})
		return
	}

	c.JSON(http.StatusCreated, endpoint)

}

func (h *EndpointHandler) GetEndpoints(c *gin.Context) {

	userID, _ := c.Get("userID")

	var endpoints []models.Endpoint

	if err := h.DB.
		Where("user_id = ?", userID).
		Find(&endpoints).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch endpoints",
		})
		return
	}

	c.JSON(http.StatusOK, endpoints)
}

func (h *EndpointHandler) GetEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("userID")

	var endpoint models.Endpoint

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&endpoint).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func (h *EndpointHandler) UpdateEndpoint(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("userID")

	var endpoint models.Endpoint

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&endpoint).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	var req UpdateEndpointRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	endpoint.Name = req.Name
	endpoint.URL = req.URL
	endpoint.ExpectedStatus = req.ExpectedStatus

	if err := h.DB.Save(&endpoint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update endpoint",
		})
		return
	}

	c.JSON(http.StatusOK, endpoint)
}

func (h *EndpointHandler) DeleteEndpoint(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var endpoint models.Endpoint

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&endpoint).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "endpoint not found",
		})
		return
	}

	if err := h.DB.Delete(&endpoint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete endpoint",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "endpoint deleted successfully",
	})

}

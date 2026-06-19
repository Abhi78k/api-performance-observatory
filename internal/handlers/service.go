package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateServiceRequest struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	ExpectedStatus int    `json:"expected_status"`
}

type UpdateServiceRequest struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	ExpectedStatus int    `json:"expected_status"`
}

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) CreateService(c *gin.Context) {

	var req CreateServiceRequest

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

	service := models.Service{
		Name:           req.Name,
		URL:            req.URL,
		ExpectedStatus: req.ExpectedStatus,
		UserID:         userID.(uint),
	}

	if err := h.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create service",
		})
		return
	}

	c.JSON(http.StatusCreated, service)

}

func (h *ServiceHandler) GetServices(c *gin.Context) {

	userID, _ := c.Get("userID")

	var services []models.Service

	if err := h.DB.
		Where("user_id = ?", userID).
		Find(&services).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch services",
		})
		return
	}

	c.JSON(http.StatusOK, services)
}

func (h *ServiceHandler) GetService(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("userID")

	var service models.Service

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&service).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "service not found",
		})
		return
	}

	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {

	id := c.Param("id")
	userID, _ := c.Get("userID")

	var service models.Service

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&service).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "service not found",
		})
		return
	}

	var req UpdateServiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	service.Name = req.Name
	service.URL = req.URL
	service.ExpectedStatus = req.ExpectedStatus

	if err := h.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update service",
		})
		return
	}

	c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var service models.Service

	if err := h.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&service).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "service not found",
		})
		return
	}

	if err := h.DB.Delete(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "service deleted successfully",
	})

}

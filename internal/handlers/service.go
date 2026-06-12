package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceHandler struct {
	DB *gorm.DB
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Service - not implemented",
	})
}

func (h *ServiceHandler) GetServices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Services - not implemented",
	})
}

func (h *ServiceHandler) GetService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Service - not implemented",
	})
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Service - not implemented",
	})
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Service - not implemented",
	})
}

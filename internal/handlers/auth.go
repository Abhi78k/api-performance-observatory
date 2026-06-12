package handlers

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}
	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"message": "email and password are required.",
		})
		return
	}
}

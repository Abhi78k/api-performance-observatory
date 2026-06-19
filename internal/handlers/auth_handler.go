package handlers

import (
	"net/http"

	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body.",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "email and password are required.",
		})
		return
	}

	err = h.authService.Register(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created.",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body.",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "email and password are required.",
		})
		return
	}

	accessToken, err := h.authService.Login(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("UserID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found.",
		})
		return
	}

	uid, ok := userID.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user.",
		})
		return
	}

	user, err := h.authService.GetUserByID(uid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

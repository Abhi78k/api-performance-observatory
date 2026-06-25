package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/internal/dto"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationError(c, utils.FormatValidationErrors(err))
		return
	}

	err = h.authService.Register(req.Email, req.Password)

	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Created(c, "User created successfully.")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		utils.BadRequest(c, "Invalid request body.")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationError(c, utils.FormatValidationErrors(err))
		return
	}

	accessToken, err := h.authService.Login(req.Email, req.Password)

	if err != nil {
		utils.Unauthorized(c, "Invalid credentials.")
		return
	}

	utils.OK(c, gin.H{
		"access_token": accessToken,
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("UserID")

	if !exists {
		utils.Unauthorized(c, "User not found.")
		return
	}

	uid, ok := userID.(uint)

	if !ok {
		utils.Unauthorized(c, "Invalid user.")
		return
	}

	user, err := h.authService.GetUserByID(uid)

	if err != nil {
		utils.NotFound(c, "User not found.")
		return
	}

	utils.OK(c, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

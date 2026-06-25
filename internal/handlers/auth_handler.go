package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	"github.com/Abhi78k/api-performance-observatory/internal/utils"
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
		utils.BadRequest(c, err.Error())
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.BadRequest(c, "Email and password are required.")
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
	var req LoginRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		utils.BadRequest(c, "Invalid request body.")
		return
	}

	if req.Email == "" || req.Password == "" {
		utils.BadRequest(c, "Email and password are required.")
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

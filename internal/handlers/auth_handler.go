package handlers

import (
	"net/http"

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

// Register godoc
//
// @Summary Register a new user
// @Description Creates a new user account.
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body dto.RegisterRequest true "Registration details"
//
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ValidationErrorResponse
//
// @Router /auth/register [post]
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

	utils.Message(c, http.StatusCreated, "User created successfully.")
}

// Login godoc
//
// @Summary Authenticate user
// @Description Authenticates a user and returns a JWT access token.
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body dto.LoginRequest true "Login credentials"
//
// @Success 200 {object} dto.LoginSuccessResponse
// @Failure 400 {object} utils.ValidationErrorResponse
// @Failure 401 {object} utils.ErrorResponse
//
// @Router /auth/login [post]
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

	utils.OK(c, dto.LoginResponse{
		AccessToken: accessToken,
	})
}

// GetMe godoc
//
// @Summary Get current user
// @Description Returns the currently authenticated user.
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Security BearerAuth
//
// @Success 200 {object} dto.UserSuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
//
// @Router /auth/me [get]
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

	utils.OK(c, dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

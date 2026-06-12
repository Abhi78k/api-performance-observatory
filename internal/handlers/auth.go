package handlers

import (
	"github.com/Abhi78k/api-performance-observatory/internal/auth"
	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
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

	var existingUser models.User

	result := database.DB.Where("email = ?", req.Email).First(&existingUser)

	if result.Error == nil {
		c.JSON(409, gin.H{
			"error": "user already exists.",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to hash password.",
		})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	result = database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "failed to create user.",
		})
		return
	}

	c.JSON(201, gin.H{
		"id":      user.ID,
		"email":   user.Email,
		"message": "user created successfully.",
	})
}

func Login(c *gin.Context) {
	var req LoginRequest

	err := c.ShouldBindJSON(&req)

	if req.Email == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "email and password are required.",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body.",
		})
		return
	}
	var user models.User

	result := database.DB.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		c.JSON(401, gin.H{
			"error": "invalid credentials.",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid credentials.",
		})
		return
	}

	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to generate token.",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Me(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{
			"error": "user not found in context.",
		})
		return
	}

	uid := userID.(uint)

	var user models.User

	result := database.DB.First(&user, uid)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "invalid user.",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func Refresh(c *gin.Context) {

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{
			"error": "user not found.",
		})
		return
	}

	token, err := auth.GenerateToken(
		userID.(uint),
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "failed to generate token.",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logged out successfully.",
	})
}

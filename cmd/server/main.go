package main

import (
	"log"

	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/middleware"
	"github.com/Abhi78k/api-performance-observatory/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load Config 🐽
	cfg := config.Load()

	// Database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected!")

	// Migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.HealthCheck{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Handlers
	serviceHandler := handlers.ServiceHandler{
		DB: db,
	}

	// Router
	r := gin.Default()

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Observatory",
		})
	})

	// -------------------
	// Public Auth Routes
	// -------------------
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
	}

	// -------------------
	// Protected Auth Routes
	// -------------------
	protectedAuth := r.Group("/auth")
	protectedAuth.Use(middleware.AuthMiddleware())
	{
		protectedAuth.GET("/me", handlers.Me)
		protectedAuth.POST("/refresh", handlers.Refresh)
		protectedAuth.POST("/logout", handlers.Logout)

		protectedAuth.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")

			c.JSON(200, gin.H{
				"user_id": userID,
			})
		})
	}

	// -------------------
	// Protected Service Routes
	// -------------------
	services := r.Group("/services")
	services.Use(middleware.AuthMiddleware())
	{
		services.POST("", serviceHandler.CreateService)
		services.GET("", serviceHandler.GetServices)
		services.GET("/:id", serviceHandler.GetService)
		services.PUT("/:id", serviceHandler.UpdateService)
		services.DELETE("/:id", serviceHandler.DeleteService)
	}

	log.Println("Server running on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"

	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/middleware"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/gin-gonic/gin"
	
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	_ = db

	db.AutoMigrate(
		&models.User{},
		&models.Service{},
	)

	// serviceHandler := handlers.ServiceHandler{
	// 	DB: db,
	// }
	r := gin.Default()

	// r.POST("/services", serviceHandler.CreateService)
	// r.GET("/services", serviceHandler.GetServices)
	// r.GET("/services/:id", serviceHandler.GetService)
	// r.PUT("/services/:id", serviceHandler.UpdateService)
	// r.DELETE("/services/:id", serviceHandler.DeleteService)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Observatory",
		})
	})

	authRoutes := r.Group("/auth")
	authRoutes.POST("/register", handlers.Register)
	authRoutes.POST("/login", handlers.Login)

	protected := r.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/me", handlers.Me)
	protected.POST("/refresh", handlers.Refresh)
	protected.POST("/logout", handlers.Logout)

	protected.GET("/profile", func(c *gin.Context) {
		userID, _ := c.Get("userID")

		c.JSON(200, gin.H{
			"user_id": userID,
		})
	})

	r.Run(":8080")

}

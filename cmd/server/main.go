package main

import (
	"fmt"

	"github.com/Abhi78k/api-performance-observatory/internal/auth"
	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	token, err := auth.GenerateToken(1)

	fmt.Println(token)
	fmt.Println(err)

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

	serviceHandler := handlers.ServiceHandler{
		DB: db,
	}

	r.POST("/services", serviceHandler.CreateService)
	r.GET("/services", serviceHandler.GetServices)
	r.GET("/services/:id", serviceHandler.GetService)
	r.PUT("/services/:id", serviceHandler.UpdateService)
	r.DELETE("/services/:id", serviceHandler.DeleteService)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Observatory",
		})
	})

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)
	r.Run(":8080")

}

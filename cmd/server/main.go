package main

import (
	"log"

	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
	"github.com/Abhi78k/api-performance-observatory/internal/routes"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
)

func main() {

	// Load Config
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
		&models.Endpoint{},
		&models.HealthCheck{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	authRepo := repositories.NewUserRepository(db)
	endpointRepo := repositories.NewEndpointRepository(db)
	healthCheckRepo := repositories.NewHealthCheckRepo(db)
	incidentRepo := repositories.NewIncidentRepository(db)

	authService := services.NewAuthService(cfg, authRepo)
	endpointService := services.NewEndpointService(endpointRepo)
	incidentService := services.NewIncidentService(incidentRepo)
	healthCheckService := services.NewHealthCheckService(endpointRepo, healthCheckRepo, incidentService)
	schedulerService := services.NewSchedulerService(endpointRepo, healthCheckService)

	go schedulerService.Start()

	authHandler := handlers.NewAuthHandler(authService)
	endpointHandler := handlers.NewEndpointHandler(endpointService)
	statsHandler := handlers.NewStatsHandler(healthCheckService)
	healthCheckHandler := handlers.NewHealthCheckHandler(healthCheckService)

	router := routes.SetupRouter(cfg, authHandler, endpointHandler, statsHandler, healthCheckHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// // Handlers
	// endpointHandler := handlers.EndpointHandler{
	// 	DB: db,
	// }

	// // Router
	// r := gin.Default()

	// // Root route
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "API Observatory",
	// 	})
	// })

	// // -------------------
	// // Public Auth Routes
	// // -------------------
	// authRoutes := r.Group("/auth")
	// {
	// 	authRoutes.POST("/register", handlers.Register)
	// 	authRoutes.POST("/login", handlers.Login)
	// }

	// // -------------------
	// // Protected Auth Routes
	// // -------------------
	// protectedAuth := r.Group("/auth")
	// protectedAuth.Use(middleware.AuthMiddleware())
	// {
	// 	protectedAuth.GET("/me", handlers.Me)
	// 	protectedAuth.POST("/refresh", handlers.Refresh)
	// 	protectedAuth.POST("/logout", handlers.Logout)

	// 	protectedAuth.GET("/profile", func(c *gin.Context) {
	// 		userID, _ := c.Get("userID")

	// 		c.JSON(200, gin.H{
	// 			"user_id": userID,
	// 		})
	// 	})
	// }

	// // -------------------
	// // Protected Endpoint Routes
	// // -------------------
	// endpoints := r.Group("/endpoints")
	// endpoints.Use(middleware.AuthMiddleware())
	// {
	// 	endpoints.POST("", endpointHandler.CreateEndpoint)
	// 	endpoints.GET("", endpointHandler.GetEndpoints)
	// 	endpoints.GET("/:id", endpointHandler.GetEndpoint)
	// 	endpoints.PUT("/:id", endpointHandler.UpdateEndpoint)
	// 	endpoints.DELETE("/:id", endpointHandler.DeleteEndpoint)
	// }

	// log.Println("Server running on :8080")

	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal(err)
	// }
}

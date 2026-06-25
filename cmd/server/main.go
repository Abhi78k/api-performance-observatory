package main

import (
	"os"

	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/database"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
	"github.com/Abhi78k/api-performance-observatory/internal/routes"
	"github.com/Abhi78k/api-performance-observatory/internal/services"
	_ "github.com/Abhi78k/api-performance-observatory/docs"
)

// @title API Performance Observatory API
// @version 1.0
// @description Backend API for monitoring HTTP endpoints.

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	logger.Init()

	// Load Config
	cfg := config.Load()

	// Database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Error(
			"Failed to connect to database",
			"error",
			err,
		)
		os.Exit(1)
	}

	logger.Info("Database connected!")

	// Migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Endpoint{},
		&models.HealthCheck{},
		&models.Incident{},
		&models.Monitoring{},
	)

	if err != nil {
		logger.Error(
			"Migration failed",
			"error",
			err,
		)
		os.Exit(1)
	}

	authRepo := repositories.NewUserRepository(db)
	endpointRepo := repositories.NewEndpointRepository(db)
	healthCheckRepo := repositories.NewHealthCheckRepo(db)
	incidentRepo := repositories.NewIncidentRepository(db)
	monitoringRepo := repositories.NewMonitoringRepository(db)

	authService := services.NewAuthService(cfg, authRepo)
	monitoringService := services.NewMonitoringService(monitoringRepo)
	endpointService := services.NewEndpointService(endpointRepo, monitoringService)
	incidentService := services.NewIncidentService(incidentRepo)
	healthCheckService := services.NewHealthCheckService(endpointRepo, healthCheckRepo, incidentService)
	schedulerService := services.NewSchedulerService(endpointRepo, healthCheckService)
	dashboardService := services.NewDashboardService(endpointRepo, healthCheckRepo, incidentRepo, monitoringRepo)
	go schedulerService.Start()

	authHandler := handlers.NewAuthHandler(authService)
	endpointHandler := handlers.NewEndpointHandler(endpointService)
	statsHandler := handlers.NewStatsHandler(healthCheckService)
	healthCheckHandler := handlers.NewHealthCheckHandler(healthCheckService)
	incidentHandler := handlers.NewIncidentHandler(incidentService)
	monitoringHandler := handlers.NewMonitoringHandler(monitoringService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := routes.SetupRouter(cfg, authHandler, endpointHandler, statsHandler, healthCheckHandler, incidentHandler, monitoringHandler, dashboardHandler)

	if err := router.Run(":8080"); err != nil {
		logger.Error(
			"Failed to start server",
			"error",
			err,
		)
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

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Abhi78k/api-performance-observatory/backend/docs"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/config"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/database"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/repositories"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/routes"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/services"
)

// @title API Performance Observatory API
// @version 1.0
// @description Backend API for monitoring HTTP endpoints.

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
	endpointService := services.NewEndpointService(endpointRepo, monitoringService, healthCheckRepo)
	incidentService := services.NewIncidentService(incidentRepo, endpointRepo)
	healthCheckService := services.NewHealthCheckService(endpointRepo, healthCheckRepo, incidentService)
	schedulerService := services.NewSchedulerService(endpointRepo, healthCheckService)
	dashboardService := services.NewDashboardService(endpointRepo, healthCheckRepo, incidentRepo, monitoringRepo)

	appCtx, stop := context.WithCancel(context.Background())
	defer stop()
	go schedulerService.Start(appCtx)

	authHandler := handlers.NewAuthHandler(authService)
	endpointHandler := handlers.NewEndpointHandler(endpointService)
	statsHandler := handlers.NewStatsHandler(healthCheckService)
	healthCheckHandler := handlers.NewHealthCheckHandler(healthCheckService)
	incidentHandler := handlers.NewIncidentHandler(incidentService)
	monitoringHandler := handlers.NewMonitoringHandler(monitoringService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := routes.SetupRouter(cfg, authHandler, endpointHandler, statsHandler, healthCheckHandler, incidentHandler, monitoringHandler, dashboardHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	logger.Info("Server started on :" + port)

	go func() {

		logger.Info("Server started on :8080")

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			logger.Error(
				"Server failed",
				"error",
				err,
			)
		}

	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		logger.Error(
			"Server shutdown failed",
			"error",
			err,
		)

		os.Exit(1)
	}

	logger.Info("Server stopped gracefully.")
	stop()
}

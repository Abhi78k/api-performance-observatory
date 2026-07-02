package routes

import (
	"net/http"
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/config"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	cfg *config.Config,
	authHandler *handlers.AuthHandler,
	endpointHandler *handlers.EndpointHandler,
	statsHandler *handlers.StatsHandler,
	healthCheckHandler *handlers.HealthCheckHandler,
	incidentsHandler *handlers.IncidentHandler,
	monitoringHandler *handlers.MonitoringHandler,
	dashboardHandler *handlers.DashboardHandler,
) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			cfg.FrontendURL,
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},

		ExposeHeaders: []string{
			"Content-Length",
		},

		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))

	auth := router.Group("/auth")
	{
		auth.POST(
			"/register",
			authHandler.Register,
		)

		auth.POST(
			"/login",
			authHandler.Login,
		)

		auth.POST(
			"/logout",
			authHandler.Logout,
		)
	}

	stats := router.Group("/endpoints")
	stats.Use(
		middleware.AuthMiddleware(cfg),
	)

	{
		stats.GET(
			"/:id/stats",
			statsHandler.GetEndpointStats,
		)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	protected := router.Group("/")
	protected.Use(
		middleware.AuthMiddleware(cfg),
	)

	{
		protected.GET(
			"/auth/me",
			authHandler.GetMe,
		)

		protected.POST(
			"/endpoints",
			endpointHandler.CreateEndpoint,
		)

		protected.GET(
			"/endpoints",
			endpointHandler.GetEndpoints,
		)

		protected.GET(
			"/endpoints/:id",
			endpointHandler.GetEndpoint,
		)

		protected.GET(
			"/endpoints/:id/incidents",
			incidentsHandler.GetIncidentByEndpointID,
		)

		protected.GET(
			"/endpoints/:id/monitoring",
			monitoringHandler.GetMonitoring,
		)
		protected.PUT(
			"/endpoints/:id",
			endpointHandler.UpdateEndpoint,
		)

		protected.DELETE(
			"/endpoints/:id",
			endpointHandler.DeleteEndpoint,
		)
	}

	healthCheck := router.Group("/healthchecks")
	healthCheck.Use(middleware.AuthMiddleware(cfg))

	{
		healthCheck.GET(
			"",
			healthCheckHandler.GetAllHealthChecks,
		)

		healthCheck.GET(
			"/:id",
			healthCheckHandler.GetByEndpointID,
		)
	}

	incidents := router.Group("/incidents")
	incidents.Use(middleware.AuthMiddleware(cfg))

	{
		incidents.GET(
			"",
			incidentsHandler.ListIncidents,
		)

		incidents.GET(
			"/:id",
			incidentsHandler.GetIncidentByID,
		)

		incidents.GET(
			"/active",
			incidentsHandler.GetActiveIncidents,
		)
	}

	dashboard := router.Group("/dashboard")
	dashboard.Use(middleware.AuthMiddleware(cfg))

	{
		dashboard.GET(
			"/overview",
			dashboardHandler.GetOverview,
		)

		dashboard.GET(
			"/status",
			dashboardHandler.GetStatus,
		)

		dashboard.GET(
			"/incidents",
			dashboardHandler.GetRecentIncidents,
		)

		dashboard.GET(
			"/performance",
			dashboardHandler.GetPerformance,
		)

		dashboard.GET(
			"/success-rate",
			dashboardHandler.GetSuccessRate,
		)

		dashboard.GET(
			"/uptime",
			dashboardHandler.GetUptime,
		)

		dashboard.GET(
			"/history",
			dashboardHandler.GetHistory,
		)

		dashboard.GET(
			"/monitoring",
			dashboardHandler.GetMonitoring,
		)
	}

	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	return router
}

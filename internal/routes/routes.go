package routes

import (
	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/handlers"
	"github.com/Abhi78k/api-performance-observatory/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	cfg *config.Config,
	authHandler *handlers.AuthHandler,
	endpointHandler *handlers.EndpointHandler,
	statsHandler *handlers.StatsHandler,
	healthCheckHandler *handlers.HealthCheckHandler,
	incidentsHandler *handlers.IncidentHandler,
	dashboardHandler *handlers.DashboardHandler,
) *gin.Engine {

	router := gin.Default()

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
	protected.Use(middleware.AuthMiddleware(cfg))

	{
		healthCheck.GET(
			"/",
			healthCheckHandler.GetByEndpointID,
		)

		healthCheck.GET(
			"/:id",
			healthCheckHandler.GetAllHealthChecks,
		)
	}

	incidents := router.Group("/incidents")
	protected.Use(middleware.AuthMiddleware(cfg))

	{
		incidents.GET(
			"/",
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
			"/dashboard/performance",
			dashboardHandler.GetPerformance,
		)

		dashboard.GET(
			"/dashboard/success-rate",
			dashboardHandler.GetSuccessRate,
		)

		dashboard.GET(
			"/dashboard/uptime",
			dashboardHandler.GetUptime,
		)

		dashboard.GET(
			"/dashboard/history",
			dashboardHandler.GetHistory,
		)
	}

	return router
}

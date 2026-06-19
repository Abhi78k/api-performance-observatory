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

		protected.PUT(
			"/endpoints/:id",
			endpointHandler.UpdateEndpoint,
		)

		protected.DELETE(
			"/endpoints/:id",
			endpointHandler.DeleteEndpoint,
		)
	}

	return router
}

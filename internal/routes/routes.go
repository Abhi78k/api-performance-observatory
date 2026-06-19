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
) *gin.Engine {

	router := gin.Default()

	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware(cfg))

	auth := router.Group("/auth")

	auth.POST(
		"/register",
		authHandler.Register,
	)

	auth.POST(
		"/login",
		authHandler.Login,
	)

	protected.GET(
		"/me",
		authHandler.GetMe,
	)

	return router
}

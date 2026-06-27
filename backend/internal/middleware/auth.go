package middleware

import (
	"github.com/Abhi78k/api-performance-observatory/backend/internal/auth"
	"github.com/Abhi78k/api-performance-observatory/backend/internal/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(401, gin.H{
				"error": "unauthorized: access token cookie missing.",
			})
			c.Abort()
			return
		}

		token, err := auth.ValidateToken(cfg, tokenString)

		// if invalid
		if err != nil {
			c.JSON(401, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		// verify claims
		claims, ok := token.Claims.(*auth.Claims)

		if !ok {
			c.JSON(401, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		// set the user id
		c.Set("UserID", claims.UserID)

		c.Next()
	}
}

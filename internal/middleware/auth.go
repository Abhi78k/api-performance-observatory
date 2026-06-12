package middleware

import (
	"strings"

	"github.com/Abhi78k/api-performance-observatory/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve the header from jwt token
		header := c.GetHeader("Authorization")

		// if header empty, return error
		if header == "" {
			c.JSON(401, gin.H{
				"error": "authorization header required.",
			})
			c.Abort()
			return
		}

		// extract token
		parts := strings.Split(header, " ")

		// verify token format
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"error": "invalid authorization format.",
			})
			c.Abort()
			return
		}

		// extract the token string
		tokenString := parts[1]

		token, err := auth.ValidateToken(tokenString)

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
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

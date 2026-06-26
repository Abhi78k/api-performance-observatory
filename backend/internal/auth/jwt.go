package auth

import (
	"time"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(cfg *config.Config, userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(cfg.JWTSecret),
	)
}

func ValidateToken(
	cfg *config.Config,
	tokenString string,
) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			return []byte(
				cfg.JWTSecret,
			), nil
		},
	)
}

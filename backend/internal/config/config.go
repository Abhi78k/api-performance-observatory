package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBPassword string
	DBUser     string

	DatabaseURL string
	FrontendURL string

	JWTSecret string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	return &Config{
		DBName:      os.Getenv("DB_NAME"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBUser:      os.Getenv("DB_USER"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		FrontendURL: os.Getenv("FRONTEND_URL"),
	}
}

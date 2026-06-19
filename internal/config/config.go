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

	JWTSecret string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env")
	}

	return &Config{
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}

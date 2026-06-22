package models

import (
	"time"
)

type HealthCheck struct {
	ID           uint `gorm:"primaryKey"`
	EndpointID   uint
	StatusCode   int
	ResponseTime int64
	Success      bool
	CheckedAt    time.Time
}

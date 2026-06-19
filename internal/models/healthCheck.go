package models

import (
	"time"
)

type HealthCheck struct {
	ID           uint `gorm: "primaryKey"`
	ServiceID    uint
	StatusCode   int
	ResponseTime int64
	Success      bool
	CheckedAt    time.Time
}

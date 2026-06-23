package models

import "time"

type Monitoring struct {
	ID                  uint      `gorm:"primaryKey"`
	EndpointID          uint      `gorm:"not null;unique"`
	MonitoringStartedAt time.Time `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import (
	"time"
)

type Incident struct {
	ID         uint       `gorm:"primaryKey"`
	EndpointID uint       `gorm:"not null"`
	StartedAt  time.Time  `gorm:"not null"`
	ResolvedAt *time.Time `gorm:"default:null"`
	IsResolved bool       `gorm:"default:false"`
}

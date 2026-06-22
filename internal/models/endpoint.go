package models

import "time"

type Endpoint struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	URL            string `gorm:"not null"`
	ExpectedStatus int    `gorm:"default:200"`
	CheckInterval  int
	LastCheckedAt  *time.Time
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

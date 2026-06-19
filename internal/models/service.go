package models

import "time"

type Service struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null"`
	URL            string `gorm:"not null"`
	ExpectedStatus int    `gorm:"default:200"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

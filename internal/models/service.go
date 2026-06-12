package models

type Service struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	URL            string
	ExpectedStatus int
	UserID         uint
}

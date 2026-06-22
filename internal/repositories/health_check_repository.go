package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type HealthCheckRepository struct {
	DB *gorm.DB
}

func NewHealthCheckRepo(db *gorm.DB) *HealthCheckRepository {
	return &HealthCheckRepository{
		DB: db,
	}
}

// Create saves a health check record
func (r *HealthCheckRepository) Create(check *models.HealthCheck) error {
	return r.DB.Create(check).Error
}

// GetByEndpointID returns all health checks for a endpoint
func (r *HealthCheckRepository) GetByEndpointID(endpointID uint) ([]models.HealthCheck, error) {

	var checks []models.HealthCheck

	err := r.DB.
		Where("endpoint_id = ?", endpointID).
		Find(&checks).Error

	return checks, err
}

func (r *HealthCheckRepository) GetAll() ([]models.HealthCheck, error) {
	var checks []models.HealthCheck

	err := r.DB.Order("checked_at DESC").Find(&checks).Error

	return checks, err
}

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

// GetByServiceID returns all health checks for a service
func (r *HealthCheckRepository) GetByEndpointID(endpointID uint) ([]models.HealthCheck, error) {

	var checks []models.HealthCheck

	err := r.DB.
		Where("endpoint_id = ?", endpointID).
		Find(&checks).Error

	return checks, err
}

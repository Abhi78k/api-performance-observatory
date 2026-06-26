package repositories

import (
	"context"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type HealthCheckRepositoryInterface interface {
	Create(ctx context.Context, check *models.HealthCheck) error
	GetByEndpointID(ctx context.Context, endpointID uint) ([]models.HealthCheck, error)
	GetAll(ctx context.Context) ([]models.HealthCheck, error)
	GetLatestByEndpointID(
		ctx context.Context,
		endpointID uint,
	) (*models.HealthCheck, error)
}

type HealthCheckRepository struct {
	DB *gorm.DB
}

func NewHealthCheckRepo(db *gorm.DB) *HealthCheckRepository {
	return &HealthCheckRepository{
		DB: db,
	}
}

// Create saves a health check record
func (r *HealthCheckRepository) Create(ctx context.Context, check *models.HealthCheck) error {
	return r.DB.WithContext(ctx).Create(check).Error
}

// GetByEndpointID returns all health checks for a endpoint
func (r *HealthCheckRepository) GetByEndpointID(ctx context.Context, endpointID uint) ([]models.HealthCheck, error) {

	var checks []models.HealthCheck

	err := r.DB.WithContext(ctx).
		Where("endpoint_id = ?", endpointID).
		Find(&checks).Error

	return checks, err
}

func (r *HealthCheckRepository) GetAll(ctx context.Context) ([]models.HealthCheck, error) {
	var checks []models.HealthCheck

	err := r.DB.WithContext(ctx).Order("checked_at DESC").Find(&checks).Error

	return checks, err
}

func (r *HealthCheckRepository) GetLatestByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.HealthCheck, error) {

	var check models.HealthCheck

	err := r.DB.WithContext(ctx).Where("endpoint_id = ?", endpointID).Order("checked_at DESC").First(&check).Error

	if err != nil {
		return nil, err
	}

	return &check, nil
}

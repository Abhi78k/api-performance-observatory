package repositories

import (
	"context"

	"github.com/Abhi78k/api-performance-observatory/backend/internal/models"
	"gorm.io/gorm"
)

type MonitoringRepositoryInterface interface {
	Create(
		ctx context.Context,
		monitoring *models.Monitoring,
	) error
	GetByEndpointID(
		ctx context.Context,
		endpointID uint,
	) (*models.Monitoring, error)
}

type MonitoringRepository struct {
	db *gorm.DB
}

func NewMonitoringRepository(db *gorm.DB) *MonitoringRepository {
	return &MonitoringRepository{
		db: db,
	}
}

func (r *MonitoringRepository) Create(
	ctx context.Context,
	monitoring *models.Monitoring,
) error {
	return r.db.WithContext(ctx).Create(monitoring).Error
}

func (r *MonitoringRepository) GetByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.Monitoring, error) {

	var monitoring models.Monitoring

	err := r.db.WithContext(ctx).
		Where("endpoint_id = ?", endpointID).
		First(&monitoring).
		Error

	if err != nil {
		return nil, err
	}

	return &monitoring, nil
}

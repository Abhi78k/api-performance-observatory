package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type MonitoringRepository struct {
	db *gorm.DB
}

func NewMonitoringRepository(db *gorm.DB) *MonitoringRepository {
	return &MonitoringRepository{
		db: db,
	}
}

func (r *MonitoringRepository) Create(
	monitoring *models.Monitoring,
) error {
	return r.db.Create(monitoring).Error
}

func (r *MonitoringRepository) GetByEndpointID(
	endpointID uint,
) (*models.Monitoring, error) {

	var monitoring models.Monitoring

	err := r.db.
		Where("endpoint_id = ?", endpointID).
		First(&monitoring).
		Error

	if err != nil {
		return nil, err
	}

	return &monitoring, nil
}

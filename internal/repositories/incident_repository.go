package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{
		db: db,
	}
}

func (r *IncidentRepository) Create(
	incident *models.Incident,
) error {
	return r.db.Create(incident).Error
}

func (r *IncidentRepository) GetActiveIncident(
	endpointID uint,
) (*models.Incident, error) {

	var incident models.Incident

	err := r.db.
		Where(
			"endpoint_id = ? AND is_resolved = ?",
			endpointID,
			false,
		).
		First(&incident).
		Error

	if err != nil {
		return nil, err
	}

	return &incident, nil
}

func (r *IncidentRepository) Update(
	incident *models.Incident,
) error {
	return r.db.Save(incident).Error
}

func (r *IncidentRepository) GetByEndpointID(
	endpointID uint,
) ([]models.Incident, error) {

	var incidents []models.Incident

	err := r.db.
		Where("endpoint_id = ?", endpointID).
		Find(&incidents).
		Error

	return incidents, err
}

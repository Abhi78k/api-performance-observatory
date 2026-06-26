package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type IncidentRepositoryInterface interface {
	Create(
		incident *models.Incident,
	) error
	GetActiveIncidentByEndpointID(
		endpointID uint,
	) (*models.Incident, error)
	Update(
		incident *models.Incident,
	) error
	GetByEndpointID(
		endpointID uint,
	) ([]models.Incident, error)
	GetAllIncidents() ([]models.Incident, error)
	GetIncidentByID(id uint) (*models.Incident, error)
	GetActiveIncidents() (
		[]models.Incident,
		error,
	)
	GetRecentIncidents() (
		[]models.Incident,
		error,
	)
}

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

func (r *IncidentRepository) GetActiveIncidentByEndpointID(
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

func (r *IncidentRepository) GetAllIncidents() ([]models.Incident, error) {

	var incidents []models.Incident

	err := r.db.Order("started_at DESC").Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (r *IncidentRepository) GetIncidentByID(id uint) (*models.Incident, error) {

	var incident models.Incident

	err := r.db.Where("ID = ?", id).First(&incident).Error

	if err != nil {
		return nil, err
	}

	return &incident, nil
}

func (r *IncidentRepository) GetActiveIncidents() (
	[]models.Incident,
	error,
) {

	var incidents []models.Incident

	err := r.db.Where("is_resolved = ?", false).Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (r *IncidentRepository) GetRecentIncidents() (
	[]models.Incident,
	error,
) {

	var incidents []models.Incident

	err := r.db.Order("started_at DESC").Limit(10).Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, nil
}

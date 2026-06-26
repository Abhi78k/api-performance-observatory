package repositories

import (
	"context"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type IncidentRepositoryInterface interface {
	Create(
		ctx context.Context,
		incident *models.Incident,
	) error
	GetActiveIncidentByEndpointID(
		ctx context.Context,
		endpointID uint,
	) (*models.Incident, error)
	Update(
		ctx context.Context,
		incident *models.Incident,
	) error
	GetByEndpointID(
		ctx context.Context,
		endpointID uint,
	) ([]models.Incident, error)
	GetAllIncidents(ctx context.Context) ([]models.Incident, error)
	GetIncidentByID(ctx context.Context, id uint) (*models.Incident, error)
	GetActiveIncidents(ctx context.Context) (
		[]models.Incident,
		error,
	)
	GetRecentIncidents(ctx context.Context) (
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
	ctx context.Context,
	incident *models.Incident,
) error {
	return r.db.WithContext(ctx).Create(incident).Error
}

func (r *IncidentRepository) GetActiveIncidentByEndpointID(
	ctx context.Context,
	endpointID uint,
) (*models.Incident, error) {

	var incident models.Incident

	err := r.db.WithContext(ctx).
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
	ctx context.Context,
	incident *models.Incident,
) error {
	return r.db.WithContext(ctx).Save(incident).Error
}

func (r *IncidentRepository) GetByEndpointID(
	ctx context.Context,
	endpointID uint,
) ([]models.Incident, error) {

	var incidents []models.Incident

	err := r.db.WithContext(ctx).
		Where("endpoint_id = ?", endpointID).
		Find(&incidents).
		Error

	return incidents, err
}

func (r *IncidentRepository) GetAllIncidents(ctx context.Context) ([]models.Incident, error) {

	var incidents []models.Incident

	err := r.db.WithContext(ctx).Order("started_at DESC").Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (r *IncidentRepository) GetIncidentByID(ctx context.Context, id uint) (*models.Incident, error) {

	var incident models.Incident

	err := r.db.WithContext(ctx).Where("ID = ?", id).First(&incident).Error

	if err != nil {
		return nil, err
	}

	return &incident, nil
}

func (r *IncidentRepository) GetActiveIncidents(ctx context.Context) (
	[]models.Incident,
	error,
) {

	var incidents []models.Incident

	err := r.db.WithContext(ctx).Where("is_resolved = ?", false).Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, err
}

func (r *IncidentRepository) GetRecentIncidents(ctx context.Context) (
	[]models.Incident,
	error,
) {

	var incidents []models.Incident

	err := r.db.WithContext(ctx).Order("started_at DESC").Limit(10).Find(&incidents).Error

	if err != nil {
		return nil, err
	}

	return incidents, nil
}

package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type EndpointRepositoryInterface interface {
	Create(
		endpoint *models.Endpoint,
	) error
	GetAllEndpoints() ([]models.Endpoint, error)
	GetAllByUserID(
		userID uint,
	) ([]models.Endpoint, error)
	GetByID(
		id string,
		userID uint,
	) (*models.Endpoint, error)
	Update(
		endpoint *models.Endpoint,
	) error
	Delete(
		endpoint *models.Endpoint,
	) error
}

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *EndpointRepository {
	return &EndpointRepository{
		db: db,
	}
}

func (r *EndpointRepository) Create(
	endpoint *models.Endpoint,
) error {
	return r.db.Create(endpoint).Error
}

func (r *EndpointRepository) GetAllEndpoints() ([]models.Endpoint, error) {
	var endpoints []models.Endpoint

	err := r.db.Find(&endpoints).Error

	return endpoints, err
}

func (r *EndpointRepository) GetAllByUserID(
	userID uint,
) ([]models.Endpoint, error) {

	var endpoints []models.Endpoint

	err := r.db.
		Where("user_id = ?", userID).
		Find(&endpoints).
		Error

	return endpoints, err
}

func (r *EndpointRepository) GetByID(
	id string,
	userID uint,
) (*models.Endpoint, error) {

	var endpoint models.Endpoint

	err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		First(&endpoint).
		Error

	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}

func (r *EndpointRepository) Update(
	endpoint *models.Endpoint,
) error {
	return r.db.Save(endpoint).Error
}

func (r *EndpointRepository) Delete(
	endpoint *models.Endpoint,
) error {
	return r.db.Delete(endpoint).Error
}

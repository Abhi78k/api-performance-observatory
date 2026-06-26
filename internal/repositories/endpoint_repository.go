package repositories

import (
	"context"

	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type EndpointRepositoryInterface interface {
	Create(
		ctx context.Context,
		endpoint *models.Endpoint,
	) error
	GetAllEndpoints(ctx context.Context) ([]models.Endpoint, error)
	GetAllByUserID(
		ctx context.Context,
		userID uint,
	) ([]models.Endpoint, error)
	GetByID(
		ctx context.Context,
		id string,
		userID uint,
	) (*models.Endpoint, error)
	Update(
		ctx context.Context,
		endpoint *models.Endpoint,
	) error
	Delete(
		ctx context.Context,
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
	ctx context.Context,
	endpoint *models.Endpoint,
) error {
	return r.db.WithContext(ctx).Create(endpoint).Error
}

func (r *EndpointRepository) GetAllEndpoints(ctx context.Context) ([]models.Endpoint, error) {
	var endpoints []models.Endpoint

	err := r.db.WithContext(ctx).Find(&endpoints).Error

	return endpoints, err
}

func (r *EndpointRepository) GetAllByUserID(
	ctx context.Context,
	userID uint,
) ([]models.Endpoint, error) {

	var endpoints []models.Endpoint

	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&endpoints).
		Error

	return endpoints, err
}

func (r *EndpointRepository) GetByID(
	ctx context.Context,
	id string,
	userID uint,
) (*models.Endpoint, error) {

	var endpoint models.Endpoint

	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&endpoint).
		Error

	if err != nil {
		return nil, err
	}

	return &endpoint, nil
}

func (r *EndpointRepository) Update(
	ctx context.Context,
	endpoint *models.Endpoint,
) error {
	return r.db.WithContext(ctx).Save(endpoint).Error
}

func (r *EndpointRepository) Delete(
	ctx context.Context,
	endpoint *models.Endpoint,
) error {
	return r.db.WithContext(ctx).Delete(endpoint).Error
}

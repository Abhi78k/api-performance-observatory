package repositories

import (
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(
	user *models.User,
) error {

	err := r.db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(
	email string,
) (*models.User, error) {

	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(
	id uint,
) (*models.User, error) {

	var user models.User

	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

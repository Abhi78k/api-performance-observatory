package services

import (
	"errors"

	"github.com/Abhi78k/api-performance-observatory/internal/apperrors"
	"github.com/Abhi78k/api-performance-observatory/internal/auth"
	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	cfg      *config.Config
}

func NewAuthService(cfg *config.Config, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *AuthService) Register(
	email string,
	password string,
) error {
	_, err := s.userRepo.GetUserByEmail(email)

	if err == nil {
		return apperrors.ErrUserAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := auth.HashedPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.userRepo.CreateUser(&user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, error) {

	user, err := s.userRepo.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	err = auth.CheckPassword(
		user.Password,
		password,
	)

	if err != nil {
		return "", apperrors.ErrInvalidCredentials
	}

	accessToken, err := auth.GenerateAccessToken(s.cfg, user.ID)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) GetUserByID(
	userID uint,
) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *AuthService) GetMe(
	userID uint,
) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

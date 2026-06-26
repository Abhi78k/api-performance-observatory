package services

import (
	"context"
	"errors"

	"github.com/Abhi78k/api-performance-observatory/internal/apperrors"
	"github.com/Abhi78k/api-performance-observatory/internal/auth"
	"github.com/Abhi78k/api-performance-observatory/internal/config"
	"github.com/Abhi78k/api-performance-observatory/internal/logger"
	"github.com/Abhi78k/api-performance-observatory/internal/models"
	"github.com/Abhi78k/api-performance-observatory/internal/repositories"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo repositories.UserRepositoryInterface
	cfg      *config.Config
}

func NewAuthService(cfg *config.Config, userRepo repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	email string,
	password string,
) error {
	_, err := s.userRepo.GetUserByEmail(ctx, email)

	if err == nil {
		logger.Warn(
			"Registration failed",
			"email", email,
			"reason", "user already exists",
		)

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

	err = s.userRepo.CreateUser(ctx, &user)

	if err != nil {
		return err
	}

	logger.Info(
		"User registered",
		"user_id", user.ID,
		"email", user.Email,
	)

	return nil
}

func (s *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	user, err := s.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	err = auth.CheckPassword(
		user.Password,
		password,
	)

	if err != nil {
		logger.Warn(
			"Login failed",
			"email", email,
			"reason", "invalid credentials",
		)

		return "", apperrors.ErrInvalidCredentials
	}

	accessToken, err := auth.GenerateAccessToken(s.cfg, user.ID)

	if err != nil {
		return "", err
	}

	logger.Info(
		"User logged in",
		"user_id", user.ID,
	)

	return accessToken, nil
}

func (s *AuthService) GetUserByID(
	ctx context.Context,
	userID uint,
) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *AuthService) GetMe(
	ctx context.Context,
	userID uint,
) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

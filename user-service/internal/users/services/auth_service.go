package services

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
	"user-service/internal/users/models"
	"user-service/internal/users/repositories"
	"user-service/pkg/utils"
)

type AuthService struct {
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
	db        *gorm.DB
}

func NewAuthService(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository, db *gorm.DB) AuthService {
	return AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		db:        db,
	}
}

func (s *AuthService) Register(username string, email string, password string) (*string, error) {
	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to process password")
	}

	var jwtToken string
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Сохраняем пользователя в базу данных
		user := &models.User{
			Username: username,
			Email:    email,
			Password: hashedPassword,
		}

		if err := s.userRepo.CreateUser(user); err != nil {
			return status.Error(codes.Internal, "failed to save user")
		}

		// Генерация токена
		jwtToken, err = utils.GenerateToken(user.ID.String())
		if err != nil {
			return status.Error(codes.Internal, "failed to generate token")
		}

		token := &models.Token{
			UserID:    user.ID,
			Token:     jwtToken,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}

		if err := s.tokenRepo.CreateToken(token); err != nil {
			return status.Error(codes.Internal, "failed to save token")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	var jwtToken string

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Ищем пользователя в базе данных
		user, err := s.userRepo.FindByEmail(email)
		if err != nil {
			return status.Error(codes.NotFound, "user not found")
		}

		// Проверяем пароль
		if !utils.CheckPasswordHash(password, user.Password) {
			return status.Error(codes.Unauthenticated, "invalid credentials")
		}

		// Генерация токена
		jwtToken, err = utils.GenerateToken(user.ID.String())
		if err != nil {
			return status.Error(codes.Internal, "failed to generate token")
		}

		token := &models.Token{
			UserID:    user.ID,
			Token:     jwtToken,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}

		if err := s.tokenRepo.CreateToken(token); err != nil {
			return status.Error(codes.Internal, "failed to save token")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func (s *AuthService) ValidateToken(token string) (*string, bool, error) {
	// Ищем токен в базе
	tokenModel, err := s.tokenRepo.FindTokenByValue(token)
	if err != nil {
		return nil, false, err
	}

	// Проверяем срок действия токена
	if tokenModel.ExpiresAt.Before(time.Now()) {
		return nil, false, nil
	}

	userID := tokenModel.UserID.String()

	return &userID, true, nil
}

func (s *AuthService) GetUserInfo(token string) (*models.User, error) {
	// Ищем токен в базе
	tokenModel, err := s.tokenRepo.FindTokenByValue(token)
	if err != nil {
		return nil, status.Error(codes.NotFound, "token not found")
	}

	// Ищем пользователя, связанного с токеном
	user := tokenModel.User
	if user.ID == uuid.Nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &tokenModel.User, nil
}
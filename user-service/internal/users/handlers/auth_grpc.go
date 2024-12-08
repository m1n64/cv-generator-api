package handlers

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"user-service/internal/users/grpc/auth"
	"user-service/internal/users/models"
	"user-service/internal/users/repositories"
	"user-service/pkg/utils"
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
	userRepo  repositories.UserRepository
	tokenRepo repositories.TokenRepository
}

// Конструктор для AuthServiceServer
func NewAuthServiceServer(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) *AuthServiceServer {
	return &AuthServiceServer{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}

// Регистрация пользователя
func (s *AuthServiceServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.TokenResponse, error) {
	// Проверка данных
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username, email, and password are required")
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "failed to process password")
	}

	// Сохраняем пользователя в базу данных
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		log.Printf("Error saving user: %v", err)
		return nil, status.Error(codes.Internal, "failed to save user")
	}

	// Генерация токена
	jwtToken, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	token := &models.Token{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := s.tokenRepo.CreateToken(token); err != nil {
		log.Printf("Error saving token: %v", err)
		return nil, status.Error(codes.Internal, "failed to save token")
	}

	return &auth.TokenResponse{
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

// Логин пользователя
func (s *AuthServiceServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.TokenResponse, error) {
	// Проверка данных
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	// Ищем пользователя в базе данных
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		log.Printf("User not found: %v", err)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Проверяем пароль
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// Генерация токена
	jwtToken, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	token := &models.Token{
		UserID:    user.ID,
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := s.tokenRepo.CreateToken(token); err != nil {
		log.Printf("Error saving token: %v", err)
		return nil, status.Error(codes.Internal, "failed to save token")
	}

	return &auth.TokenResponse{
		Token:     jwtToken,
		ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

// Метод ValidateToken
func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	// Ищем токен в базе
	token, err := s.tokenRepo.FindTokenByValue(req.Token)
	if err != nil {
		log.Printf("Token not found: %v", err)
		return &auth.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	// Проверяем срок действия токена
	if token.ExpiresAt.Before(time.Now()) {
		return &auth.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &auth.ValidateTokenResponse{
		UserId: token.UserID.String(),
		Valid:  true,
	}, nil
}

// Метод GetUserInfo
func (s *AuthServiceServer) GetUserInfo(ctx context.Context, req *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	// Ищем токен в базе
	token, err := s.tokenRepo.FindTokenByValue(req.Token)
	if err != nil {
		log.Printf("Token not found: %v", err)
		return nil, status.Error(codes.NotFound, "token not found")
	}

	// Ищем пользователя, связанного с токеном
	user := token.User
	if user.ID == uuid.Nil {
		log.Printf("User not found for token: %v", req.Token)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &auth.GetUserInfoResponse{
		UserId:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

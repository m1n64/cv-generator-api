package handlers

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user-service/internal/users/grpc/auth"
	"user-service/internal/users/services"
	"user-service/pkg/utils"
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
	authService *services.AuthService
}

func NewAuthServiceServer(authService *services.AuthService) *AuthServiceServer {
	return &AuthServiceServer{
		authService: authService,
	}
}

func (s *AuthServiceServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.TokenResponse, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "username, email, and password are required")
	}

	token, err := s.authService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error registering user: %v", err))
		return nil, err
	}

	return &auth.TokenResponse{
		Token:     *token,
		ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.TokenResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "email and password are required")
	}

	token, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error logging in user: %v", err))
		return nil, err
	}

	return &auth.TokenResponse{
		Token:     *token,
		ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}, nil
}

func (s *AuthServiceServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	userID, valid, err := s.authService.ValidateToken(req.Token)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error validating token: %v", err))
		return nil, err
	}

	var userIDString string
	if userID != nil {
		userIDString = *userID
	}

	return &auth.ValidateTokenResponse{
		UserId: userIDString,
		Valid:  valid,
	}, nil
}

func (s *AuthServiceServer) GetUserInfo(ctx context.Context, req *auth.GetUserInfoRequest) (*auth.GetUserInfoResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	user, err := s.authService.GetUserInfo(req.Token)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error getting user info: %v", err))
		return nil, err
	}

	return &auth.GetUserInfoResponse{
		UserId:   user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *AuthServiceServer) Logout(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	if req.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "token is required")
	}

	err := s.authService.Logout(req.Token)
	if err != nil {
		utils.GetLogger().Info(fmt.Sprintf("Error logging out user: %v", err))
		return nil, err
	}

	return &auth.ValidateTokenResponse{
		UserId: "",
		Valid:  true,
	}, nil
}

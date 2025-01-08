package handlers

import (
	"context"
	"gateway-service/internal/auth/services"
	"gateway-service/internal/users/grpc/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type UserProxyHandler struct {
	authClient auth.AuthServiceClient
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type UserResponse struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUserProxy() *UserProxyHandler {
	authConn := services.GetAuthConnection()

	return &UserProxyHandler{
		authClient: auth.NewAuthServiceClient(authConn),
	}
}

func (h *UserProxyHandler) RegisterHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &auth.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.Register(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getAuthResponse(resp))
}

func (h *UserProxyHandler) LoginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &auth.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.Login(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getAuthResponse(resp))
}

func (h *UserProxyHandler) UserInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	grpcReq := &auth.GetUserInfoRequest{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.GetUserInfo(ctx, grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getUserResponse(resp))
}

func (h *UserProxyHandler) getAuthResponse(resp *auth.TokenResponse) AuthResponse {
	return AuthResponse{
		Token:     resp.Token,
		ExpiresAt: resp.ExpiresAt,
	}
}

func (h *UserProxyHandler) getUserResponse(resp *auth.GetUserInfoResponse) UserResponse {
	return UserResponse{
		UserId:   resp.UserId,
		Email:    resp.Email,
		Username: resp.Username,
	}
}

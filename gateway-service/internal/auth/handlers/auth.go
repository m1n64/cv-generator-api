package handlers

import (
	"context"
	"encoding/base32"
	"gateway-service/internal/users/grpc/auth"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type UserProxyHandler struct {
	authClient   auth.AuthServiceClient
	aesEncryptor *utils.AESEncryptor
}

type AuthResponse struct {
	Token     string `json:"token"`
	WsToken   string `json:"ws_token"`
	ExpiresAt string `json:"expires_at"`
}

type UserResponse struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUserProxy(client auth.AuthServiceClient, aesEncryptor *utils.AESEncryptor) *UserProxyHandler {
	return &UserProxyHandler{
		authClient:   client,
		aesEncryptor: aesEncryptor,
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
	token := h.getToken(c)

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

func (h *UserProxyHandler) LogoutHandler(c *gin.Context) {
	token := h.getToken(c)

	request := &auth.ValidateTokenRequest{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.Logout(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Valid,
	})
}

func (h *UserProxyHandler) getToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func (h *UserProxyHandler) getAuthResponse(resp *auth.TokenResponse) AuthResponse {
	key := h.aesEncryptor.GetKey()
	iv, _ := h.aesEncryptor.GenerateIV()

	wsToken, _ := h.aesEncryptor.Encrypt(resp.Token, key, iv)

	safeWsToken := base32.StdEncoding.EncodeToString([]byte(wsToken))

	return AuthResponse{
		Token:     resp.Token,
		WsToken:   safeWsToken,
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

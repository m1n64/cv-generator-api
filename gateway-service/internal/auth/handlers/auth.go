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

	c.JSON(http.StatusOK, gin.H{
		"token":      resp.Token,
		"expires_at": resp.ExpiresAt,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"token":      resp.Token,
		"expires_at": resp.ExpiresAt,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"user_id":  resp.UserId,
		"username": resp.Username,
		"email":    resp.Email,
	})
}

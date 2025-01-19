package middlewares

import (
	"gateway-service/internal/auth/services"
	"gateway-service/internal/users/grpc/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authClient auth.AuthServiceClient
}

func NewAuthMiddleware(client auth.AuthServiceClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: client,
	}
}

func (m *AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
			c.Abort()
			return
		}

		ctx := services.GetAuthContextWithToken()

		resp, err := m.authClient.ValidateToken(ctx, &auth.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			log.Printf("Error validating token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		if !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", resp.UserId)

		c.Next()
	}
}

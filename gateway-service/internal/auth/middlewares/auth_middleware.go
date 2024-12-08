package middlewares

import (
	"context"
	"gateway-service/internal/auth/services"
	"gateway-service/internal/users/grpc/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	authClient auth.AuthServiceClient
}

func NewAuthMiddleware() *AuthMiddleware {
	authConn := services.GetAuthConnection()

	return &AuthMiddleware{
		authClient: auth.NewAuthServiceClient(authConn),
	}
}

// Middleware для проверки токена
func (m *AuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		// Проверяем, начинается ли заголовок с "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
			c.Abort()
			return
		}

		// Вызываем gRPC для проверки токена
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		resp, err := m.authClient.ValidateToken(ctx, &auth.ValidateTokenRequest{
			Token: token,
		})
		if err != nil {
			log.Printf("Error validating token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Если токен недействителен
		if !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Сохраняем user_id в контекст
		c.Set("user_id", resp.UserId)

		// Продолжаем выполнение запроса
		c.Next()
	}
}

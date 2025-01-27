package handlers

import (
	"context"
	"encoding/base32"
	"gateway-service/internal/users/grpc/auth"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketPrivateHandler(manager *utils.WebSocketPrivateManager, authClient auth.AuthServiceClient, aesEncryptor *utils.AESEncryptor) gin.HandlerFunc {
	return func(c *gin.Context) {
		encryptedToken := c.Query("token")
		if encryptedToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
			return
		}

		urlDecodedToken, err := base32.StdEncoding.DecodeString(encryptedToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		key := aesEncryptor.GetKey()
		iv, encodedString := aesEncryptor.GetIVAndCipher(string(urlDecodedToken))

		token, err := aesEncryptor.Decrypt(encodedString, key, iv)
		if err != nil {
			utils.GetLogger().Error("Error decrypting token", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := authClient.ValidateToken(context.Background(), &auth.ValidateTokenRequest{Token: token})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !response.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		manager.AddClient(conn, response.UserId)

		defer manager.RemoveClient(conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}
}

package handlers

import (
	"context"
	"gateway-service/internal/users/grpc/auth"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketPrivateHandler(manager *utils.WebSocketPrivateManager, authClient auth.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
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

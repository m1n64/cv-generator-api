package middlewares

import (
	"context"
	"gateway-service/internal/cv/grpc/cv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type CVMiddleware struct {
	cvClient cv.CVServiceClient
}

func NewCVMiddleware(cvClient cv.CVServiceClient) *CVMiddleware {
	return &CVMiddleware{
		cvClient: cvClient,
	}
}

func (m *CVMiddleware) GetCVOriginalID() gin.HandlerFunc {
	return func(c *gin.Context) {
		cvID := c.Param("cv_id")
		if cvID == "" || uuid.Validate(cvID) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cv_id is required"})
			c.Abort()
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in context"})
			c.Abort()
			return
		}

		userIDStr, ok := userID.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id format"})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		resp, err := m.cvClient.GetOriginalID(ctx, &cv.GetOriginalIDRequest{
			CvId:   cvID,
			UserId: userIDStr,
		})

		if err != nil {
			log.Printf("Error getting original ID: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch original CV ID"})
			c.Abort()
			return
		}

		c.Set("original_cv_id", resp.Id)

		c.Next()
	}
}

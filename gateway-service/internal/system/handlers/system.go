package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"time":    time.Now(),
	})
}

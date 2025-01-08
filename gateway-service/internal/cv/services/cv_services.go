package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCvID(c *gin.Context) (string, error) {
	cvID, exist := c.Get("original_cv_id")
	if !exist {
		return "", fmt.Errorf("cv_id not found")
	}
	return cvID.(string), nil
}

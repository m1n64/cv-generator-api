package handlers

import (
	"context"
	"gateway-service/internal/cv/grpc/cv"
	services2 "gateway-service/internal/cv/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CVProxyHandler struct {
	cvClient cv.CVServiceClient
}

func NewCVProxy() *CVProxyHandler {
	cvConn := services2.GetCVConnection()

	return &CVProxyHandler{
		cvClient: cv.NewCVServiceClient(cvConn),
	}
}

func (h *CVProxyHandler) CreateCVHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in context"})
		return
	}

	resp, err := h.cvClient.CreateCV(ctx, &cv.CreateCVRequest{
		UserId: userID.(string),
		Name:   req.Name,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         resp.ExternalId,
		"name":       resp.Name,
		"created_at": resp.CreatedAt,
	})
}

func (h *CVProxyHandler) GetCVListHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in context"})
		return
	}

	listOfCV, err := h.cvClient.GetAllCVsByUserID(ctx, &cv.GetAllCVsByUserIDRequest{
		UserId: userID.(string),
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var cvList []map[string]interface{}
	for _, cvItem := range listOfCV.CvList {
		cvList = append(cvList, map[string]interface{}{
			"id":         cvItem.ExternalId,
			"name":       cvItem.Name,
			"created_at": cvItem.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"list": cvList,
	})
}

func (h *CVProxyHandler) GetCVHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	originalID, exists := c.Get("original_cv_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "original_cv_id not found in context"})
		return
	}

	originalIDStr, ok := originalID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid original_cv_id format"})
		return
	}

	resp, err := h.cvClient.GetCVByID(ctx, &cv.GetCVByIDRequest{
		CvId: originalIDStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         resp.ExternalId,
		"name":       resp.Name,
		"created_at": resp.CreatedAt,
	})
}

func (h *CVProxyHandler) UpdateCVHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	originalID, exists := c.Get("original_cv_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "original_cv_id not found in context"})
		return
	}

	originalIDStr, ok := originalID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid original_cv_id format"})
		return
	}

	resp, err := h.cvClient.UpdateCV(ctx, &cv.UpdateCVRequest{
		CvId: originalIDStr,
		Name: req.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         resp.ExternalId,
		"name":       resp.Name,
		"created_at": resp.CreatedAt,
	})
}

func (h *CVProxyHandler) DeleteCV(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	originalID, exists := c.Get("original_cv_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "original_cv_id not found in context"})
		return
	}

	originalIDStr, ok := originalID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid original_cv_id format"})
		return
	}

	resp, err := h.cvClient.DeleteCVByID(ctx, &cv.DeleteCVByIDRequest{
		CvId: originalIDStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
	})
}

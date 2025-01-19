package handlers

import (
	"gateway-service/internal/cv/grpc/cv"
	"gateway-service/internal/cv/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CVProxyHandler struct {
	cvClient cv.CVServiceClient
}

type CVResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func NewCVProxy(cvClient cv.CVServiceClient) *CVProxyHandler {
	return &CVProxyHandler{
		cvClient: cvClient,
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

	ctx := services.GetCvContextWithToken()

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

	c.JSON(http.StatusOK, h.getCVResponse(resp))
}

func (h *CVProxyHandler) GetCVListHandler(c *gin.Context) {
	ctx := services.GetCvContextWithToken()

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

	var cvList []CVResponse
	for _, cvItem := range listOfCV.CvList {
		cvList = append(cvList, CVResponse{
			ID:        cvItem.ExternalId,
			Name:      cvItem.Name,
			CreatedAt: cvItem.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"list": cvList,
	})
}

func (h *CVProxyHandler) GetCVHandler(c *gin.Context) {
	ctx := services.GetCvContextWithToken()

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

	c.JSON(http.StatusOK, h.getCVResponse(resp))
}

func (h *CVProxyHandler) UpdateCVHandler(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := services.GetCvContextWithToken()

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

	c.JSON(http.StatusOK, h.getCVResponse(resp))
}

func (h *CVProxyHandler) DeleteCV(c *gin.Context) {
	ctx := services.GetCvContextWithToken()

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

func (h *CVProxyHandler) getCVResponse(gRPCResponse *cv.CVResponse) CVResponse {
	return CVResponse{
		ID:        gRPCResponse.ExternalId,
		Name:      gRPCResponse.Name,
		CreatedAt: gRPCResponse.CreatedAt,
	}
}

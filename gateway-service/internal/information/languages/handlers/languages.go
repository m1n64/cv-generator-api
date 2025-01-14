package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	languages "gateway-service/internal/information/languages/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LanguagesProxyHandler struct {
	languageClient languages.LanguagesServiceClient
}

type LanguageRequest struct {
	Name  string `form:"name" binding:"required"`
	Level string `form:"level" binding:"required"`
}

type LanguageResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewLanguagesProxy(langsClient languages.LanguagesServiceClient) *LanguagesProxyHandler {
	return &LanguagesProxyHandler{
		languageClient: langsClient,
	}
}

func (h *LanguagesProxyHandler) GetLanguages(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.languageClient.GetLanguages(ctx, &languages.GetLanguagesRequest{
		CvId: cvID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var languagesResp []*LanguageResponse
	for _, language := range response.Languages {
		languagesResp = append(languagesResp, h.getLanguagesResponse(language))
	}

	c.JSON(http.StatusOK, gin.H{
		"languages": languagesResp,
	})
}

func (h *LanguagesProxyHandler) GetLanguage(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	languageID := c.Param("id")
	if languageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language_id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.languageClient.GetLanguageByID(ctx, &languages.GetLanguageByIDRequest{
		CvId: cvID,
		Id:   languageID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getLanguagesResponse(response))
}

func (h *LanguagesProxyHandler) CreateLanguage(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var request LanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.languageClient.CreateLanguage(ctx, &languages.CreateLanguageRequest{
		CvId:  cvID,
		Name:  request.Name,
		Level: request.Level,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getLanguagesResponse(response))
}

func (h *LanguagesProxyHandler) UpdateLanguage(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	languageID := c.Param("id")
	if languageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language_id is required"})
		return
	}

	var request LanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.languageClient.UpdateLanguageByID(ctx, &languages.UpdateLanguageByIDRequest{
		CvId:  cvID,
		Id:    languageID,
		Name:  request.Name,
		Level: request.Level,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getLanguagesResponse(response))
}

func (h *LanguagesProxyHandler) DeleteLanguage(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	languageID := c.Param("id")
	if languageID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "language_id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.languageClient.DeleteLanguageByID(ctx, &languages.DeleteLanguageByIDRequest{
		CvId: cvID,
		Id:   languageID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": response.Success,
	})
}

func (h *LanguagesProxyHandler) getLanguagesResponse(response *languages.LanguageResponse) *LanguageResponse {
	return &LanguageResponse{
		ID:        response.Id,
		Name:      response.Name,
		Level:     response.Level,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
}

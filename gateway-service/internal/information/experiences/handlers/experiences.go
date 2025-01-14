package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	experiences "gateway-service/internal/information/experiences/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ExperienceProxyHandler struct {
	experiencesClient experiences.ExperiencesServiceClient
}

type ExperienceRequest struct {
	Company     string  `json:"company" binding:"required"`
	Position    string  `json:"position" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type ExperienceResponse struct {
	ID          string  `json:"id"`
	Company     string  `json:"company"`
	Position    string  `json:"position"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func NewExperienceProxy(experiencesClient experiences.ExperiencesServiceClient) *ExperienceProxyHandler {
	return &ExperienceProxyHandler{
		experiencesClient: experiencesClient,
	}
}

func (h *ExperienceProxyHandler) GetExperiences(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.experiencesClient.GetExperiences(ctx, &experiences.GetExperiencesRequest{
		CvId: cvID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var experienceResponses []ExperienceResponse
	for _, experience := range response.WorkExperiences {
		experienceResponses = append(experienceResponses, h.getExperienceResponse(experience))
	}

	c.JSON(http.StatusOK, gin.H{"experiences": experienceResponses})
}

func (h *ExperienceProxyHandler) GetExperience(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experienceID := c.Param("id")
	if experienceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.experiencesClient.GetExperienceByID(ctx, &experiences.GetExperienceByIDRequest{
		CvId: cvID,
		Id:   experienceID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getExperienceResponse(response))
}

func (h *ExperienceProxyHandler) CreateExperience(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var experienceRequest ExperienceRequest
	if err := c.ShouldBindJSON(&experienceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.experiencesClient.CreateExperience(ctx, &experiences.CreateExperienceRequest{
		CvId:        cvID,
		Company:     experienceRequest.Company,
		Position:    experienceRequest.Position,
		StartDate:   experienceRequest.StartDate,
		EndDate:     experienceRequest.EndDate,
		Location:    experienceRequest.Location,
		Description: experienceRequest.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getExperienceResponse(response))
}

func (h *ExperienceProxyHandler) UpdateExperience(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experienceID := c.Param("id")
	if experienceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var request ExperienceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.experiencesClient.UpdateExperienceByID(ctx, &experiences.UpdateExperienceByIDRequest{
		CvId:        cvID,
		Id:          experienceID,
		Company:     request.Company,
		Position:    request.Position,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Location:    request.Location,
		Description: request.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getExperienceResponse(response))
}

func (h *ExperienceProxyHandler) DeleteExperience(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	experienceID := c.Param("id")
	if experienceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.experiencesClient.DeleteExperienceByID(ctx, &experiences.DeleteExperienceByIDRequest{
		Id:   experienceID,
		CvId: cvID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
	})
}

func (h *ExperienceProxyHandler) getExperienceResponse(gRPCResponse *experiences.ExperienceResponse) ExperienceResponse {
	return ExperienceResponse{
		ID:          gRPCResponse.Id,
		Company:     gRPCResponse.Company,
		Position:    gRPCResponse.Position,
		StartDate:   gRPCResponse.StartDate,
		EndDate:     gRPCResponse.EndDate,
		Location:    gRPCResponse.Location,
		Description: gRPCResponse.Description,
		CreatedAt:   gRPCResponse.CreatedAt,
		UpdatedAt:   gRPCResponse.UpdatedAt,
	}
}

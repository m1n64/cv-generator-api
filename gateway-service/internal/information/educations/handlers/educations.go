package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	educations "gateway-service/internal/information/educations/grpc"
	"gateway-service/internal/information/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EducationsProxyHandler struct {
	educationsClient educations.EducationServiceClient
}

type EducationRequest struct {
	Institute   string  `json:"institute" binding:"required"`
	Location    string  `json:"location" binding:"required"`
	Faculty     string  `json:"faculty" binding:"required"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
	Degree      *string `json:"degree"`
}

type EducationResponse struct {
	ID          string  `json:"id"`
	Institute   string  `json:"institute"`
	Location    string  `json:"location"`
	Faculty     string  `json:"faculty"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
	Degree      *string `json:"degree"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func NewEducationsProxy() *EducationsProxyHandler {
	educationsConn := services.GetInformationConnection()

	return &EducationsProxyHandler{
		educationsClient: educations.NewEducationServiceClient(educationsConn),
	}
}

func (h *EducationsProxyHandler) GetEducations(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.educationsClient.GetEducations(ctx, &educations.GetEducationsRequest{CvId: cvID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var eduResp []*EducationResponse
	for _, education := range response.Educations {
		eduResp = append(eduResp, h.getEducationResponse(education))
	}

	c.JSON(http.StatusOK, gin.H{
		"educations": eduResp,
	})
}

func (h *EducationsProxyHandler) GetEducation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	educationID := c.Param("id")
	if educationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.educationsClient.GetEducationByID(ctx, &educations.GetEducationByIDRequest{
		CvId: cvID,
		Id:   educationID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getEducationResponse(response))
}

func (h *EducationsProxyHandler) CreateEducation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req EducationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.educationsClient.CreateEducation(ctx, &educations.CreateEducationRequest{
		CvId:        cvID,
		Institution: req.Institute,
		Location:    req.Location,
		Faculty:     req.Faculty,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
		Degree:      req.Degree,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getEducationResponse(response))
}

func (h *EducationsProxyHandler) UpdateEducation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	educationID := c.Param("id")
	if educationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req EducationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.educationsClient.UpdateEducationByID(ctx, &educations.UpdateEducationByIDRequest{
		CvId:        cvID,
		Id:          educationID,
		Institution: req.Institute,
		Location:    req.Location,
		Faculty:     req.Faculty,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Description: req.Description,
		Degree:      req.Degree,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getEducationResponse(response))
}

func (h *EducationsProxyHandler) DeleteEducation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	educationID := c.Param("id")
	if educationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.educationsClient.DeleteEducationByID(ctx, &educations.DeleteEducationByIDRequest{
		CvId: cvID,
		Id:   educationID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

func (h *EducationsProxyHandler) getEducationResponse(gRPCResponse *educations.EducationResponse) *EducationResponse {
	return &EducationResponse{
		ID:          gRPCResponse.Id,
		Institute:   gRPCResponse.Institution,
		Location:    gRPCResponse.Location,
		Faculty:     gRPCResponse.Faculty,
		StartDate:   gRPCResponse.StartDate,
		EndDate:     gRPCResponse.EndDate,
		Description: gRPCResponse.Description,
		Degree:      gRPCResponse.Degree,
		CreatedAt:   gRPCResponse.CreatedAt,
		UpdatedAt:   gRPCResponse.UpdatedAt,
	}
}

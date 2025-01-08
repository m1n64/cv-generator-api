package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	"gateway-service/internal/information/services"
	skills "gateway-service/internal/information/skills/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type SkillsProxyHandler struct {
	skillClient skills.SkillsServiceClient
}

type SkillRequest struct {
	Name string `json:"name"`
}

type SkillResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewSkillsProxy() *SkillsProxyHandler {
	skillConn := services.GetInformationConnection()

	return &SkillsProxyHandler{
		skillClient: skills.NewSkillsServiceClient(skillConn),
	}
}

func (h *SkillsProxyHandler) GetSkills(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.skillClient.GetSkills(ctx, &skills.GetSkillsRequest{
		CvId: cvID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var skillsResp []SkillResponse
	for _, skill := range response.Skills {
		skillsResp = append(skillsResp, h.getSkillResponse(skill))
	}

	c.JSON(http.StatusOK, gin.H{
		"skills": skillsResp,
	})
}

func (h *SkillsProxyHandler) GetSkill(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	skillID := c.Param("id")
	if skillID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.skillClient.GetSkillByID(ctx, &skills.GetSkillByIDRequest{
		CvId: cvID,
		Id:   skillID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getSkillResponse(response))
}

func (h *SkillsProxyHandler) CreateSkill(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var request SkillRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.skillClient.CreateSkill(ctx, &skills.CreateSkillRequest{
		CvId: cvID,
		Name: request.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getSkillResponse(response))
}

func (h *SkillsProxyHandler) UpdateSkill(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	skillID := c.Param("id")
	if skillID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var request SkillRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.skillClient.UpdateSkillByID(ctx, &skills.UpdateSkillByIDRequest{
		CvId: cvID,
		Id:   skillID,
		Name: request.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getSkillResponse(response))
}

func (h *SkillsProxyHandler) DeleteSkill(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	skillID := c.Param("id")
	if skillID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.skillClient.DeleteSkillByID(ctx, &skills.DeleteSkillByIDRequest{
		CvId: cvID,
		Id:   skillID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": response.Success,
	})
}

func (h *SkillsProxyHandler) getSkillResponse(gRPCResponse *skills.SkillResponse) SkillResponse {
	return SkillResponse{
		ID:        gRPCResponse.Id,
		Name:      gRPCResponse.Name,
		CreatedAt: gRPCResponse.CreatedAt,
		UpdatedAt: gRPCResponse.UpdatedAt,
	}
}

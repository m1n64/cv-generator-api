package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	information "gateway-service/internal/information/information/grpc"
	"gateway-service/internal/information/services"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

type InformationProxyHandler struct {
	informationClient information.InformationServiceClient
}

type CVInformationRequest struct {
	FullName  string                `form:"full_name" binding:"required"`
	Position  string                `form:"position"`
	Location  string                `form:"location"`
	Biography string                `form:"biography"`
	Photo     *multipart.FileHeader `form:"photo"`
}

type CVInformationResponse struct {
	ID        string  `json:"id"`
	FullName  string  `json:"full_name"`
	Photo     []byte  `json:"photo"`
	PhotoUrl  *string `json:"photo_url"`
	Position  *string `json:"position"`
	Location  *string `json:"location"`
	Biography *string `json:"biography"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewInformationProxy() *InformationProxyHandler {
	informationConn := services.GetInformationConnection()

	return &InformationProxyHandler{
		informationClient: information.NewInformationServiceClient(informationConn),
	}
}

func (h *InformationProxyHandler) GetCVInformation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.informationClient.GetInformationByCvID(ctx, &information.GetInformationByCvIDRequest{
		CvId: cvID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getCVInfoResponse(response))
}

func (h *InformationProxyHandler) CreateOrUpdateCVInformation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var request CVInformationRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var photoBytes []byte
	if request.Photo != nil {
		file, err := request.Photo.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read photo"})
			return
		}
		defer file.Close()

		photoBytes, err = io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read photo"})
			return
		}
	}

	response, err := h.informationClient.CreateOrUpdateInformation(ctx, &information.CreateOrUpdateInformationRequest{
		CvId:      cvID,
		FullName:  request.FullName,
		Photo:     photoBytes,
		Position:  &request.Position,
		Location:  &request.Location,
		Biography: &request.Biography,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getCVInfoResponse(response))
}

func (h *InformationProxyHandler) DeleteCVInformation(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.informationClient.DeleteInformationByCvID(ctx, &information.DeleteInformationByCvIDRequest{
		CvId: cvID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": response.Success,
	})
}

func (h *InformationProxyHandler) getCVInfoResponse(gRPCResponse *information.InformationResponse) *CVInformationResponse {
	return &CVInformationResponse{
		ID:        gRPCResponse.Id,
		FullName:  gRPCResponse.FullName,
		Photo:     gRPCResponse.PhotoFile,
		PhotoUrl:  utils.ChangeDomainFromMinio(gRPCResponse.PhotoUrl),
		Position:  gRPCResponse.Position,
		Location:  gRPCResponse.Location,
		Biography: gRPCResponse.Biography,
		CreatedAt: gRPCResponse.CreatedAt,
		UpdatedAt: gRPCResponse.UpdatedAt,
	}
}

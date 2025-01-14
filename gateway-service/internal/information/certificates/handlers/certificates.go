package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	certificates "gateway-service/internal/information/certificates/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CertificatesProxyHandler struct {
	certificateClient certificates.CertificatesServiceClient
}

type CertificateRequest struct {
	Title       string  `json:"title"`
	Vendor      string  `json:"vendor"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
}

type CertificateResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Vendor      string  `json:"vendor"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func NewCertificatesProxy(certificatesClient certificates.CertificatesServiceClient) *CertificatesProxyHandler {
	return &CertificatesProxyHandler{
		certificateClient: certificatesClient,
	}
}

func (h *CertificatesProxyHandler) GetCertificates(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.certificateClient.GetCertificates(ctx, &certificates.GetCertificatesRequest{
		CvId: cvID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var certificatesResp []CertificateResponse

	for _, certificate := range response.Certificates {
		certificatesResp = append(certificatesResp, *h.getCertificateResponse(certificate))
	}

	c.JSON(http.StatusOK, gin.H{"certificates": certificatesResp})
}

func (h *CertificatesProxyHandler) GetCertificate(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	certificateId := c.Param("id")
	if certificateId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.certificateClient.GetCertificateByID(ctx, &certificates.GetCertificateByIDRequest{
		CvId: cvID,
		Id:   certificateId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getCertificateResponse(response))
}

func (h *CertificatesProxyHandler) CreateCertificate(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var request CertificateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.certificateClient.CreateCertificate(ctx, &certificates.CreateCertificateRequest{
		CvId:        cvID,
		Title:       request.Title,
		Vendor:      request.Vendor,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getCertificateResponse(response))
}

func (h *CertificatesProxyHandler) UpdateCertificate(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	certificateId := c.Param("id")
	if certificateId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var request CertificateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.certificateClient.UpdateCertificateByID(ctx, &certificates.UpdateCertificateByIDRequest{
		CvId:        cvID,
		Id:          certificateId,
		Title:       request.Title,
		Vendor:      request.Vendor,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getCertificateResponse(response))
}

func (h *CertificatesProxyHandler) DeleteCertificate(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	certificateId := c.Param("id")
	if certificateId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.certificateClient.DeleteCertificateByID(ctx, &certificates.DeleteCertificateByIDRequest{
		CvId: cvID,
		Id:   certificateId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

func (h *CertificatesProxyHandler) getCertificateResponse(gRPCResponse *certificates.CertificateResponse) *CertificateResponse {
	return &CertificateResponse{
		ID:          gRPCResponse.Id,
		Title:       gRPCResponse.Title,
		Vendor:      gRPCResponse.Vendor,
		StartDate:   gRPCResponse.StartDate,
		EndDate:     gRPCResponse.EndDate,
		Description: gRPCResponse.Description,
		CreatedAt:   gRPCResponse.CreatedAt,
		UpdatedAt:   gRPCResponse.UpdatedAt,
	}
}

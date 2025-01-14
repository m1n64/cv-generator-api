package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	generator "gateway-service/internal/generator/grpc"
	"gateway-service/internal/generator/services"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GeneratorProxyHandler struct {
	generatorClient generator.GeneratorServiceClient
}

type GeneratedPdfResponse struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	PdfFile   []byte  `json:"pdf_file"`
	PdfUrl    *string `json:"pdf_url"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type GeneratedLinkResponse struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	PdfFile []byte  `json:"pdf_file"`
	Link    *string `json:"link"`
}

func NewGeneratorProxy() *GeneratorProxyHandler {
	generatorConnection := services.GetGeneratorConnection()

	return &GeneratorProxyHandler{
		generatorClient: generator.NewGeneratorServiceClient(generatorConnection),
	}
}

func (h *GeneratorProxyHandler) GetAllGeneratedPdfs(c *gin.Context) {
	userId, ok := h.getUserID(c)
	if !ok {
		return
	}

	ctx, cancel := h.createContext()
	defer cancel()

	response, err := h.generatorClient.GetAllListGenerated(ctx, &generator.AllListGeneratedRequest{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var generatedPdfs []*GeneratedPdfResponse
	for _, gRPCResponse := range response.Pdfs {
		generatedPdfs = append(generatedPdfs, h.getGeneratedPdfResponse(gRPCResponse))
	}

	c.JSON(http.StatusOK, gin.H{"generated_pdfs": generatedPdfs})
}

func (h *GeneratorProxyHandler) GetGeneratedPdfByCV(c *gin.Context) {
	cvID, ok := h.getCvID(c)
	if !ok {
		return
	}

	userId, ok := h.getUserID(c)
	if !ok {
		return
	}

	ctx, cancel := h.createContext()
	defer cancel()

	response, err := h.generatorClient.GetListGenerated(ctx, &generator.GeneratedRequest{CvId: cvID, UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var generatedPdfs []*GeneratedPdfResponse
	for _, gRPCResponse := range response.Pdfs {
		generatedPdfs = append(generatedPdfs, h.getGeneratedPdfResponse(gRPCResponse))
	}

	c.JSON(http.StatusOK, gin.H{"generated_pdfs": generatedPdfs})
}

func (h *GeneratorProxyHandler) GeneratePdf(c *gin.Context) {
	cvID, ok := h.getCvID(c)
	if !ok {
		return
	}

	userId, ok := h.getUserID(c)
	if !ok {
		return
	}

	genId := c.Param("id")
	if genId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	ctx, cancel := h.createContext()
	defer cancel()

	response, err := h.generatorClient.GetGeneratedPDF(ctx, &generator.GeneratedPDFRequest{CvId: cvID, UserId: userId, Id: genId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getGeneratedPdfResponse(response))
}

func (h *GeneratorProxyHandler) DeleteGeneratedPdf(c *gin.Context) {
	cvID, ok := h.getCvID(c)
	if !ok {
		return
	}

	userId, ok := h.getUserID(c)
	if !ok {
		return
	}

	genId := c.Param("id")
	if genId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	ctx, cancel := h.createContext()
	defer cancel()

	resp, err := h.generatorClient.DeleteGenerated(ctx, &generator.GeneratedPDFRequest{CvId: cvID, UserId: userId, Id: genId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

func (h *GeneratorProxyHandler) DownloadGeneratedPdf(c *gin.Context) {
	cvID, ok := h.getCvID(c)
	if !ok {
		return
	}

	userId, ok := h.getUserID(c)
	if !ok {
		return
	}

	genId := c.Param("id")
	if genId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	ctx, cancel := h.createContext()
	defer cancel()

	resp, err := h.generatorClient.GetPDFLink(ctx, &generator.GeneratedPDFRequest{CvId: cvID, UserId: userId, Id: genId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &GeneratedLinkResponse{
		ID:      resp.Id,
		Title:   resp.Title,
		PdfFile: resp.PdfFile,
		Link:    resp.PdfUrl,
	})
}

func (h *GeneratorProxyHandler) createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *GeneratorProxyHandler) getUserID(c *gin.Context) (string, bool) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return "", false
	}
	return userId.(string), true
}

func (h *GeneratorProxyHandler) getCvID(c *gin.Context) (string, bool) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", false
	}
	return cvID, true
}

func (h *GeneratorProxyHandler) getGeneratedPdfResponse(gRPCResponse *generator.GeneratedPdf) *GeneratedPdfResponse {
	return &GeneratedPdfResponse{
		ID:        gRPCResponse.Id,
		Title:     gRPCResponse.Title,
		PdfFile:   gRPCResponse.PdfFile,
		PdfUrl:    utils.ChangeDomainFromMinio(gRPCResponse.PdfUrl),
		Status:    gRPCResponse.Status,
		CreatedAt: gRPCResponse.CreatedAt,
		UpdatedAt: gRPCResponse.UpdatedAt,
	}
}

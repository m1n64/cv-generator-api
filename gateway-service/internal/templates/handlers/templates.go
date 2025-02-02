package handlers

import (
	"context"
	templates "gateway-service/internal/templates/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TemplateProxyHandler struct {
	templateClient templates.TemplateServiceClient
}

type TemplateResponse struct {
	Id       string `json:"id"`
	Template string `json:"html_template"`
	Title    string `json:"title"`
}

func NewTemplateProxyHandler(templateClient templates.TemplateServiceClient) *TemplateProxyHandler {
	return &TemplateProxyHandler{
		templateClient: templateClient,
	}
}

func (h *TemplateProxyHandler) GetDefaultTemplate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.templateClient.GetDefaultTemplate(ctx, &templates.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"html_template": resp.Template,
	})
}

func (h *TemplateProxyHandler) GetColors(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.templateClient.GetColorScheme(ctx, &templates.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"colors": resp.Colors,
	})
}

func (h *TemplateProxyHandler) GetTemplates(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.templateClient.GetTemplates(ctx, &templates.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var templatesResp []*TemplateResponse
	for _, template := range resp.Templates {
		templatesResp = append(templatesResp, &TemplateResponse{
			Id:       template.Id,
			Template: template.Template,
			Title:    template.Title,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"templates": templatesResp,
	})
}

func (h *TemplateProxyHandler) GetTemplate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.templateClient.GetTemplateById(ctx, &templates.TemplateByIdRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &TemplateResponse{
		Id:       resp.Id,
		Template: resp.Template,
		Title:    resp.Title,
	})
}

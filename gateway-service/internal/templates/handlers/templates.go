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

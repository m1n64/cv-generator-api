package handlers

import (
	"context"
	"errors"
	"gateway-service/internal/ai/enums"
	ai "gateway-service/internal/ai/grpc"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type AiProxyHandler struct {
	aiClient ai.AiServiceClient
}

type GenerateRequest struct {
	Prompt string `json:"prompt" binding:"required"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Response string `json:"response"`
}

func NewAiProxy(aiClient ai.AiServiceClient) *AiProxyHandler {
	return &AiProxyHandler{
		aiClient: aiClient,
	}
}

func (h *AiProxyHandler) Generate(c *gin.Context) {
	generateType := c.Param("type")

	if !enums.IsValidGenerateType(generateType) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	var request GenerateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	key, _ := enums.GetKeyByGenerateType(generateType)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if request.Stream {
		h.handleStream(ctx, c, key, request.Prompt)
		return
	}

	h.handleRequest(ctx, c, key, request.Prompt)
}

func (h *AiProxyHandler) handleRequest(ctx context.Context, c *gin.Context, serviceId, prompt string) {
	response, err := h.aiClient.Generate(ctx, &ai.GenerateRequest{
		Prompt:    prompt,
		ServiceId: serviceId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GenerateResponse{
		Response: response.Response,
	})
}

func (h *AiProxyHandler) handleStream(ctx context.Context, c *gin.Context, serviceId, prompt string) {
	response, err := h.aiClient.StreamGenerate(ctx, &ai.GenerateRequest{
		Prompt:    prompt,
		ServiceId: serviceId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "streaming unsupported",
		})
		return
	}

	c.Stream(func(w io.Writer) bool {
		for {
			msg, err := response.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return false
				}
				c.SSEvent("error", gin.H{"message": "stream error"})
				flusher.Flush()
				return false
			}

			c.SSEvent("message", GenerateResponse{
				Response: msg.GetResponse(),
			})
			flusher.Flush()
		}
	})
}

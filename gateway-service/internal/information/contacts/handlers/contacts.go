package handlers

import (
	"context"
	services2 "gateway-service/internal/cv/services"
	contacts "gateway-service/internal/information/contacts/grpc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ContactsProxyHandler struct {
	contactClient contacts.ContactsServiceClient
}

type ContactRequest struct {
	Title string  `json:"title" binding:"required"`
	Link  *string `json:"link"`
}

type ContactResponse struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Link      *string `json:"link"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func NewContactsProxy(contactsClient contacts.ContactsServiceClient) *ContactsProxyHandler {
	return &ContactsProxyHandler{
		contactClient: contactsClient,
	}
}

func (h *ContactsProxyHandler) GetContacts(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.contactClient.GetContacts(ctx, &contacts.GetContactsRequest{
		CvId: cvID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var contactsResp []*ContactResponse

	for _, contact := range response.Contacts {
		contactsResp = append(contactsResp, h.getContactResponse(contact))
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contactsResp,
	})
}

func (h *ContactsProxyHandler) GetContact(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.contactClient.GetContactByID(ctx, &contacts.GetContactByIDRequest{
		CvId: cvID,
		Id:   contactID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getContactResponse(response))
}

func (h *ContactsProxyHandler) CreateContact(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.contactClient.CreateContact(ctx, &contacts.CreateContactRequest{
		CvId:  cvID,
		Title: req.Title,
		Link:  req.Link,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getContactResponse(response))
}

func (h *ContactsProxyHandler) UpdateContact(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := h.contactClient.UpdateContactByID(ctx, &contacts.UpdateContactByIDRequest{
		CvId:  cvID,
		Id:    contactID,
		Title: req.Title,
		Link:  req.Link,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.getContactResponse(response))
}

func (h *ContactsProxyHandler) DeleteContact(c *gin.Context) {
	cvID, err := services2.GetCvID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.contactClient.DeleteContactByID(ctx, &contacts.DeleteContactByIDRequest{
		CvId: cvID,
		Id:   contactID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": resp.Success})
}

func (h *ContactsProxyHandler) getContactResponse(gRPCResponse *contacts.ContactResponse) *ContactResponse {
	return &ContactResponse{
		ID:        gRPCResponse.Id,
		Title:     gRPCResponse.Title,
		Link:      gRPCResponse.Link,
		CreatedAt: gRPCResponse.CreatedAt,
		UpdatedAt: gRPCResponse.UpdatedAt,
	}
}

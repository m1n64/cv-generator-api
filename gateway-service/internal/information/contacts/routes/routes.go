package routes

import (
	"gateway-service/internal/auth/middlewares"
	middlewares2 "gateway-service/internal/cv/middlewares"
	"gateway-service/internal/information/contacts/handlers"
	"github.com/gin-gonic/gin"
)

func CVContactsRoutes(r *gin.Engine, authMiddleware *middlewares.AuthMiddleware, cvMiddleware *middlewares2.CVMiddleware) {
	contactsHandler := handlers.NewContactsProxy()

	contactsGroup := r.Group("/information/contacts/:cv_id")
	contactsGroup.Use(authMiddleware.ValidateToken())
	contactsGroup.Use(cvMiddleware.GetCVOriginalID())
	contactsGroup.GET("/", contactsHandler.GetContacts)
	contactsGroup.POST("/", contactsHandler.CreateContact)

	contactsIdGroup := contactsGroup.Group("/:id")
	contactsIdGroup.GET("/", contactsHandler.GetContact)
	contactsIdGroup.POST("/", contactsHandler.UpdateContact)
	contactsIdGroup.DELETE("/", contactsHandler.DeleteContact)
}
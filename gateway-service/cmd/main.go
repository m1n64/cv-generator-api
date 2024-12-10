package main

import (
	"fmt"
	"gateway-service/internal/auth/handlers"
	"gateway-service/internal/auth/middlewares"
	handlers3 "gateway-service/internal/cv/handlers"
	middlewares2 "gateway-service/internal/cv/middlewares"
	handlers2 "gateway-service/internal/system/handlers"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	fmt.Println("Gateway service started!")

	utils.InitLogs()
	utils.LoadEnv()

	authHandler := handlers.NewUserProxy()
	authMiddleware := middlewares.NewAuthMiddleware()

	r := gin.Default()

	r.GET("/ping", handlers2.PingHandler)

	r.GET("/documentation/openapi.json", func(c *gin.Context) {
		c.File("./config/swagger/openapi.json")
	})

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.POST("/login", authHandler.LoginHandler)

	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware.ValidateToken())
	userGroup.GET("/info", authHandler.UserInfoHandler)

	cvHandler := handlers3.NewCVProxy()
	cvMiddleware := middlewares2.NewCVMiddleware()

	cvGroup := r.Group("/cv")
	cvGroup.Use(authMiddleware.ValidateToken())
	cvGroup.POST("/", cvHandler.CreateCVHandler)
	cvGroup.GET("/", cvHandler.GetCVListHandler)

	cvWithIDMiddleware := cvGroup.Group("/:cv_id")
	cvWithIDMiddleware.Use(cvMiddleware.GetCVOriginalID())
	cvWithIDMiddleware.GET("/", cvHandler.GetCVHandler)
	cvWithIDMiddleware.POST("/", cvHandler.UpdateCVHandler)
	cvWithIDMiddleware.DELETE("/", cvHandler.DeleteCV)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	fmt.Println("Gateway service run successfully!")
}

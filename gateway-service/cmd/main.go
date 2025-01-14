package main

import (
	"fmt"
	"gateway-service/internal/auth/middlewares"
	"gateway-service/internal/auth/routes"
	middlewares2 "gateway-service/internal/cv/middlewares"
	routes2 "gateway-service/internal/cv/routes"
	routes10 "gateway-service/internal/generator/routes"
	routes6 "gateway-service/internal/information/certificates/routes"
	routes7 "gateway-service/internal/information/contacts/routes"
	routes8 "gateway-service/internal/information/educations/routes"
	routes9 "gateway-service/internal/information/experiences/routes"
	routes3 "gateway-service/internal/information/information/routes"
	routes4 "gateway-service/internal/information/languages/routes"
	routes5 "gateway-service/internal/information/skills/routes"
	"gateway-service/internal/system/consumers"
	handlers2 "gateway-service/internal/system/handlers"
	middlewares3 "gateway-service/internal/system/middlewares"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Gateway service started!")

	utils.InitLogs()
	utils.LoadEnv()

	utils.ConnectRabbitMQ()
	utils.InitializeQueues()

	logger := utils.GetLogger()

	r := gin.Default()

	webSocketManager := utils.NewWebSocketPrivateManager()
	r.GET("/ws/private", handlers2.WebSocketPrivateHandler(webSocketManager))

	go func() {
		err := utils.ListenToQueue(utils.GatewayEventsQueue, consumers.NewEventConsumer(logger, webSocketManager).Handle)
		if err != nil {
			logger.Error("Error listening to queue", zap.Error(err))
		}
	}()

	authMiddleware := middlewares.NewAuthMiddleware()
	cvMiddleware := middlewares2.NewCVMiddleware()

	r.GET("/ping", handlers2.PingHandler)

	r.LoadHTMLFiles("./config/asyncapi/output/index.html")
	r.Static("/ws-docs", "./config/asyncapi/output")

	openApiDoc := r.Group("/documentation")
	openApiDoc.Use(middlewares3.CORSMiddleware())
	openApiDoc.GET("/openapi.json", func(c *gin.Context) {
		c.File("./config/swagger/openapi.json")
	})
	openApiDoc.GET("/ws-docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./config/asyncapi/output/index.html", nil)
	})

	routes.AuthRoutes(r, authMiddleware)
	routes2.CVRoutes(r, authMiddleware, cvMiddleware)
	routes3.CVInfoRoutes(r, authMiddleware, cvMiddleware)
	routes4.CVLanguagesRoutes(r, authMiddleware, cvMiddleware)
	routes5.CVSkillsRoutes(r, authMiddleware, cvMiddleware)
	routes6.CVCertificatesRoutes(r, authMiddleware, cvMiddleware)
	routes7.CVContactsRoutes(r, authMiddleware, cvMiddleware)
	routes8.CVEducationsRoutes(r, authMiddleware, cvMiddleware)
	routes9.CVExperiencesRoutes(r, authMiddleware, cvMiddleware)
	routes2.CVGeneratorRoutes(r, authMiddleware, cvMiddleware)
	routes10.GeneratorRoutes(r, authMiddleware, cvMiddleware)

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	fmt.Println("Gateway service run successfully!")
}

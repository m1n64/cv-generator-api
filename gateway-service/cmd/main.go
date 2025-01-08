package main

import (
	"fmt"
	"gateway-service/internal/auth/middlewares"
	"gateway-service/internal/auth/routes"
	middlewares2 "gateway-service/internal/cv/middlewares"
	routes2 "gateway-service/internal/cv/routes"
	routes6 "gateway-service/internal/information/certificates/routes"
	routes7 "gateway-service/internal/information/contacts/routes"
	routes8 "gateway-service/internal/information/educations/routes"
	routes9 "gateway-service/internal/information/experiences/routes"
	routes3 "gateway-service/internal/information/information/routes"
	routes4 "gateway-service/internal/information/languages/routes"
	routes5 "gateway-service/internal/information/skills/routes"
	handlers2 "gateway-service/internal/system/handlers"
	middlewares3 "gateway-service/internal/system/middlewares"
	"gateway-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	fmt.Println("Gateway service started!")

	utils.InitLogs()
	utils.LoadEnv()

	authMiddleware := middlewares.NewAuthMiddleware()
	cvMiddleware := middlewares2.NewCVMiddleware()

	r := gin.Default()

	r.GET("/ping", handlers2.PingHandler)

	openApiDoc := r.Group("/documentation")
	openApiDoc.Use(middlewares3.CORSMiddleware())
	openApiDoc.GET("/openapi.json", func(c *gin.Context) {
		c.File("./config/swagger/openapi.json")
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

	r.Run(fmt.Sprintf(":%s", os.Getenv("SERVICE_PORT")))
	fmt.Println("Gateway service run successfully!")
}

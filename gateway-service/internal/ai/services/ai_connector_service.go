package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"os"
)

func GetAIConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("AI_SERVICE_HOST"), os.Getenv("AI_SERVICE_PORT")))
}

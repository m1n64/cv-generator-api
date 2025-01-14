package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"os"
)

func GetTemplatesConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("TEMPLATES_SERVICE_HOST"), os.Getenv("TEMPLATES_SERVICE_PORT")))
}

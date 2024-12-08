package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"os"
)

func GetAuthConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("AUTH_SERVICE_PORT")))
}

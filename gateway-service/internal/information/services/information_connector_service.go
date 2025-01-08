package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"os"
)

func GetInformationConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("INFORMATION_SERVICE_HOST"), os.Getenv("INFORMATION_SERVICE_PORT")))
}

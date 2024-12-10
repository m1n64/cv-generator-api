package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"os"
)

func GetCVConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("CV_SERVICE_HOST"), os.Getenv("CV_SERVICE_PORT")))
}

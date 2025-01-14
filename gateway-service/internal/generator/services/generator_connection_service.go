package services

import (
	"fmt"
	"gateway-service/pkg/utils"
	grpc "google.golang.org/grpc"
	"os"
)

func GetGeneratorConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("GENERATOR_SERVICE_HOST"), os.Getenv("GENERATOR_SERVICE_PORT")))
}

package services

import (
	"context"
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
)

func GetCVConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("CV_SERVICE_HOST"), os.Getenv("CV_SERVICE_PORT")))
}

func GetCvContextWithToken() context.Context {
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + os.Getenv("CV_GRPC_TOKEN"),
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}

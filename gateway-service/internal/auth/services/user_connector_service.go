package services

import (
	"context"
	"fmt"
	"gateway-service/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
)

func GetAuthConnection() *grpc.ClientConn {
	return utils.CreateGRPCClient(fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("AUTH_SERVICE_PORT")))
}

func GetAuthContextWithToken() context.Context {
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + os.Getenv("AUTH_GRPC_TOKEN"),
	})

	return metadata.NewOutgoingContext(context.Background(), md)
}

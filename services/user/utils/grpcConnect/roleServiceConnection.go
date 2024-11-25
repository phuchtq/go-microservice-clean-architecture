package grpcconnect

import (
	grpc_gateway "architecture_template/constants/grpcGateway"
	"architecture_template/constants/notis"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectRoleService() (*grpc.ClientConn, error) {
	cnn, err := grpc.NewClient(grpc_gateway.RPCRolePort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Print("Fail to connect role service: ", err)
		return nil, errors.New(notis.InternalErr)
	}

	return cnn, nil
}

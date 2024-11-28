package grpcconnect

import (
	grpc_gateway "architecture_template/constants/grpcGateway"
	"architecture_template/constants/notis"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectRoleService() (*grpc.ClientConn, error) {
	cnn, err := grpc.NewClient(grpc_gateway.RPCRolePort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(fmt.Sprintf(notis.GrpcConnectMsg, "Role") + err.Error())
		return nil, errors.New(notis.InternalErr)
	}

	return cnn, nil
}

package routes

import (
	"architecture_template/constants/notis"
	"architecture_template/protocols/roleService/pb"
	role_grpc "architecture_template/services/role/adapters/gRPC"
	envvar "architecture_template/services/role/constants/envVar"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	backUpGrpcPort string = "Your back up grpc port"
)

func InitializeGRPCRoute() {
	var logger = &log.Logger{}

	var service string = "Role"

	var port string = os.Getenv(envvar.GrpcPort)
	if port == "" {
		logger.Println(fmt.Sprintf(notis.GrpcPortEnvNotSetMsg, service))
		port = backUpApiPort
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		logger.Println(fmt.Sprintf(notis.NetListeningMsg, port) + err.Error())
		return
	}

	logger.Println(service+" service grpc starts listening on port ", port)

	rsServer, err := role_grpc.GenerateGRPCService()
	if err != nil {
		logger.Println(fmt.Sprintf(notis.GrpcGenerateMsg, service))
		return
	}

	var grpcServer = grpc.NewServer()

	pb.RegisterRoleServiceServer(grpcServer, rsServer)

	if err := grpcServer.Serve(l); err != nil {
		logger.Println(fmt.Sprintf(notis.GrpcServeMsg, service) + err.Error())
	}
}

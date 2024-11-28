package grpc

import (
	"architecture_template/constants/notis"
	"architecture_template/protocols/roleService/pb"
	"architecture_template/services/role/interfaces"
	businesslogics "architecture_template/services/role/usecases/businessLogics"
	"errors"
)

type grpcServer struct {
	service interfaces.IRoleService
	pb.UnimplementedRoleServiceServer
}

func GenerateGRPCService() (*grpcServer, error) {
	service, err := businesslogics.GenerateService()

	if err != nil {
		return nil, errors.New(notis.InternalErr)
	}

	return &grpcServer{
		service: service,
	}, nil
}

package grpc

import (
	"architecture_template/protocols/roleService/pb"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *grpcServer) GetAllRoles(c context.Context, empty *emptypb.Empty) (*pb.RolesResp, error) {
	roles, err := s.service.GetAllRoles(c)

	if err != nil {
		return nil, err
	}

	var res *pb.RolesResp
	var tmpStorage []*pb.Role

	for _, role := range *roles {
		tmpStorage = append(tmpStorage, &pb.Role{
			Id:   role.RoleId,
			Name: role.RoleName,
		})
	}

	res.Roles = tmpStorage
	return res, nil
}

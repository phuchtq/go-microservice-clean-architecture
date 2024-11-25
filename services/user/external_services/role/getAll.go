package role

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"architecture_template/protocols/roleService/pb"
	user_notis "architecture_template/services/user/constants/notis"
	redis_key "architecture_template/services/user/constants/redisKey"
	grpc_connect "architecture_template/services/user/utils/grpcConnect"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *rolService) GetRoleStorage(isFindId bool, c context.Context) map[string]string {
	var key string = redis_key.GetAllRolesBasedIdKey
	if !isFindId {
		key = redis_key.GetAllRolesBasedNameKey
	}

	if res := generateRoleDictionaryFromRedisCache(r.redisClient, key, r.logger, c); res != nil || len(res) > 0 {
		return res
	}

	return generateRoleDictionFromGrpc(r.redisClient, key, isFindId, r.logger, c)
}

func generateRoleDictionaryFromRedisCache(client *redis.Client, key string, logger *log.Logger, c context.Context) map[string]string {
	res, _, isValid := helper.GetDataFromRedis[map[string]string](client, key, c)

	if !isValid {
		logger.Print(user_notis.RedisMsg)
		return nil
	}

	return *res
}

func generateRoleDictionFromGrpc(client *redis.Client, key string, isFindId bool, logger *log.Logger, c context.Context) map[string]string {
	var res = make(map[string]string)

	var roles = fetchRoles(logger, c)

	for _, role := range roles {
		if !isFindId {
			res[role.Name] = role.Id
		} else {
			res[role.Id] = role.Name
		}
	}

	if helper.SaveDataToRedis(client, key, response.DataStorage{
		Data: res,
	}, c) != nil {
		logger.Print(notis.RedisMsg + helper.ConvertModelToString(res))
	}

	return res
}

func fetchRoles(logger *log.Logger, c context.Context) []*pb.Role {
	cnn, err := grpc_connect.ConnectRoleService()

	if err != nil {
		return getBackUpStaticRoles()
	}
	defer cnn.Close()

	var client = pb.NewRoleServiceClient(cnn)

	roles, err := client.GetAllRoles(c, &emptypb.Empty{})
	if err != nil {
		logger.Print(notis.RoleRpcMsg + err.Error())
		return getBackUpStaticRoles()
	}

	return roles.Roles
}

func getBackUpStaticRoles() []*pb.Role {
	return []*pb.Role{
		{
			Id:   "R001",
			Name: "Admin",
		},

		{
			Id:   "R002",
			Name: "Staff",
		},

		{
			Id:   "R003",
			Name: "Customer",
		},
	}
}

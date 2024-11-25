package externalservices

import (
	"architecture_template/services/user/external_services/role"
	external_services "architecture_template/services/user/interfaces/external_services"
	"log"

	"github.com/redis/go-redis/v9"
)

func GenerateExternalServices(redisClient *redis.Client, logger *log.Logger) external_services.IRole {
	return role.InitializeExternalRoleService(redisClient, logger)
}

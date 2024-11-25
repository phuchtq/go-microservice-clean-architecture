package role

import (
	external_services "architecture_template/services/user/interfaces/external_services"
	"log"

	"github.com/redis/go-redis/v9"
)

type rolService struct {
	redisClient *redis.Client
	logger      *log.Logger
}

func InitializeExternalRoleService(redisClient *redis.Client, logger *log.Logger) external_services.IRole {
	return &rolService{
		logger:      logger,
		redisClient: redisClient,
	}
}

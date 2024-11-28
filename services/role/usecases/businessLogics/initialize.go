package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/external_services/caches"
	envvar "architecture_template/services/role/constants/envVar"
	"architecture_template/services/role/infrastructures/db"
	"architecture_template/services/role/infrastructures/repositories"
	"architecture_template/services/role/interfaces"
	"fmt"
	"log"
	"os"
)

type service struct {
	roleRepo interfaces.IRepository
	logger   *log.Logger
}

const (
	backUpRedisPort string = "Your back up redis port"
)

// This func used by adapter layers - controllers or handlers to generate service in order to process client requests.
func GenerateService() (interfaces.IRoleService, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	var logger = &log.Logger{}

	var redisPort string = os.Getenv(envvar.RedisPort)
	if redisPort == "" {
		logger.Println(fmt.Sprintf(notis.RedisPortEnvNotSetMsg, "Role"))
		redisPort = backUpRedisPort
	}

	return InitializeService(repositories.InitializeRepository(db, logger, caches.InitializeRedisTrigger("localhost:"+redisPort).GetRedisClient()), logger), nil
}

// This func as a intermediary to generate service.
// It is seperated from GenerateService() as main purpose to support for generating testing object used to call testing methods.
func InitializeService(repo interfaces.IRepository, logger *log.Logger) interfaces.IRoleService {
	return &service{
		roleRepo: repo,
		logger:   logger,
	}
}

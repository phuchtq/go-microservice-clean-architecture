package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/external_services/caches"
	envvar "architecture_template/services/user/constants/envVar"
	"architecture_template/services/user/infrastructures/db"
	"architecture_template/services/user/infrastructures/repositories"
	"architecture_template/services/user/interfaces"
	"errors"
	"fmt"
	"log"
	"os"
)

type service struct {
	repo   interfaces.IRepository
	logger *log.Logger
}

const (
	backUpRedisPort string = "Your back up redis port"
)

// This func used by adapter layers - controllers or handlers to generate service in order to process client requests.
func GenerateService() (interfaces.IService, error) {
	db, err := db.ConnectDB()

	if err != nil {
		return nil, errors.New(notis.InternalErr)
	}

	var logger *log.Logger = &log.Logger{}

	var redisPort string = os.Getenv(envvar.RedisPort)
	if redisPort == "" {
		logger.Println(fmt.Sprintf(notis.RedisPortEnvNotSetMsg, "User"))
		redisPort = backUpRedisPort
	}

	return InitializeService(repositories.InitializeRepository(db, logger, caches.InitializeRedisTrigger("localhost:"+redisPort).GetRedisClient()), logger), nil
}

// This func as a intermediary to generate service.
// It is seperated from GenerateService() as main purpose to support for generating testing object used to call testing methods.
func InitializeService(repo interfaces.IRepository, logger *log.Logger) interfaces.IService {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

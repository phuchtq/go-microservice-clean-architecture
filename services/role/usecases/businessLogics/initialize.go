package businesslogics

import (
	"architecture_template/external_services/caches"
	envvar "architecture_template/services/role/constants/envVar"
	"architecture_template/services/role/infrastructures/db"
	"architecture_template/services/role/infrastructures/repositories"
	"architecture_template/services/role/interfaces"
	"log"
	"os"
)

type service struct {
	roleRepo interfaces.IRepository
	logger   *log.Logger
}

func GenerateService() (interfaces.IRoleService, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}

	var logger = &log.Logger{}

	return &service{
		roleRepo: repositories.InitializeRepository(db, logger, caches.InitializeRedisTrigger("localhost:"+os.Getenv(envvar.RedisPort)).GetRedisClient()),
		logger:   logger,
	}, nil
}

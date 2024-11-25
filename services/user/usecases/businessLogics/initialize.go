package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/external_services/caches"
	envvar "architecture_template/services/user/constants/envVar"
	"architecture_template/services/user/infrastructures/db"
	"architecture_template/services/user/infrastructures/repositories"
	"architecture_template/services/user/interfaces"
	"errors"
	"log"
	"os"
)

type service struct {
	repo   interfaces.IRepository
	logger *log.Logger
}

func GenerateService() (interfaces.IService, error) {
	db, err := db.ConnectDB()

	if err != nil {
		return nil, errors.New(notis.InternalErr)
	}

	var logger *log.Logger = &log.Logger{}

	return &service{
		repo:   repositories.InitializeRepository(db, logger, caches.InitializeRedisTrigger("localhost:"+os.Getenv(envvar.RedisPort)).GetRedisClient()),
		logger: logger,
	}, nil
}

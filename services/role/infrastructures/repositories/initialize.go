package repositories

import (
	"architecture_template/services/role/interfaces"
	"database/sql"
	"log"

	"github.com/redis/go-redis/v9"
)

// export ~~ public

type repo struct {
	db          *sql.DB
	logger      *log.Logger
	redisClient *redis.Client
}

func InitializeRepository(db *sql.DB, logger *log.Logger, redisClient *redis.Client) interfaces.IRepository {
	return &repo{
		db:          db,
		logger:      logger,
		redisClient: redisClient,
	}
}

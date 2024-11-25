package caches

import (
	"github.com/redis/go-redis/v9"
)

type redisTrigger struct {
	address string
}

func InitializeRedisTrigger(address string) *redisTrigger {
	return &redisTrigger{
		address: address,
	}
}

var redisClient *redis.Client

func (rd *redisTrigger) GetRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = initializeRedisClient(rd.address)
	return redisClient
}

func initializeRedisClient(address string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: address,
	})
}

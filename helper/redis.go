package helper

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func SaveDataToRedis(client *redis.Client, key string, data response.DataStorage, c context.Context) error {
	var expiration time.Duration = CacheDuration
	if data.Data == nil || data.ErrMsg != nil {
		expiration = NotExistCacheDuration
	}

	value, _ := ToJson(data)

	return client.Set(c, key, value, expiration).Err()
}

func GetDataFromRedis[T any](client *redis.Client, key string, c context.Context) (*T, error, bool) {
	// The boolean value is a flag to decide whether these data: T and error can be used

	dataStorage, err := extractDataFromRedis[response.DataStorage](client, key, c)

	if err != nil {
		// In this case, the error get from extractDataFromRedis is due to redis connection - system error
		// -> These data can't be used so the flag would be false in this case
		return nil, err, false
	}

	return dataStorage.Data.(*T), dataStorage.ErrMsg, true // Flag true as there is no error from the system, the error just from data itself
}

func RefreshRedisCache[T any](keys, messages []string, logger *log.Logger, client *redis.Client, c context.Context) {
	go processRefrechCache[[]T](keys[0], messages, logger, client, c)

	if len(keys) > 1 {
		go processRefrechCache[T](keys[1], messages, logger, client, c)
	}
}

func processRefrechCache[T any](key string, messages []string, logger *log.Logger, client *redis.Client, c context.Context) {
	data, _, isValid := GetDataFromRedis[T](client, key, c)

	if !isValid {
		logger.Println(messages[0])
	}

	if data != nil {
		if removeDataFromRedis(client, key, c) != nil {
			logger.Println(messages[1])
		}
	}
}

func removeDataFromRedis(client *redis.Client, key string, c context.Context) error {
	return client.Del(c, key).Err()
}

func extractDataFromRedis[T any](client *redis.Client, key string, c context.Context) (*T, error) {
	cachedData, err := client.Get(c, key).Result()

	if err == redis.Nil {
		return nil, errors.New(notis.UndefinedDataWarnMsg + " with keyword: " + key) // No matched keyword found
	}

	if err != nil {
		return nil, err
	}

	return ConvertJsonToModel[T](cachedData), nil
}

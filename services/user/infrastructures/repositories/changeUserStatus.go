package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	redis_key "architecture_template/services/user/constants/redisKey"
	"architecture_template/services/user/entities"
	"context"
	"errors"
	"fmt"
	"time"
)

func (tr *repo) ChangeUserStatus(status bool, id string, c context.Context) error {
	errLogMsg := user_notis.UserRepoMsg + "ChangeUserStatus - "
	lastFailValueQuery := "NULL"
	if status {
		lastFailValueQuery = fmt.Sprint(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))
	}

	query := "Update " + entities.GetTable() + " set activeStatus = " + fmt.Sprint(status) + ", failAccess = 0, lastFail = " + lastFailValueQuery + ", accessToken = NULL, refreshToken = NULL, actionPeriod = NULL, actionToken = NULL, isHaveToResetPw = NULL where id = ?"
	if _, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	// users, err := helper.GetDataFromRedis[[]entities.User](tr.redisCache, redis_key.GetAllKey, c)
	// user, _ := tr.GetUserById(id, c)

	// Refresh data in cache
	go helper.RefreshRedisCache[entities.User](
		[]string{ // keys
			redis_key.GetAllKey,
			id,
		},

		[]string{ // messages
			user_notis.RedisMsg,
			"",
		},

		tr.logger, // logger

		tr.redisCache, // redis client

		c, // context
	)

	tr.db.Close()
	return nil
}

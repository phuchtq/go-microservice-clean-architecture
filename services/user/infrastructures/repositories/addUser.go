package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	redis_key "architecture_template/services/user/constants/redisKey"
	"architecture_template/services/user/entities"
	"context"
	"errors"
)

func (tr *repo) AddUser(u entities.User, c context.Context) error {
	var errLogMsg string = user_notis.UserRepoMsg + "AddUser - "
	var query string = "Insert into " + entities.GetTable() + "(id, email, password, roleId, activeStatus, failAccess, lastFail) values (?, ?, ?, ?, ?, ?, ?)"

	if _, err := tr.db.Exec(query, u.UserId, u.Email, u.Pasword, u.RoleId, u.ActiveStatus, u.FailAccess, u.LastFail); err != nil {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	// Refresh data in cache
	helper.RefreshRedisCache[entities.User](
		[]string{ // keys
			redis_key.GetAllKey,
		},

		[]string{ // messages
			notis.RedisExtractDataMsg,
			notis.RedisRefreshKeyMsg,
		},

		tr.logger, // logger

		tr.redisCache, // redis client

		c, // context
	)

	tr.db.Close()
	return nil
}

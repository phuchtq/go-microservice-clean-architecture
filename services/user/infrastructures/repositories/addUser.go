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
	errLogMsg := user_notis.UserRepoMsg + "AddUser - "
	query := "Insert into " + entities.GetTable() + "(id, email, password, roleId, activeStatus, failAccess, lastFail) values (?, ?, ?, ?, ?, ?, ?)"

	if _, err := tr.db.Exec(query, u.UserId, u.Email, u.Pasword, u.RoleId, u.ActiveStatus, u.FailAccess, u.LastFail); err != nil {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	// Refresh data in cache
	go helper.RefreshRedisCache[entities.User](
		[]string{ // keys
			redis_key.GetAllKey,
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

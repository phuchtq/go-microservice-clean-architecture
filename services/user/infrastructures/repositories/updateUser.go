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
)

func (tr *repo) UpdateUser(u entities.User, c context.Context) error {
	var errLogMsg string = user_notis.UserRepoMsg + "UpdateUser - "
	var query string = "Update " + entities.GetTable() + " set email = ?, password = ?, roleId = ?, accessToken = ?, refreshToken = ?, activeStatus = ?, failAccess = ?, lastFail = ? where id = ?"

	res, err := tr.db.Exec(query, u.Email, u.Pasword, u.RoleId, u.AccessToken, u.RefreshToken, u.ActiveStatus, u.FailAccess, u.LastFail, u.UserId)
	if err != nil {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	if rowsAffected == 0 {
		tr.db.Close()
		return errors.New(user_notis.UndefinedUserWarnMsg)
	}

	// Refresh cache if exists
	go func() {
		helper.RefreshRedisCache[entities.User](
			[]string{ // keys
				redis_key.GetAllKey,
				fmt.Sprintf(redis_key.GetByIdKey, u.UserId),
			},

			[]string{ // messages
				user_notis.RedisMsg,
				"",
			},

			tr.logger, // logger

			tr.redisCache, // redis client

			c, // context
		)
	}()

	tr.db.Close()
	return nil
}

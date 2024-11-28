package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	redis_key "architecture_template/services/role/constants/redisKey"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"fmt"
)

func (tr *repo) UpdateRole(r entities.Role, c context.Context) error {
	var errLogMsg string = notis.RoleRepoMsg + "UpdateRole - "
	var query string = "Update " + entities.GetTable() + " set roleName = ?, activeStatus = ? where id = ?"

	if res, err := tr.db.Exec(query, r.RoleName, r.ActiveStatus, r.RoleId); err != nil {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)
		return errors.New(notis.InternalErr)
	} else {
		if rowsAffected, err := res.RowsAffected(); err != nil {
			tr.db.Close()
			tr.logger.Println(errLogMsg, err)
			return errors.New(notis.InternalErr)
		} else if rowsAffected == 0 {
			return errors.New(notis.UndefinedRoleWarnMsg)
		}
	}

	tr.db.Close()

	// Refresh cache if exists
	helper.RefreshRedisCache[entities.Role](
		[]string{ // keys
			redis_key.GetAllKey,
			fmt.Sprintf(redis_key.GetByIdKey, r.RoleId),
		},

		[]string{ // messages
			notis.RedisExtractDataMsg,
			notis.RedisRefreshKeyMsg,
		},

		tr.logger, // logger

		tr.redisClient, // redis client

		c, // context
	)

	return nil
}

package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	redis_key "architecture_template/services/role/constants/redisKey"

	"architecture_template/services/role/entities"
	"context"
	"errors"
)

func (tr *repo) CreateRole(r entities.Role, c context.Context) error {
	var errLogMsg string = notis.RoleRepoMsg + "CreateRole - "
	var query string = "Insert into " + entities.GetTable() + "(roleId, roleName, activeStatus) values (?, ?, ?)"

	if _, err := tr.db.Exec(query, r.RoleId, r.RoleName, r.ActiveStatus); err != nil {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	tr.db.Close()

	// Refresh cache if exists
	go helper.RefreshRedisCache[entities.Role](
		[]string{ // keys
			redis_key.GetAllKey,
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

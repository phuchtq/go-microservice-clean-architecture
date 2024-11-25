package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	role_notis "architecture_template/services/role/constants/notis"
	redis_key "architecture_template/services/role/constants/redisKey"

	"architecture_template/services/role/entities"
	"context"
	"errors"
	"log"
)

func (tr *repo) CreateRole(r entities.Role, c context.Context) error {
	errLogMsg := notis.RoleRepoMsg + "CreateRole - "
	query := "Insert into " + entities.GetTable() + "(roleId, roleName, activeStatus) values (?, ?, ?)"

	if _, err := tr.db.Exec(query, r.RoleId, r.RoleName, r.ActiveStatus); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	tr.db.Close()

	// Refresh cache if exists
	go helper.RefreshRedisCache[entities.Role](
		[]string{ // keys
			redis_key.GetAllKey,
		},

		[]string{ // messages
			role_notis.RedisMsg,
			"",
		},

		tr.logger, // logger

		tr.redisClient, // redis client

		c, // context
	)

	return nil
}

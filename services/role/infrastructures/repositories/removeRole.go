package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	role_notis "architecture_template/services/role/constants/notis"
	redis_key "architecture_template/services/role/constants/redisKey"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"fmt"

	"log"
)

func (tr *repo) RemoveRole(id string, c context.Context) error {
	var errLogMsg string = notis.RoleRepoMsg + "RemoveRole - "
	var query string = "Update " + entities.GetTable() + " set activeStatus = false where id = ?"

	if res, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	} else {
		if rowsAffected, err := res.RowsAffected(); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return errors.New(notis.InternalErr)
		} else if rowsAffected == 0 {
			return errors.New(notis.UndefinedRoleWarnMsg)
		}
	}

	tr.db.Close()

	// Refresh cache if exists
	go helper.RefreshRedisCache[entities.Role](
		[]string{ // keys
			redis_key.GetAllKey,
			fmt.Sprintf(redis_key.GetByIdKey, id),
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

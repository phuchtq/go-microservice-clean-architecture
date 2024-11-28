package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"fmt"

	redis_key "architecture_template/services/role/constants/redisKey"
	"architecture_template/services/role/entities"
	"context"
	"errors"
)

func (tr *repo) ActivateRole(id string, c context.Context) error {
	var errLogMsg string = notis.RoleRepoMsg + "ActivateRole - "
	var query string = "Update " + entities.GetTable() + " set activeStatus = true where id = ?"

	if _, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

	// Refresh cache if exists
	helper.RefreshRedisCache[entities.Role](
		[]string{ // keys
			redis_key.GetAllKey,
			fmt.Sprintf(redis_key.GetByIdKey, id),
		},

		[]string{ // messages
			notis.RedisExtractDataMsg,
			notis.RedisRefreshKeyMsg,
		},

		tr.logger, // logger

		tr.redisClient, // redis client

		c, // context
	)

	tr.db.Close()
	return nil
}

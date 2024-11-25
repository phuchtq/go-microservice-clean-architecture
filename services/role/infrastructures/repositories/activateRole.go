package repositories

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"fmt"

	role_notis "architecture_template/services/role/constants/notis"
	redis_key "architecture_template/services/role/constants/redisKey"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"log"
)

func (tr *repo) ActivateRole(id string, c context.Context) error {
	errLogMsg := notis.RoleRepoMsg + "ActivateRole - "

	query := "Update " + entities.GetTable() + " set activeStatus = true where id = ?"
	if _, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}

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

	tr.db.Close()
	return nil
}

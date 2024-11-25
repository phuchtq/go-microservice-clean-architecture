package repositories

import (
	"architecture_template/common_dtos/response"
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

func (tr *repo) GetRolesByStatus(status bool, c context.Context) (*[]entities.Role, error) {
	var key string = fmt.Sprintf(redis_key.GetByStatusKey, status)

	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[[]entities.Role](tr.redisClient, key, c); isValid {
		return res, err
	} else {
		tr.logger.Print(role_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = notis.RoleRepoMsg + "GetRolesByStatus - "
	var query string = "Select * from " + entities.GetTable() + " where activeStatus = $1"
	var InternalErr error = errors.New(notis.InternalErr)

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)

		if helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
			ErrMsg: InternalErr,
		}, c) != nil {
			tr.logger.Print(notis.RedisMsg)
		}

		return nil, InternalErr
	}

	defer rows.Close()

	var res *[]entities.Role
	for rows.Next() {
		var x entities.Role

		if err := rows.Scan(&x.RoleId, &x, x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)

			if helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
				ErrMsg: InternalErr,
			}, c) != nil {
				tr.logger.Print(notis.RedisMsg)
			}

			return nil, InternalErr
		}

		*res = append(*res, x)
	}

	if err := helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
		Data: res,
	}, c); err != nil { // Save data to cache for next request
		tr.logger.Print(notis.RedisMsg + helper.ConvertModelToString(res))
	}

	return res, nil
}

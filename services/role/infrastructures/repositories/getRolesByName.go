package repositories

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"architecture_template/services/role/entities"
	"fmt"
	"strings"

	role_notis "architecture_template/services/role/constants/notis"
	redis_key "architecture_template/services/role/constants/redisKey"
	"context"
	"errors"
	"log"
)

func (tr *repo) GetRolesByName(name string, c context.Context) (*[]entities.Role, error) {
	var key string = fmt.Sprintf(redis_key.GetByNameKey, strings.ToLower(name))

	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[[]entities.Role](tr.redisClient, key, c); isValid {
		return res, err
	} else {
		tr.logger.Print(role_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	errLogMsg := notis.RoleRepoMsg + "GetRolesByName - "
	query := "Select * from " + entities.GetTable() + " where lower(roleName) like lower($1)"
	var internalErr error = errors.New(notis.InternalErr)
	rows, err := tr.db.Query(query, "%"+name+"%")

	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)

		if helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
			ErrMsg: internalErr,
		}, c) != nil {
			tr.logger.Print(notis.RedisMsg)
		}

		return nil, internalErr
	}
	defer rows.Close()

	var res *[]entities.Role
	for rows.Next() {
		var x entities.Role

		if err := rows.Scan(&x.RoleId, &x, x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)

			if helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
				ErrMsg: internalErr,
			}, c) != nil {
				tr.logger.Print(notis.RedisMsg)
			}

			return nil, internalErr
		}

		*res = append(*res, x)
	}

	if helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
		Data: res,
	}, c) != nil { // Save data to cache for next request
		tr.logger.Print(notis.RedisMsg + helper.ConvertModelToString(res))
	}

	return res, nil
}

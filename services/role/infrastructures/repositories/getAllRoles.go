package repositories

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	role_notis "architecture_template/services/role/constants/notis"
	"fmt"

	redis_key "architecture_template/services/role/constants/redisKey"

	"architecture_template/helper"
	"architecture_template/services/role/entities"
	"context"
	"errors"
)

func (tr *repo) GetAllRoles(c context.Context) (*[]entities.Role, error) {
	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[[]entities.Role](tr.redisClient, redis_key.GetAllKey, c); isValid {
		return res, err
	} else {
		tr.logger.Println(role_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = notis.RoleRepoMsg + "GetAllRoles - "
	var query string = "Select * from " + entities.GetTable()
	var internalErr error = errors.New(notis.InternalErr)
	var storeDataLogPrefixMsg string = fmt.Sprintf(notis.RedisStoreDataMsg, redis_key.GetAllKey)

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)

		if err := helper.SaveDataToRedis(tr.redisClient, redis_key.GetAllKey, response.DataStorage{
			ErrMsg: internalErr,
		}, c); err != nil { // Save data to cache for next request
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, internalErr
	}
	defer rows.Close()

	var res *[]entities.Role
	for rows.Next() {
		var x entities.Role

		if err := rows.Scan(&x.RoleId, &x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			tr.logger.Println(errLogMsg, err)

			if helper.SaveDataToRedis(tr.redisClient, redis_key.GetAllKey, response.DataStorage{
				ErrMsg: internalErr,
			}, c) != nil { // Save data to cache for next request
				tr.logger.Println(storeDataLogPrefixMsg + err.Error())
			}

			return nil, internalErr
		}

		*res = append(*res, x)
	}

	if helper.SaveDataToRedis(tr.redisClient, redis_key.GetAllKey, response.DataStorage{
		Data: res,
	}, c) != nil { // Save data to cache for next request
		tr.logger.Println(storeDataLogPrefixMsg + helper.ConvertModelToString(res))
	}

	return res, nil
}

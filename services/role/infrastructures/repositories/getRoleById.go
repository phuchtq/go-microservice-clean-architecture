package repositories

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	role_notis "architecture_template/services/role/constants/notis"
	redis_key "architecture_template/services/role/constants/redisKey"
	"fmt"

	"architecture_template/services/role/entities"
	"context"
	"database/sql"
	"errors"
)

func (tr *repo) GetRoleById(id string, c context.Context) (*entities.Role, error) {
	var key string = fmt.Sprintf(redis_key.GetByIdKey, id)
	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[entities.Role](tr.redisClient, key, c); isValid {
		return res, err
	} else {
		tr.logger.Println(role_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = notis.RoleRepoMsg + "GetRoleById - "
	var query string = "Select * from " + entities.GetTable() + " where roleId = $1"
	var storeDataLogPrefixMsg string = fmt.Sprintf(notis.RedisStoreDataMsg, key)
	var res *entities.Role

	if err := tr.db.QueryRow(query, id).Scan(&res.RoleId, &res.RoleName, &res.ActiveStatus); err != nil && err == sql.ErrNoRows {
		tr.db.Close()

		if err := helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{}, c); err != nil {
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, nil // No data found with incoming ID parameter - actually not considered as an error -> no data and error returned
	} else if err != nil && err != sql.ErrNoRows {
		var internalErr error = errors.New(notis.InternalErr)

		tr.db.Close()
		tr.logger.Println(errLogMsg, err) // Error but bot caused of None-data found - Return error

		if err := helper.SaveDataToRedis(tr.redisClient, key, response.DataStorage{
			ErrMsg: internalErr,
		}, c); err != nil {
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, internalErr
	}
	tr.db.Close()

	if err := helper.SaveDataToRedis(tr.redisClient, fmt.Sprintf(redis_key.GetByIdKey, id), response.DataStorage{}, c); err != nil { // Save data to cache for next request
		tr.logger.Println(storeDataLogPrefixMsg + helper.ConvertModelToString(res))
	}

	return res, nil
}

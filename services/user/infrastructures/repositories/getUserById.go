package repositories

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	redis_key "architecture_template/services/user/constants/redisKey"
	"architecture_template/services/user/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (tr *repo) GetUserById(id string, c context.Context) (*entities.User, error) {
	var key string = fmt.Sprintf(redis_key.GetByIdKey, id)

	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[entities.User](tr.redisCache, key, c); isValid {
		return res, err
	} else {
		tr.logger.Println(user_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = user_notis.UserRepoMsg + "GetUserById - "
	var query string = "Select id, email, password, roleId, activeStatus from " + entities.GetTable() + " where id = ?"
	var storeDataLogPrefixMsg string = fmt.Sprintf(notis.RedisStoreDataMsg, key)
	var res *entities.User

	if err := tr.db.QueryRow(query, id).Scan(&res.UserId, &res.Email, &res.Pasword, &res.RoleId, &res.ActiveStatus); err != nil && err == sql.ErrNoRows {
		tr.db.Close()

		if err := helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{}, c); err != nil {
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, nil // No data found with incoming ID parameter - actually not considered as an error -> no data and error returned
	} else if err != nil && err != sql.ErrNoRows {
		tr.db.Close()
		tr.logger.Println(errLogMsg, err)

		var internalErr error = errors.New(notis.InternalErr)

		if err := helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{
			ErrMsg: internalErr,
		}, c); err != nil {
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, internalErr
	}

	if helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{
		Data: res,
	}, c) != nil {
		tr.logger.Println(storeDataLogPrefixMsg + helper.ConvertModelToString(res))
	}

	tr.db.Close()

	return res, nil
}

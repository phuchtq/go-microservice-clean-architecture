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
	"strings"
)

func (tr *repo) GetUserByEmail(email string, c context.Context) (*entities.User, error) {
	var key string = fmt.Sprintf(redis_key.GetByEmailKey, strings.ToLower(email))

	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[entities.User](tr.redisCache, key, c); isValid {
		return res, err
	} else {
		tr.logger.Println(user_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = user_notis.UserRepoMsg + "GetUserByEmail - "
	var query string = "Select id, roleId, activeStatus, lastFail, failAccess, password from " + entities.GetTable() + " where lower(email) = lower($1)"
	var storeDataLogPrefixMsg string = fmt.Sprintf(notis.RedisStoreDataMsg, key)
	var res *entities.User

	if err := tr.db.QueryRow(query, email).Scan(&res.UserId, &res.RoleId, &res.ActiveStatus, &res.LastFail, &res.FailAccess, &res.Pasword); err != nil && err == sql.ErrNoRows {
		tr.db.Close()

		if err := helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{}, c); err != nil {
			tr.logger.Println(storeDataLogPrefixMsg + err.Error())
		}

		return nil, nil
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

	res.Email = email
	return res, nil
}

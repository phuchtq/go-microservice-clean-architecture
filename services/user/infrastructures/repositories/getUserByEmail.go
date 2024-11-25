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
		tr.logger.Print(user_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = user_notis.UserRepoMsg + "GetUserByEmail - "
	var query string = "Select id, roleId, activeStatus, lastFail, failAccess, password from " + entities.GetTable() + " where lower(email) = lower($1)"
	var res *entities.User

	if err := tr.db.QueryRow(query, email).Scan(&res.UserId, &res.RoleId, &res.ActiveStatus, &res.LastFail, &res.FailAccess, &res.Pasword); err != nil && err == sql.ErrNoRows {
		tr.db.Close()

		if helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{}, c) != nil {
			tr.logger.Print(notis.RedisMsg)
		}

		return nil, nil
	} else if err != nil && err != sql.ErrNoRows {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)

		var internalErr error = errors.New(notis.InternalErr)

		if helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{
			ErrMsg: internalErr,
		}, c) != nil {
			tr.logger.Print(notis.RedisMsg)
		}

		return nil, internalErr
	}

	go func() {
		if helper.SaveDataToRedis(tr.redisCache, key, response.DataStorage{
			Data: res,
		}, c) != nil {
			tr.logger.Print(notis.RedisMsg + helper.ConvertModelToString(res))
		}
	}()

	tr.db.Close()

	res.Email = email
	return res, nil
}

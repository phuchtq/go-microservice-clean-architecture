package repositories

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	redis_key "architecture_template/services/user/constants/redisKey"
	"architecture_template/services/user/entities"
	"context"
	"errors"
)

func (tr *repo) GetAllUsers(c context.Context) (*[]entities.User, error) {
	// Retrieve data from redis cache
	if res, err, isValid := helper.GetDataFromRedis[[]entities.User](tr.redisCache, redis_key.GetAllKey, c); isValid {
		return res, err
	} else {
		tr.logger.Print(user_notis.RedisMsg + err.Error()) // Fetching data from cache meets problem
	}
	//-------------------------------------------

	// Retrieve database
	var errLogMsg string = user_notis.UserRepoMsg + "GetAllUsers - "
	var query string = "Select id, email, password, roleId, activeStatus from " + entities.GetTable()

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		tr.logger.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res *[]entities.User
	for rows.Next() {
		var x entities.User
		if err := rows.Scan(&x.UserId, &x.Email, &x.Pasword, &x.RoleId, &x.ActiveStatus); err != nil {
			tr.db.Close()
			tr.logger.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}

		*res = append(*res, x)
	}

	go func() {
		if helper.SaveDataToRedis(tr.redisCache, redis_key.GetAllKey, response.DataStorage{
			Data: res,
		}, c) != nil { // Save data to cache for next request
			tr.logger.Print(notis.RedisMsg + helper.ConvertModelToString(res))
		}
	}()

	return res, nil
}

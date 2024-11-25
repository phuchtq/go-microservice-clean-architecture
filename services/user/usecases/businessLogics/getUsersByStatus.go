package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/user/entities"
	"context"
	"errors"
	"strconv"
)

func (tr *service) GetUsersByStatus(rawStatus string, c context.Context) (*[]entities.User, error) {
	status, err := strconv.ParseBool(rawStatus)

	if err != nil {
		return nil, errors.New(notis.InvalidStatusWarnMsg)
	}

	return tr.repo.GetUsersByStatus(status, c)
}

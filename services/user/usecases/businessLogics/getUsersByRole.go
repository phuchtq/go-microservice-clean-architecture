package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/user/entities"
	"context"
	"errors"
	"strings"
)

func (tr *service) GetUsersByRole(role string, c context.Context) (*[]entities.User, error) {
	if role == "" {
		return tr.repo.GetAllUsers(c)
	}

	role = strings.TrimSpace(role)
	res, err := tr.repo.GetUsersByRole(role, c)

	if res == nil && err == nil {
		return nil, errors.New(notis.UndefinedRoleWarnMsg)
	}

	return res, nil
}

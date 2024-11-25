package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"architecture_template/services/role/entities"
	"context"
	"errors"
)

func (tr *service) GetRolesByStatus(rawStatus string, c context.Context) (*[]entities.Role, error) {
	if rawStatus == "" {
		return tr.roleRepo.GetAllRoles(c)
	}

	status, err := helper.IsStatusValid(rawStatus)

	if err != nil {
		return nil, errors.New(notis.InvalidStatusWarnMsg)
	}

	return tr.roleRepo.GetRolesByStatus(status, c)
}

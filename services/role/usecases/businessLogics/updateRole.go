package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"strings"
)

func (tr *service) UpdateRole(x entities.Role, c context.Context) error {
	res, err := tr.roleRepo.GetRoleById(x.RoleId, c)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New(notis.UndefinedRoleWarnMsg)
	}

	if x.RoleName != "" {
		res.RoleName = strings.TrimSpace(x.RoleName)
	}

	return tr.roleRepo.UpdateRole(*res, c)
}

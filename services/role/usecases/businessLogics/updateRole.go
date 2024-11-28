package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"strings"
)

func (s *service) UpdateRole(x entities.Role, c context.Context) error {
	res, err := s.roleRepo.GetRoleById(x.RoleId, c)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New(notis.UndefinedRoleWarnMsg)
	}

	if x.RoleName != "" {
		res.RoleName = strings.TrimSpace(x.RoleName)
	}

	return s.roleRepo.UpdateRole(*res, c)
}

package businesslogics

import (
	model_types "architecture_template/constants/modelTypes"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"architecture_template/services/role/entities"
	"context"
	"errors"
	"strings"
)

func (s *service) CreateRole(name string, c context.Context) error {
	if name := strings.TrimSpace(name); name == "" {
		return errors.New(notis.FieldEmptyWarnMsg)
	}
	//---------------------------------------
	list, err := s.roleRepo.GetAllRoles(c)
	if err != nil {
		return err
	}
	//---------------------------------------
	return s.roleRepo.CreateRole(entities.Role{
		RoleId:       helper.GenerateId(model_types.ROLE_TYPE, len(*list)),
		RoleName:     name,
		ActiveStatus: true,
	}, c)
}

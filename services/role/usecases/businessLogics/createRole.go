package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/entities"
	"architecture_template/services/role/interfaces"
	"context"
	"errors"
	"fmt"
	"strings"
)

func (tr *service) CreateRole(name string, c context.Context) error {
	if name := strings.TrimSpace(name); name == "" {
		return errors.New(notis.FieldEmptyWarnMsg)
	}
	//---------------------------------------
	id, err := generateRoleId(&tr.roleRepo, c)
	if err != nil {
		return err
	}
	//---------------------------------------
	return tr.roleRepo.CreateRole(entities.Role{
		RoleId:       id,
		RoleName:     name,
		ActiveStatus: true,
	}, c)
}

func generateRoleId(repo *interfaces.IRepository, c context.Context) (string, error) {
	list, err := (*repo).GetAllRoles(c)
	if err != nil {
		return "", err
	}
	//-----------------------------------
	return "R" + fmt.Sprintf("%03d", len(*list)+1), nil
}

package businesslogics

import (
	"architecture_template/services/role/entities"
	"context"
	"strings"
)

func (tr *service) GetRolesByName(name string, c context.Context) (*[]entities.Role, error) {
	if trimStr := strings.TrimSpace(name); trimStr != "" {
		return tr.roleRepo.GetRolesByName(trimStr, c)
	}

	return tr.roleRepo.GetAllRoles(c)
}

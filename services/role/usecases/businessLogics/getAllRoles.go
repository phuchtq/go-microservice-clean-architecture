package businesslogics

import (
	"architecture_template/services/role/entities"
	"context"
)

func (tr *service) GetAllRoles(c context.Context) (*[]entities.Role, error) {
	return tr.roleRepo.GetAllRoles(c)
}

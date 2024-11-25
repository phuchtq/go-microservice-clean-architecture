package repositories

import (
	"context"
)

func (tr *repo) GetAllRoles(isFindId bool, c context.Context) map[string]string {
	return tr.externalRoleService.GetRoleStorage(isFindId, c)
}

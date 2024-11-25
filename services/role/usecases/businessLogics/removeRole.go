package businesslogics

import "context"

func (tr *service) RemoveRole(id string, c context.Context) error {
	return tr.roleRepo.RemoveRole(id, c)
}

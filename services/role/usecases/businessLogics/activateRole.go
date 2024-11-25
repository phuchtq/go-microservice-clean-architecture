package businesslogics

import "context"

func (tr *service) ActivateRole(id string, c context.Context) error {
	return tr.roleRepo.ActivateRole(id, c)
}

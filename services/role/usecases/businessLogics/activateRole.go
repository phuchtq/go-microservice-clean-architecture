package businesslogics

import "context"

func (s *service) ActivateRole(id string, c context.Context) error {
	return s.roleRepo.ActivateRole(id, c)
}

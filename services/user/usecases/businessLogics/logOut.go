package businesslogics

import (
	"architecture_template/services/user/entities"
	"context"
)

func (s *service) LogOut(userId string, c context.Context) error {
	var user *entities.User

	if err := verifyAccount(userId, idValidateType, user, s.repo, c); err != nil {
		return err
	}

	user.AccessToken = nil
	user.RefreshToken = nil

	return s.repo.UpdateUser(*user, c)
}

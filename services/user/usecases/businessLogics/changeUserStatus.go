package businesslogics

import (
	"architecture_template/helper"
	"architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"context"
	"fmt"
)

func (s *service) ChangeUserStatus(rawStatus, userId, actorId string, c context.Context) (error, string) {
	var user *entities.User
	if err := verifyAccount(userId, idValidateType, user, s.repo, c); err != nil {
		return err, ""
	}

	var actor *entities.User
	if err := verifyAccount(actorId, idValidateType, actor, s.repo, c); err != nil {
		return err, ""
	}

	var roles = s.repo.GetAllRoles(false, c)

	if err := verifyEditAuthorization(request.PublicUserInfo{
		UserId:       userId,
		RoleId:       user.RoleId,
		Email:        user.Email,
		Pasword:      user.Pasword,
		ActiveStatus: fmt.Sprint(user.ActiveStatus),
	}, user, actor, roles); err != nil {
		return err, ""
	}

	if isRemain, err := helper.IsStatusRemain(user.ActiveStatus, rawStatus); err != nil {
		return err, ""
	} else if isRemain { // Status not change
		return nil, ""
	}

	if err := s.repo.ChangeUserStatus(!user.ActiveStatus, userId, c); err != nil { // As updated status is different from prev one, negativate the old one
		return err, ""
	}

	if actorId == userId { // As self locking -> Redirect to login page as log out
		return nil, getLoginPageUrl()
	}

	return nil, ""
}

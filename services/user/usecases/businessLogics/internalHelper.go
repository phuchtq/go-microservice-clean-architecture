package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/interfaces"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"log"
	"os"
)

const (
	seperateChar      string = ":"
	activateType      string = "1"
	resetPassType     string = "2"
	updateProfileType string = "3"
	LoginPageUrl      string = "Your-login-page-url"
	resetFlag         string = "Reset"
	activateFlag      string = "Activate"
	idValidateType    string = "id"
	emailValidateType string = "email"
)

func getProcessUrl() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//------------------------------------
	return "http://localhost:" + port + "/VerifyAction?rawToken="
}

func getResetPassUrl() string {
	return "Your reset-pass URL page?token="
}

func getLoginPageUrl() string {
	return "Your-login-page-url"
}

func isEmailExisted(email string, repo *interfaces.IRepository, c context.Context) (bool, error) {
	user, err := (*repo).GetUserByEmail(email, c)

	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil // Email not registered -> Approve
}

func verifyAccount(searchField, validateType string, user *entities.User, repo interfaces.IRepository, c context.Context) error {
	if searchField == "" {
		return errors.New(notis.GenericsErrorWarnMsg)
	}

	var res error

	switch validateType {
	case idValidateType:
		user, res = repo.GetUserById(searchField, c)
	case emailValidateType:
		user, res = repo.GetUserByEmail(searchField, c)
	}

	if user == nil && res == nil {
		switch validateType {
		case idValidateType:
			res = errors.New(user_notis.UndefinedUserWarnMsg)
		case emailValidateType:
			res = errors.New(user_notis.WrongCredentialsWarnMsg)
		}
	}

	return res
}

func verifyUserState(userId, validateType string, user *entities.User, repo interfaces.IRepository, c context.Context) error {
	if err := verifyAccount(userId, idValidateType, user, repo, c); err != nil {
		return err
	}

	if !user.ActiveStatus {
		return errors.New(user_notis.LockWarnMsg)
	}

	return nil
}

func verifyDataFromToken(actionToken string, user *entities.User, logger *log.Logger) error {
	tmpUserId, tmpRole, _, err := utils.ExtractDataFromToken(actionToken, logger)
	if err != nil {
		return err
	}

	if user != nil {
		if tmpUserId != user.UserId || tmpRole != user.RoleId {
			return errors.New(notis.GenericsErrorWarnMsg)
		}
	}

	if user.ActionToken == nil || *user.ActionToken != actionToken {
		return errors.New(notis.GenericsErrorWarnMsg)
	}

	if user.ActionPeriod == nil || helper.IsActionExpired(*user.ActionPeriod, utils.NormalActionDuration) {
		return errors.New(notis.ExpirationWarnMsg)
	}

	return nil
}

func isActionTypeValid(actionType string) bool {
	return actionType == activateType || actionType == resetPassType || actionType == updateProfileType
}

func verifyEditAuthorization(user request.PublicUserInfo, originUser, actor *entities.User, roles map[string]string) error {
	if user.UserId != actor.UserId {
		return verifyEditedAuth(user.RoleId, actor.RoleId, originUser.RoleId, roles) // Edited by other
	}

	return verifyEditAuth(user.RoleId, originUser.RoleId) // Edit themselves
}

func verifyEditedAuth(role, actorRole, originRole string, roles map[string]string) error {
	var res error

	switch actorRole {
	case roles["Admin"]:
		if originRole == roles["Admin"] {
			res = errors.New(user_notis.AdminEditAdmin)
		}
	case roles["Staff"]:
		if role != "" && role != originRole {
			res = errors.New(notis.GenericsRightAccessWarnMsg)
		}
	case roles["Customer"]:
		res = errors.New(notis.GenericsRightAccessWarnMsg)
	}

	return res
}

func verifyEditAuth(role, originRole string) error {
	if role != "" && role != originRole {
		return errors.New(user_notis.EditOwnRoleWarnMsg)
	}

	return nil
}

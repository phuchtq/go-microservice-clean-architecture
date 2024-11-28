package businesslogics

import (
	common_dtos "architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	model_types "architecture_template/constants/modelTypes"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"strings"
	"time"
)

// This method used for registering new account from a guest or a staff-role worker creates new account for someone with specific conditions.
func (s *service) AddUser(u request.SignUpModel, actorId string, c context.Context) (error, string) {
	var actor *entities.User

	// Validate if actor id not empty
	if actorId != "" {
		if err := verifyAccount(actorId, idValidateType, actor, s.repo, c); err != nil {
			return err, ""
		}
	}

	// New email exists?
	if isExist, err := isEmailExisted(u.Email, &s.repo, c); err != nil {
		return err, ""
	} else if isExist {
		return errors.New(notis.EmailRegisteredWarnMsg), ""
	}

	// Check password secure
	if !helper.IsPasswordSecure(u.Password) {
		return errors.New(user_notis.PasswordNotSecureWarnMsg), ""
	}

	// Hash password
	var orgPass string = u.Password
	u.Password = helper.ToHashString(u.Password)

	// Define role for new account
	var roles = s.repo.GetAllRoles(false, c)
	if actorId == "" || u.RoleId == "" {
		u.RoleId = roles["Customer"]
	} else {
		if err := validateRole(u.RoleId, actor.RoleId, roles); err != nil {
			return err, ""
		}
	}

	// Generate id
	list, err := s.repo.GetAllUsers(c)
	if err != nil {
		return err, ""
	}

	var id string = helper.GenerateId(model_types.USER_TYPE, len(*list))
	//-------------------------------------------------------

	// Generate token
	token, _, err := utils.GenerateTokens(u.Email, id, u.RoleId, s.logger)
	if err != nil {
		return err, ""
	}

	// Belongs to last fail access of a new account
	var tmpTime = helper.GetPrimitiveTime()

	// Flag if account need to reset password (in case staff role creates)
	var isHaveToResetPw *bool = nil
	if actorId != "" {
		var flag bool = true
		isHaveToResetPw = &flag
	}

	// Belongs to the moment to request this action
	var tmpCurTime = time.Now().UTC()

	// Save new account to database
	if err := s.repo.AddUser(entities.User{
		UserId:          id,
		RoleId:          u.RoleId,
		Email:           strings.ToLower(strings.TrimSpace(u.Email)),
		Pasword:         u.Password,
		ActiveStatus:    false,
		FailAccess:      0,
		LastFail:        &tmpTime,
		ActionToken:     &token,
		ActionPeriod:    &tmpCurTime,
		IsHaveToResetPw: isHaveToResetPw,
	}, c); err != nil {
		return err, ""
	}

	// Send confirmation mail
	if err := helper.SendMail(common_dtos.SendMailReqDto{
		Body: common_dtos.MailBody{ // Mail body
			Email:    u.Email,
			Password: orgPass,
			Url: utils.GenerateCallBackUrl( // Call back url when guest clicks to the confirmation, it will call back to the api endpoint which generate here to verify and finish the registration process
				[]string{
					getProcessUrl(),
					token,
					id,
					activateType,
				},
				seperateChar,
			),
		},

		TemplatePath: mailconst.AccountRecoveryTemplate, // Template path

		Subject: notis.RegistrationAccountSubject, // Mail subject

		Logger: s.logger, // Logger

	}); err != nil {
		return err, ""
	}

	// Define response message
	var msg string = "Success"
	if actorId == "" {
		msg = user_notis.RegistrationAccountMsg
	}

	return nil, msg
}

func validateRole(roleId, actorRole string, roles map[string]string) error {
	if _, isExist := roles[roleId]; !isExist {
		return errors.New(notis.UndefinedRoleWarnMsg)
	}

	// Actor is a staff
	if actorRole == roles["Staff"] {
		// Staff creates an Admin account
		if roleId == roles["Admin"] {
			return errors.New(notis.GenericsRightAccessWarnMsg)
		}
	}

	return nil
}

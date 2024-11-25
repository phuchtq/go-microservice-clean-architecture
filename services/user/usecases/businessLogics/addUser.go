package businesslogics

import (
	common_dtos "architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/interfaces"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"strings"
	"time"
)

func (s *service) AddUser(u request.SignUpModel, actorId string, c context.Context) (error, string) {
	var actor *entities.User
	if err := verifyAccount(actorId, idValidateType, actor, s.repo, c); err != nil {
		return err, ""
	}

	if isExist, err := isEmailExisted(u.Email, &s.repo, c); err != nil {
		return err, ""
	} else if isExist {
		return errors.New(notis.EmailRegisteredWarnMsg), ""
	}

	if !helper.IsPasswordSecure(u.Password) {
		return errors.New(user_notis.PasswordNotSecureWarnMsg), ""
	}

	var orgPass string = u.Password
	u.Password = helper.ToHashString(u.Password)

	var roles = s.repo.GetAllRoles(false, c)
	if actorId == "" || u.RoleId == "" {
		u.RoleId = roles["Customer"]
	} else {
		if err := validateRole(u.RoleId, actor.RoleId, roles); err != nil {
			return err, ""
		}
	}

	tmpTime := helper.GetPrimitiveTime()
	id, err := generatId(&s.repo, c)
	if err != nil {
		return err, ""
	}

	tmpToken, _, err := utils.GenerateTokens(u.Email, id, u.RoleId, s.logger)
	if err != nil {
		return err, ""
	}

	tmpCurTime := time.Now().UTC()

	var isHaveToResetPw *bool = nil
	if actorId != "" {
		var flag bool = true
		isHaveToResetPw = &flag
	}

	if err := s.repo.AddUser(entities.User{
		UserId:          id,
		RoleId:          u.RoleId,
		Email:           strings.ToLower(strings.TrimSpace(u.Email)),
		Pasword:         u.Password,
		ActiveStatus:    false,
		FailAccess:      0,
		LastFail:        &tmpTime,
		ActionToken:     &tmpToken,
		ActionPeriod:    &tmpCurTime,
		IsHaveToResetPw: isHaveToResetPw,
	}, c); err != nil {
		return err, ""
	}

	var url string = utils.GenerateCallBackUrl(
		[]string{
			getProcessUrl(),
			tmpToken,
			id,
			activateType,
		},
		seperateChar,
	)

	if err := helper.SendMail(common_dtos.SendMailReqDto{
		Body: common_dtos.MailBody{ // Mail body
			Email:    u.Email,
			Password: orgPass,
			Url:      url,
		},
		TemplatePath: mailconst.AccountRecoveryTemplate, // Template path
		Subject:      notis.RegistrationAccountSubject,  // Mail subject
		Logger:       s.logger,                          // Logger
	}); err != nil {
		return err, ""
	}

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

func generatId(repo *interfaces.IRepository, c context.Context) (string, error) {
	list, err := (*repo).GetAllUsers(c)

	if err != nil {
		return "", err
	}

	var id string = helper.GenerateId(strings.ToUpper(entities.GetTable()), len(*list))
	var errMsg error

	if id == "" {
		errMsg = errors.New(notis.GenericsErrorWarnMsg)
	}

	return id, errMsg
}

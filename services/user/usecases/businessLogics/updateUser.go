package businesslogics

import (
	"architecture_template/common_dtos/request"
	common_request "architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	user_requests "architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/interfaces"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"log"
	"strings"
	"time"
)

func (s *service) UpdateUser(user user_requests.PublicUserInfo, actorId string, c context.Context) (string, error) {
	var originUser *entities.User
	if err := verifyAccount(user.UserId, idValidateType, originUser, s.repo, c); err != nil { // Account exists?
		return "", err
	}

	var actor *entities.User
	if err := verifyAccount(actorId, idValidateType, actor, s.repo, c); err != nil { // Account exists?
		return "", err
	}

	var roles = s.repo.GetAllRoles(false, c)

	if err := verifyEditAuthorization(user, originUser, actor, roles); err != nil {
		return "", err
	}

	if actorId == user.UserId { // Update themselves
		if !helper.IsPasswordSecure(user.Pasword) {
			return "", errors.New(user_notis.PasswordNotSecureWarnMsg)
		}
	}

	originUser.Pasword = helper.ToHashString(user.Pasword)

	if err := verifyEditStatus(originUser, user.ActiveStatus); err != nil { // Invalid status
		return "", err
	}

	var msg string = "Success"

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email != "" && user.Email != originUser.Email { // New email's different
		if err := verifyEditEmail(originUser, user.Email, s.repo, c, s.logger); err != nil {
			return "", err
		}

		msg = user_notis.UpdateMailMsg
	}

	if err := s.repo.UpdateUser(*originUser, c); err != nil {
		return "", err
	}

	return msg, nil
}

func verifyEditStatus(originUser *entities.User, updatedRawStatus string) error {
	var activeStatus bool

	if updatedRawStatus != "" {
		var err error
		if activeStatus, err = helper.IsStatusValid(updatedRawStatus); err != nil {
			return err
		}

		if activeStatus != originUser.ActiveStatus {
			var tmpLastFail *time.Time = nil
			if activeStatus {
				*tmpLastFail = helper.GetPrimitiveTime()
			}

			originUser.FailAccess = 0
			originUser.LastFail = tmpLastFail
			originUser.ActiveStatus = activeStatus
		}
	}

	return nil
}

func verifyEditEmail(originUser *entities.User, email string, repo interfaces.IRepository, c context.Context, logger *log.Logger) error {
	var tmpUser *entities.User
	if err := verifyAccount(email, emailValidateType, tmpUser, repo, c); err != nil && err.Error() != user_notis.WrongCredentialsWarnMsg { // Check if new email is exist
		return err // Database connection error
	}

	if tmpUser != nil {
		return errors.New(user_notis.EmailRegisteredWarnMsg) // User found -> email exists -> Deny
	}

	actionToken, _, err := utils.GenerateTokens(email, originUser.UserId, originUser.RoleId, logger)
	if err != nil {
		return err
	}

	*originUser.ActionPeriod = time.Now().UTC()
	*originUser.ActionToken = actionToken

	if err := helper.SendMail(request.SendMailReqDto{
		Body: common_request.MailBody{ // Mail body
			Email: email,
			Url: utils.GenerateCallBackUrl([]string{
				getProcessUrl(),
				actionToken,
				originUser.UserId,
				updateProfileType,
				email,
			}, seperateChar),
		},

		TemplatePath: mailconst.UpdateMailTemplate,

		Subject: notis.UpdateMailSubject,

		Logger: logger,
	}); err != nil {
		return err
	}

	return nil
}

package businesslogics

import (
	"architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	"architecture_template/constants/notis"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"log"
	"strings"
	"time"
)

func (s *service) RecoverAccountByCustomer(email string, c context.Context) (string, error) {
	var account *entities.User
	if err := verifyAccount(strings.TrimSpace(email), emailValidateType, account, s.repo, c); err != nil {
		return "", err
	}

	if account.ActiveStatus {
		return "", errors.New(user_notis.StillActiveAccountMsg)
	}

	if account.RoleId != s.repo.GetAllRoles(false, c)["Customer"] { // Recover staff account -> Just admins have rights to recover these cases -> Deny
		return "", errors.New(user_notis.RecoverStaffAccountWarnMsg)
	}

	if account.FailAccess > maxCustomerLockPeriod {
		if account.LastFail != nil { // Banned -> Contact admin
			return "", errors.New(user_notis.AccountBanWarnMsg)
		}

		if err := setUpRecoverAccount(account, s.logger); err != nil {
			return "", err
		}

		if err := s.repo.UpdateUser(*account, c); err != nil {
			return "", err
		}

		return user_notis.RecoverAccountMsg, nil
	}

	return "", errors.New(user_notis.StillActiveAccountMsg)
}

func setUpRecoverAccount(account *entities.User, logger *log.Logger) error {
	token, _, err := utils.GenerateTokens(account.Email, account.UserId, account.RoleId, logger)

	if err != nil {
		return err
	}

	*account.ActionPeriod = time.Now().UTC()
	*account.ActionToken = token

	return helper.SendMail(request.SendMailReqDto{
		Body: request.MailBody{ // Mail body
			Email: account.Email,
			Url: utils.GenerateCallBackUrl([]string{
				getProcessUrl(),
				token,
			}, seperateChar),
		},

		TemplatePath: mailconst.AccountRecoveryTemplate, // Template path for generating mail

		Subject: notis.RecoverAccountSubject, // Mail subject

		Logger: logger, // Logger
	})
}

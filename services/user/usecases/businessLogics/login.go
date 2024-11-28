package businesslogics

import (
	common_request "architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	"architecture_template/constants/notis"
	post_types "architecture_template/constants/postTypes"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/interfaces"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"log"
	"strings"
	"time"
)

const (
	minStaffLockPeriod    int = 4
	minCustomerLockPeriod int = 3
	maxCustomerLockPeriod int = 5
)

var customerFailRange = []int{
	minCustomerLockPeriod,
	4,
	maxCustomerLockPeriod,
}

func (s *service) Login(email string, password string, c context.Context) (string, string, error) {
	var acc *entities.User
	if err := verifyAccount(strings.TrimSpace(email), emailValidateType, acc, s.repo, c); err != nil {
		return "", "", err
	}

	var isCorrectCredentials bool = utils.IsLoginPasswordMatched(acc.Pasword, password)
	var isActivated bool = isAccountActivated(acc)
	var roles = s.repo.GetAllRoles(false, c)

	if !isCorrectCredentials {
		return processFailCase(acc, isActivated, s.repo, roles, c)
	}

	return processSuccessCase(acc, isActivated, s.repo, roles, s.logger, c)
}

func isAccountActivated(acc *entities.User) bool {
	return !(!acc.ActiveStatus && acc.FailAccess == 0 && acc.ActionPeriod == nil && acc.LastFail == nil || *acc.LastFail == helper.GetPrimitiveTime())
}

func processSuccessCase(acc *entities.User, isActivated bool, repo interfaces.IRepository, roles map[string]string, logger *log.Logger, c context.Context) (string, string, error) {
	if acc.LastFail == nil { // Self locking account
		return "", "", errors.New(user_notis.InactiveAccountMsg)
	}

	var res1 string // Can be flag for type of login or access token
	var res2 string // Can be message to user or a redirect url or refresh token

	if !isActivated { // Still not activated
		if err := prepareActivateAccount(acc, logger); err != nil {
			return "", "", err
		}

		res1 = post_types.ActivateCase
		res2 = user_notis.ActivateAccountMsg
	}

	if !isLockExpired(acc, roles) {
		return "", "", errors.New(user_notis.LockWarnMsg)
	}

	if acc.FailAccess > maxCustomerLockPeriod && *acc.LastFail != helper.GetPrimitiveTime() { //Banned
		return "", "", errors.New(user_notis.AccountBanWarnMsg)
	}

	if res1 == "" {
		accessToken, refreshToken, err := utils.GenerateTokens(acc.Email, acc.UserId, acc.RoleId, logger)

		if err != nil {
			return "", "", err
		}

		if acc.IsHaveToResetPw != nil && *acc.IsHaveToResetPw { // Have to reset password
			res1 = post_types.ResetCase
			res2 = utils.GenerateCallBackUrl([]string{ // Generate reset password url
				getResetPassUrl(),
				accessToken,
			}, seperateChar)

			*acc.ActionToken = accessToken
			*acc.ActionPeriod = time.Now().UTC()
		} else {
			*acc.AccessToken = accessToken
			*acc.RefreshToken = refreshToken

			res1 = accessToken
			res2 = refreshToken
		}
	}

	if err := refreshAccount(acc, repo, c); err != nil {
		return "", "", err
	}

	return res1, res2, nil
}

func prepareActivateAccount(acc *entities.User, logger *log.Logger) error {
	actionToken, _, err := utils.GenerateTokens(acc.Email, acc.UserId, acc.RoleId, logger)

	if err != nil {
		return err
	}

	acc.ActionToken = &actionToken
	*acc.ActionPeriod = time.Now().UTC()

	return helper.SendMail(common_request.SendMailReqDto{
		Body: common_request.MailBody{ // Mail body
			Email: acc.Email,
			Url: utils.GenerateCallBackUrl([]string{
				getProcessUrl(),
				actionToken,
				acc.UserId,
				activateType,
			}, seperateChar),
		},

		TemplatePath: mailconst.AccountRegistrationTemplate, // Template path

		Subject: notis.RegistrationAccountSubject, // Mail subject

		Logger: logger, // Logger
	})
}

func refreshAccount(acc *entities.User, repo interfaces.IRepository, c context.Context) error {
	*acc.LastFail = helper.GetPrimitiveTime()
	acc.ActiveStatus = true

	return repo.UpdateUser(*acc, c)
}

func processFailCase(acc *entities.User, isActivated bool, repo interfaces.IRepository, roles map[string]string, c context.Context) (string, string, error) {
	if !isActivated {
		return "", "", errors.New(user_notis.WrongCredentialsWarnMsg)
	}

	if isLockExpired(acc, roles) {
		if err := updateFailCase(acc, repo, c); err != nil {
			return "", "", err
		}

		return "", "", errors.New(user_notis.WrongCredentialsWarnMsg)
	}

	return "", "", errors.New(user_notis.LockWarnMsg)
}

func isLockExpired(acc *entities.User, roles map[string]string) bool {
	lockDuration, lockPeriod := getLockDurationAndPeriod(acc.RoleId, acc.FailAccess, roles)

	if acc.FailAccess >= lockPeriod {
		return helper.IsActionExpired(*acc.LastFail, lockDuration)
	}

	return true
}

func getCustomerLockPeriods() map[int]time.Duration {
	res := make(map[int]time.Duration)
	//-----------------------------------
	var primitiveDuration time.Duration = 15 * time.Minute
	//-----------------------------------
	for _, v := range customerFailRange {
		res[v] = primitiveDuration
		primitiveDuration *= 2
	}
	//-----------------------------------
	return res
}

func getLockDurationAndPeriod(role string, fail int, roles map[string]string) (time.Duration, int) {
	var duration time.Duration
	var failPeriod int

	switch role {
	case roles["Admin"]:
		duration = utils.AdminLockDuration
		failPeriod = minStaffLockPeriod
	case roles["Staff"]:
		duration = utils.StaffLockDuration
		failPeriod = minStaffLockPeriod
	case roles["Customer"]:
		duration = getCustomerLockPeriods()[fail]
		failPeriod = minCustomerLockPeriod
	default:
		duration = 0
		failPeriod = 0
	}

	return duration, failPeriod
}

func updateFailCase(acc *entities.User, repo interfaces.IRepository, c context.Context) error {
	var tmpCurTime time.Time = time.Now().UTC()
	acc.LastFail = &tmpCurTime

	acc.FailAccess += 1

	acc.ActiveStatus = false

	return repo.UpdateUser(*acc, c)
}

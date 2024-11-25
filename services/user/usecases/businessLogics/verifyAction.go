package businesslogics

import (
	"architecture_template/constants/notis"
	"architecture_template/helper"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/utils"
	"context"
	"errors"
	"log"
	"strings"
	"time"
)

func (s *service) VerifyAction(rawToken string, c context.Context) (error, string) {
	var cmp []string = strings.Split(rawToken, seperateChar)
	if len(cmp) < 3 { // Min length of combination of information in a call back url
		return errors.New(notis.GenericsErrorWarnMsg), ""
	}

	var actionToken string = cmp[0]
	var userId string = cmp[1]
	var actionType string = cmp[2]
	var user *entities.User

	if err := verifyUserState(userId, idValidateType, user, s.repo, c); err != nil { // Check if user data extract from call back url is valid
		return err, ""
	}

	if err := verifyDataFromToken(actionToken, user, s.logger); err != nil { // Check if user data from token is valid
		return err, ""
	}

	if !isActionTypeValid(actionType) { // Check if action type found in url is valid
		return errors.New(notis.GenericsErrorWarnMsg), ""
	}

	go categorizeActionType(actionType, cmp, user) // Categorize action type, update user data based on that type

	var res string
	if *user.IsHaveToResetPw || actionType == resetPassType {
		redirectUrl, err := setUpBeforeResetPw(user, s.logger) // Generate url to redirect user to reset page

		if err != nil {
			return err, ""
		}

		res = redirectUrl
	}

	if err := s.repo.UpdateUser(*user, c); err != nil {
		return err, ""
	}

	return nil, res
}

func categorizeActionType(actionType string, rawTokenCmps []string, user *entities.User) {
	switch actionType {
	case actionType:
		user.ActiveStatus = true
		*user.LastFail = helper.GetPrimitiveTime()
	case updateProfileType:
		user.Email = rawTokenCmps[len(rawTokenCmps)-1]
	default:
	}

	user.ActionToken = nil
	user.ActionPeriod = nil
}

func setUpBeforeResetPw(user *entities.User, logger *log.Logger) (string, error) { // Generate reset page url for user if needed to reset pw
	token, _, err := utils.GenerateTokens(user.Email, user.UserId, user.RoleId, logger)
	if err != nil {
		return "", err
	}

	*user.ActionToken = token
	*user.ActionPeriod = time.Now().UTC()

	return utils.GenerateCallBackUrl([]string{
		getResetPassUrl(),
		token,
	}, seperateChar), nil
}

package businesslogics

import (
	"architecture_template/common_dtos/response"
	"architecture_template/helper"
	user_notis "architecture_template/services/user/constants/notis"
	"architecture_template/services/user/entities"
	"architecture_template/services/user/utils"
	"context"
)

func (s *service) ResetPassword(newPass, re_newPass, token string, c context.Context) response.ResetPasswordResponse {
	var resetPageUrl string = utils.GenerateCallBackUrl([]string{
		getResetPassUrl(),
		token,
	}, seperateChar)

	if newPass != re_newPass {
		return response.ResetPasswordResponse{
			RedirectUrl: resetPageUrl,
			ErrorMsg:    user_notis.PasswordsNotMatchWarnMsg,
		}
	}

	if !helper.IsPasswordSecure(newPass) {
		return response.ResetPasswordResponse{
			RedirectUrl: resetPageUrl,
			ErrorMsg:    user_notis.PasswordNotSecureWarnMsg,
		}
	}

	var user *entities.User
	userId, _, _, _ := utils.ExtractDataFromToken(token, s.logger)

	if err := verifyUserState(userId, idValidateType, user, s.repo, c); err != nil {
		return response.ResetPasswordResponse{
			RedirectUrl: resetPageUrl,
			ErrorMsg:    err.Error(),
		}
	}

	if err := verifyDataFromToken(token, user, s.logger); err != nil { // Check if user data from token is valid
		return response.ResetPasswordResponse{
			RedirectUrl: resetPageUrl,
			ErrorMsg:    err.Error(),
		}
	}

	user.Pasword = helper.ToHashString(newPass)

	if err := s.repo.UpdateUser(*user, c); err != nil {
		return response.ResetPasswordResponse{
			RedirectUrl: resetPageUrl,
			ErrorMsg:    err.Error(),
		}
	}

	return response.ResetPasswordResponse{
		RedirectUrl: getLoginPageUrl(),
	}
}

package interfaces

import (
	"architecture_template/common_dtos/response"
	"architecture_template/services/user/dtos/request"
	"architecture_template/services/user/entities"
	"context"
)

type IService interface {
	GetAllUsers(c context.Context) (*[]entities.User, error)
	GetUsersByRole(role string, c context.Context) (*[]entities.User, error)
	GetUsersByStatus(rawStatus string, c context.Context) (*[]entities.User, error)
	GetUserById(id string, c context.Context) (*entities.User, error)
	AddUser(u request.SignUpModel, actorId string, c context.Context) (error, string)
	UpdateUser(user request.PublicUserInfo, actorId string, c context.Context) (string, error)
	ChangeUserStatus(rawStatus, userId, actorId string, c context.Context) (error, string)
	Login(email string, password string, c context.Context) (string, string, error)
	LogOut(userId string, c context.Context) error
	VerifyAction(rawToken string, c context.Context) (error, string)
	ResetPassword(newPass, re_newPass, token string, c context.Context) response.ResetPasswordResponse
	RecoverAccountByCustomer(email string, c context.Context) (string, error)
}

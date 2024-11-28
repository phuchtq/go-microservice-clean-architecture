package interfaces

import (
	"architecture_template/services/user/entities"
	"context"
)

type IRepository interface {
	GetAllUsers(c context.Context) (*[]entities.User, error)
	GetUsersByRole(id string, c context.Context) (*[]entities.User, error)
	GetUsersByStatus(status bool, c context.Context) (*[]entities.User, error)
	GetUserById(id string, c context.Context) (*entities.User, error)
	GetUserByEmail(email string, c context.Context) (*entities.User, error)
	AddUser(u entities.User, c context.Context) error
	UpdateUser(u entities.User, c context.Context) error
	ChangeUserStatus(status bool, id string, c context.Context) error

	// External service
	GetAllRoles(isFindId bool, c context.Context) map[string]string
}

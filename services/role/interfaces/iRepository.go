package interfaces

import (
	"architecture_template/services/role/entities"
	"context"
)

type IRepository interface {
	GetAllRoles(c context.Context) (*[]entities.Role, error)
	GetRolesByName(name string, c context.Context) (*[]entities.Role, error)
	GetRolesByStatus(status bool, c context.Context) (*[]entities.Role, error)
	GetRoleById(id string, c context.Context) (*entities.Role, error)
	CreateRole(r entities.Role, c context.Context) error
	RemoveRole(id string, c context.Context) error
	UpdateRole(r entities.Role, c context.Context) error
	ActivateRole(id string, c context.Context) error
}

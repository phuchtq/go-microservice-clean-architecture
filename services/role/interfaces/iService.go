package interfaces

import (
	"architecture_template/services/role/entities"
	"context"
)

type IRoleService interface {
	GetAllRoles(c context.Context) (*[]entities.Role, error)
	GetRolesByName(name string, c context.Context) (*[]entities.Role, error)
	GetRolesByStatus(rawStatus string, c context.Context) (*[]entities.Role, error)
	GetRoleById(id string, c context.Context) (*entities.Role, error)
	CreateRole(name string, c context.Context) error
	UpdateRole(x entities.Role, c context.Context) error
	RemoveRole(id string, c context.Context) error
	ActivateRole(id string, c context.Context) error
}

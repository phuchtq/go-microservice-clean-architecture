package tests

import "architecture_template/services/role/entities"

const (
	existed     string = "Your role id 1"
	newRole     string = "Your new role name"
	searchKw    string = "Your search key word"
	rawActive   string = "true"
	rawInactive string = "false"
)

var sampleRoles = &[]entities.Role{
	{
		RoleId:       "Your role id 1",
		RoleName:     "Your role name 1",
		ActiveStatus: true,
	},

	{
		RoleId:       "Your role id 2",
		RoleName:     "Your role name 2",
		ActiveStatus: false,
	},

	{
		RoleId:       "Your role id 3",
		RoleName:     "Your role name 3",
		ActiveStatus: true,
	},
}

var sampleActiveRoles = &[]entities.Role{
	{
		RoleId:       "Your role id 1",
		RoleName:     "Your role name 1",
		ActiveStatus: true,
	},

	{
		RoleId:       "Your role id 3",
		RoleName:     "Your role name 3",
		ActiveStatus: true,
	},
}

var sampleInactiveRoles = &[]entities.Role{
	{
		RoleId:       "Your role id 2",
		RoleName:     "Your role name 2",
		ActiveStatus: false,
	},
}

var sampleNameBasedRoles = &[]entities.Role{
	{
		RoleId:       "Your role id 1",
		RoleName:     "Your role name 1",
		ActiveStatus: true,
	},

	{
		RoleId:       "Your role id 2",
		RoleName:     "Your role name 2",
		ActiveStatus: false,
	},
}

var sampleRole = &entities.Role{
	RoleId:       "Your role id 1",
	RoleName:     "Your role name 1",
	ActiveStatus: true,
}

// Struct data types
const (
	role             string = "entities.Role"
	pointerRole      string = "*entities.Role"
	pointerSliceRole string = "*[]entities.Role"
)

// Repository method list
const (
	repoGetAll      string = "GetAllRoles"
	repoGetByName   string = "GetRolesByName"
	repoGetByStatus string = "GetRolesByStatus"
	repoGetById     string = "GetRoleById"
	repoCreate      string = "CreateRole"
	repoRemove      string = "RemoveRole"
	repoUpdate      string = "UpdateRole"
	repoActivate    string = "ActivateRole"
)

// Service method list
const (
	serviceGetAll      string = "GetAllRoles"
	serviceGetByName   string = "GetRolesByName"
	serviceGetByStatus string = "GetRolesByStatus"
	serviceGetById     string = "GetRoleById"
	serviceCreate      string = "CreateRole"
	serviceRemove      string = "RemoveRole"
	serviceUpdate      string = "UpdateRole"
	serviceActivate    string = "ActivateRole"
)

package entities

type Role struct {
	RoleId       string `json:"role_id"`
	RoleName     string `json:"role_name" validate:"required, min=1"`
	ActiveStatus bool   `json:"active_status"`
}

func GetTable() string {
	return "Roles"
}

func GetDatabaseServer() string {
	return "postgre"
}

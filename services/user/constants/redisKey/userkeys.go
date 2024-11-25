package rediskey

const (
	GetAllKey string = "users"

	GetByRoleKey string = "users:role:%s"

	GetByStatusKey string = "users:status:%t"

	GetByIdKey string = "user:id:%s"

	GetByEmailKey string = "user:email:%s"
)

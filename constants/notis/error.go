package notis

const (
	InternalErr string = "There is something wrong in the system during the process. Please try again later."

	RoleRepoMsg string = "Error in role service - repository at "

	RoleRpcMsg string = "Error while fetching roles from RoleService GRPC - "

	MailHelperMsg string = "Error while generating mail in helper at "

	JsonMsg string = "Error in json helper - "

	RedisMsg string = "Error in storing data to redis cache - data: "
)

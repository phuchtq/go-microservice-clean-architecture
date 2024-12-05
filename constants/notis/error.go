package notis

const (
	InternalErr string = "There is something wrong in the system during the process. Please try again later."

	EnvLoadErr string = "Error while loading .env file in %s service - "

	EnvSetErrMsg string = "Error while setting environment variable %s with value %s - "

	CmdExecuteErrMsg string = "Error while executing command in %s service - "

	DbConnectionMsg string = "Error while connecting to database in %s service - "

	RoleRepoMsg string = "Error in role service - repository at "

	RoleRpcMsg string = "Error while fetching roles from RoleService GRPC - "

	MailHelperMsg string = "Error while generating mail in helper at "

	JsonMsg string = "Error in json helper - "

	RedisMsg string = "Error in storing data to redis cache - data: "

	GinMsg string = "Error while starting gin server in %s service - "

	NetListeningMsg string = "Error while listening on port %s - "
)

// gRPC
const (
	GrpcConnectMsg string = "Error while connecting %s service grpc - "

	GrpcGenerateMsg string = "Error while generating %s service grpc - "

	GrpcServeMsg string = "Error while serving %s service grpc - "
)

// Redis
const (
	RedisExtractDataMsg string = "Error while extracting data in Redis as key '%s' - "

	RedisStoreDataMsg string = "Error while storing data with Redis as key '%s' - "

	RedisRefreshKeyMsg string = "Error while refreshing key with Redis as key '%s' - "
)

// RabbitMQ
const (
	RabbitmqDeclareMsg string = "Error while declaring queue '%s' - "

	RabbitmqConsumeMsg string = "Error while consuming queue '%s' - "

	RabbitmqPublishMsg string = "Error while publishing queue '%s' - "

	RabbitmqConnectMsg string = "Error while connecting RabbitMQ client with connection string '%s' - "
)

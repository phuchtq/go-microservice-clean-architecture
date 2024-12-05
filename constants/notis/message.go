package notis

// Database
const (
	DbServerNotSetMsg string = "Database server in %s service is empty."

	DbCnnStrNotSetMsg string = "Database connection string in %s service has not been set yet."

	DbMigrationInformMsg string = "There is something wrong during database migration."
)

// API
const (
	ApiPortEnvNotSetMsg string = "Api port in %s service has not been set yet."
)

// gRPC
const (
	GrpcPortEnvNotSetMsg string = "%s grpc port has not been set yet."
)

// Redis
const (
	RedisPortEnvNotSetMsg string = "Redis port in %s service has not been set yet."
)

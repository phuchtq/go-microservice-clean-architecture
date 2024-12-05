package routes

import (
	"architecture_template/constants/notis"
	"architecture_template/middlewares/authorization"
	envvar "architecture_template/services/role/constants/envVar"
	"architecture_template/services/user/adapters/api"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	backUpApiPort string = "Your back up api port"
)

func InitializeAPIRoutes() {
	var server = gin.Default()
	var logger = &log.Logger{}

	var port string = os.Getenv(envvar.ApiPort)
	if port == "" {
		logger.Println(fmt.Sprintf(notis.ApiPortEnvNotSetMsg, "User"))

		if err := os.Setenv(envvar.ApiPort, backUpApiPort); err != nil {
			logger.Println(fmt.Sprintf(notis.EnvSetErrMsg, envvar.ApiPort, backUpApiPort) + err.Error())
		}

		port = backUpApiPort
	}

	var contextPath string = "users"

	var adminAuthGroup = server.Group(contextPath, authorization.Authorize, authorization.AdminAuhthorization)
	adminAuthGroup.GET("", api.GetAllUsers)
	adminAuthGroup.GET("/:role", api.GetUsersByRole)
	adminAuthGroup.GET("/:status", api.GetUsersByStatus)

	var authGroup = server.Group(contextPath, authorization.Authorize)
	authGroup.GET("/:id", api.GetUserById)
	authGroup.PUT("", api.UpdateUser)
	authGroup.PUT("/id/:id/status/:status", api.ChangeUserStatus)
	authGroup.PUT("/logout", api.LogOut)

	var norGroup = server.Group(contextPath)
	norGroup.PUT("/login", api.Login)
	norGroup.POST("", api.AddUser)
	norGroup.PUT("/:email", api.RecoverAccountByCustomer)
	norGroup.PUT("/password/:password/confirm-password/:confirmPassword", api.ResetPassword)
	norGroup.PUT("", api.VerifyAction)

	if err := server.Run(":" + port); err != nil {
		logger.Fatalln(fmt.Sprintf(notis.GinMsg, "User") + err.Error())
	}

	logger.Println("User service starts on port: ", port)
}

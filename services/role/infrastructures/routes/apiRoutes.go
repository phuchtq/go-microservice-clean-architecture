package routes

import (
	"architecture_template/constants/notis"
	"architecture_template/middlewares/authorization"
	"architecture_template/services/role/adapters/api"
	envvar "architecture_template/services/role/constants/envVar"
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
	//-----------------------------------------
	var port string = os.Getenv(envvar.ApiPort)
	if port == "" {
		logger.Println(fmt.Sprintf(notis.ApiPortEnvNotSetMsg, "Role"))

		if err := os.Setenv(envvar.ApiPort, backUpApiPort); err != nil {
			logger.Println(fmt.Sprintf(notis.EnvSetErrMsg, envvar.ApiPort, backUpApiPort) + err.Error())
		}

		port = backUpApiPort
	}
	//-----------------------------------------
	var contextPath string = "roles"
	//-----------------------------------------
	var adminAuthGroup = server.Group(contextPath, authorization.Authorize, authorization.AdminAuhthorization)
	//-----------------------------------------
	adminAuthGroup.GET("", api.GetAllRoles)
	adminAuthGroup.GET("/name/:name", api.GetRolesByName)
	adminAuthGroup.GET("/status/:status", api.GetRolesByStatus)
	adminAuthGroup.POST("/:name", api.CreateRole)
	adminAuthGroup.PUT("", api.UpdateRole)
	adminAuthGroup.PUT("/:id", api.ActivateRole)
	adminAuthGroup.DELETE("/:id", api.RemoveRole)
	//-----------------------------------------
	var authGroup = server.Group(contextPath, authorization.Authorize)
	//-----------------------------------------
	authGroup.GET("/:id", api.GetRoleById)
	//-----------------------------------------
	logger.Println("Role service starts on port: ", port)
	//-----------------------------------------
	if err := server.Run(":" + port); err != nil {
		logger.Fatal(fmt.Sprintf(notis.GinMsg, "Role") + err.Error())
	}
}

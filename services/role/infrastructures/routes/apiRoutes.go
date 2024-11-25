package routes

import (
	"architecture_template/services/role/adapters/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeAPIRoutes() {
	server := gin.Default()
	//-----------------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//-----------------------------------------
	if err := server.Run(":" + port); err != nil {
		log.Fatal("InitializeAPIRoutes meet an unexpected error: ", err)
	}
	//-----------------------------------------
	group := server.Group("roles", nil)
	//-----------------------------------------
	group.GET("", api.GetAllRoles)
	group.GET("/name/:name", api.GetRolesByName)
	group.GET("/status/:status", api.GetRolesByStatus)
	group.GET("/:id", api.GetRoleById)
	//-----------------------------------------
	group.POST("/:name", api.CreateRole)
	//-----------------------------------------
	group.PUT("", api.UpdateRole)
	group.PUT("/:id", api.ActivateRole)
	//-----------------------------------------
	group.DELETE("/:id", api.RemoveRole)
}

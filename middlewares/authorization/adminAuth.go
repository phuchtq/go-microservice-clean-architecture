package authorization

import (
	"architecture_template/constants"
	api_response "architecture_template/helper/api_response"

	"github.com/gin-gonic/gin"
)

func AdminAuhthorization(c *gin.Context) {
	if c.GetString("role") != constants.ADMIN {
		api_response.ProcessResponse(getUnauthBodyResponse(c))
		return
	}

	c.Next()
}

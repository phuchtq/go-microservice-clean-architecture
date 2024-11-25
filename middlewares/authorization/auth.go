package authorization

import (
	"architecture_template/helper"
	api_response "architecture_template/helper/api_response"
	"log"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	// Get token from the header
	var token string = c.Request.Header.Get("Authorization")

	var responseBody = getUnauthBodyResponse(c)

	if token == "" {
		api_response.ProcessResponse(responseBody)
		return
	}

	userId, role, expPeriod, err := helper.ExtractDataFromToken(token, &log.Logger{})
	if err != nil {
		api_response.ProcessResponse(responseBody)
		return
	}

	if helper.IsActionExpired(expPeriod, 0) { // Access token expired
		// Track refresh token
		// If expired -> return message
	}

	c.Set("userId", userId)
	c.Set("role", role)
	c.Next()
}

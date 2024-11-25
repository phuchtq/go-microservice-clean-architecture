package authorization

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	"errors"

	"github.com/gin-gonic/gin"
)

func getUnauthBodyResponse(c *gin.Context) response.ApiResponseModel {
	return response.ApiResponseModel{
		ErrMsg:  errors.New(notis.GenericsRightAccessWarnMsg),
		Context: c,
	}
}

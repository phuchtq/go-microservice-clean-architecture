package api

import (
	api_response "architecture_template/helper/api_response"
	business_logics "architecture_template/services/user/usecases/businessLogics"
	"architecture_template/services/user/utils"

	"github.com/gin-gonic/gin"
)

func ChangeUserStatus(c *gin.Context) {
	service, err := business_logics.GenerateService()

	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}

	err, res := service.ChangeUserStatus(c.Param("status"), c.Param("id"), c.GetString("userId"), c)

	utils.ProcessRedirectAndInformResponse(res, err, c)
}

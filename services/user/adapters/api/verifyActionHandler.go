package api

import (
	api_response "architecture_template/helper/api_response"
	business_logics "architecture_template/services/user/usecases/businessLogics"
	"architecture_template/services/user/utils"

	"github.com/gin-gonic/gin"
)

func VerifyAction(c *gin.Context) {
	service, err := business_logics.GenerateService()

	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}

	err, res := service.VerifyAction(c.Query("rawToken"), c)

	utils.ProcessRedirectAndInformResponse(res, err, c)
}

package api

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	business_logics "architecture_template/services/user/usecases/businessLogics"

	"github.com/gin-gonic/gin"
)

func ResetPassword(c *gin.Context) {
	service, err := business_logics.GenerateService()

	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
	}

	api_response.ProcessResponse(response.ApiResponseModel{
		Data2:    service.ResetPassword(c.Param("password"), c.Param("confirmPassword"), c.Query("token"), c).RedirectUrl,
		PostType: post_types.RedirectPost,
		Context:  c,
	})
}

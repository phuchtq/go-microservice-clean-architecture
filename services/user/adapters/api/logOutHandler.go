package api

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	"architecture_template/services/user/constants/url"
	business_logics "architecture_template/services/user/usecases/businessLogics"

	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	service, err := business_logics.GenerateService()

	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}

	api_response.ProcessResponse(response.ApiResponseModel{
		Data2:    url.LoginPageUrl,
		ErrMsg:   service.LogOut(c.GetString("userId"), c),
		PostType: post_types.RedirectPost,
		Context:  c,
	})
}

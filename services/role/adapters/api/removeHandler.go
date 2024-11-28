package api

import (
	"architecture_template/common_dtos/response"
	posttypes "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	business_logics "architecture_template/services/role/usecases/businessLogics"

	"github.com/gin-gonic/gin"
)

func RemoveRole(c *gin.Context) {
	service, err := business_logics.GenerateService()
	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}
	//-----------------------------------------
	api_response.ProcessResponse(response.ApiResponseModel{
		ErrMsg:   service.RemoveRole(c.Param("id"), c),
		PostType: posttypes.ActionPost,
		Context:  c,
	})
}

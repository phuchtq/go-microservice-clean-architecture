package api

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	"architecture_template/services/user/dtos/request"
	business_logics "architecture_template/services/user/usecases/businessLogics"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var model request.SignUpModel
	if err := c.ShouldBindJSON(&model); err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, nil))
		return
	}

	service, err := business_logics.GenerateService()
	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}

	err, res := service.AddUser(model, c.GetString("userId"), c)

	api_response.ProcessResponse(response.ApiResponseModel{
		Data1:    res,
		ErrMsg:   err,
		PostType: post_types.InformPost,
		Context:  c,
	})
}

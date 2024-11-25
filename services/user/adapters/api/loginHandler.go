package api

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	"architecture_template/services/user/dtos/request"
	business_logics "architecture_template/services/user/usecases/businessLogics"
	"architecture_template/services/user/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var model request.LoginModel
	if c.ShouldBindJSON(&model) != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, nil))
		return
	}

	service, err := business_logics.GenerateService()

	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}

	res1, res2, err := service.Login(model.Email, model.Password, c)

	if err != nil {
		api_response.ProcessResponse(response.ApiResponseModel{
			Data1:    res1,
			Data2:    res2,
			ErrMsg:   err,
			PostType: post_types.NonPost,
			Context:  c,
		})

		return
	}

	utils.ProcessLoginResponse(res1, res2, c)
}

package api

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	business_logics "architecture_template/services/role/usecases/businessLogics"

	"github.com/gin-gonic/gin"
)

func ActivateRole(c *gin.Context) {
	// if !isAdminAccess(c.GetString("role")) {
	// 	c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
	// 	return
	// }
	//-----------------------------------------
	service, err := business_logics.GenerateService()
	if err != nil {
		api_response.ProcessResponse(api_response.GenerateInvalidRequestAndSystemProblemModel(c, err))
		return
	}
	//-----------------------------------------
	api_response.ProcessResponse(response.ApiResponseModel{
		ErrMsg:   service.ActivateRole(c.Param("id"), c),
		PostType: post_types.ActionPost,
		Context:  c,
	})
}

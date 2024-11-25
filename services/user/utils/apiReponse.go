package utils

import (
	"architecture_template/common_dtos/response"
	post_types "architecture_template/constants/postTypes"
	api_response "architecture_template/helper/api_response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessLoginResponse(res1, res2 interface{}, c *gin.Context) {
	var stringRes1 string = fmt.Sprint(res1)
	var stringRes2 string = fmt.Sprint(res2)
	//--------------------------------------
	switch res1 {
	case post_types.RedirectPost:
		api_response.ProcessResponse(response.ApiResponseModel{
			Data1:    stringRes1,
			Data2:    stringRes2,
			PostType: post_types.RedirectPost,
			Context:  c,
		})
	case post_types.ActivateCase:
		c.IndentedJSON(http.StatusContinue, gin.H{"message": stringRes2})
	default:
		c.IndentedJSON(http.StatusOK, gin.H{
			"access_token":  stringRes1,
			"refresh_token": stringRes2,
		})
	}
}

func ProcessRedirectAndInformResponse(res interface{}, err error, c *gin.Context) {
	var data interface{} = fmt.Sprint(res)
	var postType string = post_types.RedirectPost

	if fmt.Sprint(res) == "" {
		data = "Success"
		postType = post_types.InformPost
	}

	api_response.ProcessResponse(response.ApiResponseModel{
		Data2:    data,
		ErrMsg:   err,
		PostType: postType,
		Context:  c,
	})
}

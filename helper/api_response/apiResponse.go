package apiresponse

import (
	"architecture_template/common_dtos/response"
	"architecture_template/constants/notis"
	post_types "architecture_template/constants/postTypes"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateInvalidRequestAndSystemProblemModel(c *gin.Context, err error) response.ApiResponseModel {
	var errMsg error = err
	if errMsg == nil {
		errMsg = errors.New(notis.GenericsErrorWarnMsg)
	}

	return response.ApiResponseModel{
		ErrMsg:   errMsg,
		Context:  c,
		PostType: post_types.NonPost,
	}
}

func ProcessResponse(data response.ApiResponseModel) {
	if data.ErrMsg != nil {
		processFailResponse(data.ErrMsg, data.Context)
		return
	}

	if data.PostType != post_types.NonPost {
		processSuccessPostReponse(data.Data2, data.PostType, data.Context)
		return
	}

	processSuccessResponse(data.Data1, data.Context)
}

func processFailResponse(err error, c *gin.Context) {
	var errCode int

	switch err.Error() {
	case notis.InternalErr:
		errCode = http.StatusInternalServerError
	case notis.GenericsRightAccessWarnMsg:
		errCode = http.StatusForbidden
	default:
		errCode = http.StatusBadRequest
	}

	if isErrorTypeOfUndefined(err) {
		errCode = http.StatusNotFound
	}

	c.IndentedJSON(errCode, gin.H{"message": err.Error()})
}

func processSuccessPostReponse(res interface{}, postType string, c *gin.Context) {
	switch postType {
	case post_types.RedirectPost:
		processRedirectResponse(fmt.Sprint(res), c)
	case post_types.InformPost:
		processInformResponse(res, c)
	default:
		c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func processRedirectResponse(redirectUrl string, c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

func processInformResponse(message interface{}, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprint(message)})
}

func processSuccessResponse(data interface{}, c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

func isErrorTypeOfUndefined(err error) bool {
	return strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "undefined")
}

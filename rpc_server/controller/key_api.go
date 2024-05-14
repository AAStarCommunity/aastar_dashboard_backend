package controller

import (
	"aastar_dashboard_back/model"
	"github.com/gin-gonic/gin"
)

// GetApiKeyList
// @Tags GetApiKeyList
// @Description GetApiKeyList
// @Accept json
// @Product json
// @Param user_id query string true "User ID"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func GetApiKeyList(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// GetApiKey
// @Tags GetApiKey
// @Description GetApiKey
// @Accept json
// @Product json
// @Param user_id query string true "User ID"
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func GetApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// UpdateApiKey
// @Tags UpdateApiKey
// @Description UpdateApiKey
// @Accept json
// @Product json
// @Param uploadApiKeyRequest  body  model.UploadApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func UpdateApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// AddApiKey
// @Tags AddApiKey
// @Description AddApiKey
// @Accept json
// @Product json
// @Param uploadApiKeyRequest  body  model.UploadApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func AddApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// DeleteApiKey
// @Tags DeleteApiKey
// @Description DeleteApiKey
// @Accept json
// @Product json
// @Param user_id query string true "User ID"
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func DeleteApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

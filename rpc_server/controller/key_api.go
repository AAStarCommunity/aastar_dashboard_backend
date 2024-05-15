package controller

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetApiKeyList
// @Tags GetApiKeyList
// @Description GetApiKeyList
// @Accept json
// @Product json
// @Param user_id header string true "User ID"
// @Router /api/v1/api_key/list  [get]
// @Success 200
func GetApiKeyList(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.Query("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	apiKeys, err := repository.SelectApiKeyListByUserId(userId)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, apiKeys)
}

// GetApiKey
// @Tags GetApiKey
// @Description GetApiKey
// @Accept json
// @Product json
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key  [get]
// @Success 200
func GetApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	apiKeyStr := ctx.Query("api_key")
	if apiKeyStr == "" {
		response.FailCode(ctx, 400, "api_key is required")
		return
	}
	apiKeyRes, err := repository.SelectApiKeyByApiKey(apiKeyStr)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, apiKeyRes)
}

// UpdateApiKey
// @Tags UpdateApiKey
// @Description UpdateApiKey
// @Accept json
// @Product json
// @Param user_id header string true "User ID"
// @Param uploadApiKeyRequest  body  model.UploadApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key  [put]
// @Success 200
func UpdateApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	request := model.UploadApiKeyRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	apikey, err := convertUploadRequestToApiKey(request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	err = repository.UpdateApiKey(apikey)
}

// ApplyApiKey
// @Tags ApplyApiKey
// @Description ApplyApiKey
// @Accept json
// @Product json
// @Param applyApiKeyRequest  body  model.ApplyApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key  [post]
// @Success 200
func ApplyApiKey(ctx *gin.Context) {

	response := model.GetResponse()
	request := model.ApplyApiKeyRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	apiKeyModule := model.ApiKeyModel{
		UserId:  request.UserId,
		KeyName: request.ApiKeyName,
	}
	apiKeySecret := uuid.New().String()
	apiKeyModule.ApiKey = apiKeySecret

	err = repository.InsertApiKey(apiKeyModule)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, gin.H{})

}

// DeleteApiKey
// @Tags DeleteApiKey
// @Description DeleteApiKey
// @Accept json
// @Product json
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key  [delete]
// @Success 200
func DeleteApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	apiKeyStr := ctx.Query("api_key")
	if apiKeyStr == "" {
		response.FailCode(ctx, 400, "api_key is required")
		return
	}
	err := repository.DeleteApiKeyByApiKey(apiKeyStr)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, gin.H{})

}
func convertUploadRequestToApiKey(uploadRequest model.UploadApiKeyRequest) (apikey model.ApiKeyModel, err error) {
	apikey = model.ApiKeyModel{
		UserId:  uploadRequest.UserId,
		ApiKey:  uploadRequest.ApiKey,
		KeyName: uploadRequest.ApiKeyName,
	}
	//if len(uploadRequest.ExtraInfo) > 0 {
	//	extraInfo, err := json.Marshal(uploadRequest.ExtraInfo)
	//	if err != nil {
	//		return apikey, err
	//	}
	//	apikey.Extra = extraInfo
	//}
	return apikey, nil
}

package controller

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type APIKeyVo struct {
	ProjectCode                   string        `json:"project_code"`
	Disable                       bool          `json:"disable"`
	UserId                        int64         `json:"user_id"`
	ApiKey                        string        `json:"api_key"`
	KeyName                       string        `json:"key_name"`
	NetWorkLimitEnable            bool          `json:"network_limit_enable"`
	DomainWhitelist               []interface{} `json:"domain_whitelist"`
	IPWhiteList                   []interface{} `json:"ip_white_list"`
	PaymasterEnable               bool          `json:"paymaster_enable"`
	Erc20PaymasterEnable          bool          `json:"erc20_paymaster_enable"`
	ProjectSponsorPaymasterEnable bool          `json:"project_sponsor_paymaster_enable"`
	UserPayPaymasterEnable        bool          `json:"user_pay_paymaster_enable"`
}

func convertAPiKeyDBModelToAPIKeyVo(apiKeyModel *model.ApiKeyModel) (*APIKeyVo, error) {
	apiKeyVo := APIKeyVo{
		ProjectCode: apiKeyModel.ProjectCode,
		Disable:     apiKeyModel.Disable,
		UserId:      apiKeyModel.UserId,
		ApiKey:      apiKeyModel.ApiKey,
		KeyName:     apiKeyModel.KeyName,
	}
	extra := make(map[string]any)
	if apiKeyModel.Extra == nil {
		return &apiKeyVo, nil
	}
	logrus.Info("apiKeyModel.Extra", apiKeyModel.Extra)
	err := json.Unmarshal(apiKeyModel.Extra, &extra)
	if err != nil {
		return &apiKeyVo, err
	}
	if extra["network_limit_enable"] != nil {

		apiKeyVo.NetWorkLimitEnable = extra["network_limit_enable"].(bool)

	}
	if extra["domain_whitelist"] != nil {
		apiKeyVo.DomainWhitelist = extra["domain_whitelist"].([]interface{})
	}
	if extra["ip_white_list"] != nil {
		iPWhiteList := extra["ip_white_list"].([]interface{})

		apiKeyVo.IPWhiteList = iPWhiteList
	}
	if extra["paymaster_enable"] != nil {
		apiKeyVo.PaymasterEnable = extra["paymaster_enable"].(bool)
	}
	if extra["erc20_paymaster_enable"] != nil {
		apiKeyVo.Erc20PaymasterEnable = extra["erc20_paymaster_enable"].(bool)
	}
	if extra["project_sponsor_paymaster_enable"] != nil {
		apiKeyVo.ProjectSponsorPaymasterEnable = extra["project_sponsor_paymaster_enable"].(bool)
	}
	if extra["user_pay_paymaster_enable"] != nil {
		apiKeyVo.UserPayPaymasterEnable = extra["user_pay_paymaster_enable"].(bool)
	}
	return &apiKeyVo, nil
}

// GetApiKeyList
// @Tags GetApiKeyList
// @Description GetApiKeyList
// @Accept json
// @Product json
// @Router /api/v1/api_key/list  [get]
// @Success 200
// @Security JWT
func GetApiKeyList(ctx *gin.Context) {
	response := model.GetResponse()

	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}

	apiKeyDBModelList, err := repository.SelectApiKeyListByUserId(userId)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	apiKeyVos := make([]APIKeyVo, 0)
	for _, apiKeyDBModel := range apiKeyDBModelList {
		apiKeyVo, err := convertAPiKeyDBModelToAPIKeyVo(&apiKeyDBModel)
		if err != nil {
			response.FailCode(ctx, 500, err.Error())
			return
		}
		apiKeyVos = append(apiKeyVos, *apiKeyVo)
	}
	response.WithDataSuccess(ctx, apiKeyVos)
}

// GetApiKey
// @Tags GetApiKey
// @Description GetApiKey
// @Accept json
// @Product json
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key  [get]
// @Success 200
// @Security JWT
func GetApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	apiKeyStr := ctx.Query("api_key")
	if apiKeyStr == "" {
		response.FailCode(ctx, 400, "api_key is required")
		return
	}
	apiKeyDBModel, err := repository.FindApiKeyByApiKey(apiKeyStr)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}

	apiKeyVo, err := convertAPiKeyDBModelToAPIKeyVo(apiKeyDBModel)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, apiKeyVo)
}

// UpdateApiKey
// @Tags UpdateApiKey
// @Description UpdateApiKey
// @Accept json
// @Product json
// @Param uploadApiKeyRequest  body  model.UploadApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key  [put]
// @Success 200
// @Security JWT
func UpdateApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	request := model.UploadApiKeyRequest{}
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}

	apikey, err := convertUploadRequestToApiKey(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	apikey.UserId, err = strconv.ParseInt(userId, 10, 64)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	//timeNow := time.Now()
	//timeStr := timeNow.Format(global_const.TimeFormat)
	//apikey.UpdatedAt = timeStr
	err = repository.UpdateApiKey(apikey)
	response.WithDataSuccess(ctx, apikey)
}

// ApplyApiKey
// @Tags ApplyApiKey
// @Description ApplyApiKey
// @Accept json
// @Product json
// @Param applyApiKeyRequest  body  model.ApplyApiKeyRequest true "UploadApiKeyRequest Model"
// @Router /api/v1/api_key/apply  [post]
// @Success 200
// @Security JWT
func ApplyApiKey(ctx *gin.Context) {

	response := model.GetResponse()
	request := model.ApplyApiKeyRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	apiKeyModule := model.ApiKeyModel{
		UserId:  userIDInt,
		KeyName: request.ApiKeyName,
	}
	apiKeySecret := uuid.New().String()
	apiKeyModule.ApiKey = apiKeySecret
	extra := make(map[string]any)
	if len(request.DomainWhitelist) > 0 {
		extra["domain_whitelist"] = request.DomainWhitelist
	}

	if len(request.IPWhiteList) > 0 {
		extra["ip_white_list"] = request.IPWhiteList
	}
	extra["network_limit_enable"] = request.NetWorkLimitEnable
	extra["paymaster_enable"] = request.PaymasterEnable
	extra["erc20_paymaster_enable"] = request.Erc20PaymasterEnable
	extra["project_sponsor_paymaster_enable"] = request.ProjectSponsorPaymasterEnable
	extra["user_pay_paymaster_enable"] = request.UserPayPaymasterEnable

	extraJson, err := json.Marshal(extra)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	apiKeyModule.Extra = extraJson

	err = repository.CreateApiKey(&apiKeyModule)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, apiKeyModule)

}

// DeleteApiKey
// @Tags DeleteApiKey
// @Description DeleteApiKey
// @Accept json
// @Product json
// @Param api_key query string true "Api Key"
// @Router /api/v1/api_key  [delete]
// @Success 200
// @Security JWT
func DeleteApiKey(ctx *gin.Context) {
	response := model.GetResponse()
	apiKeyStr := ctx.Query("api_key")
	if apiKeyStr == "" {
		response.SetHttpCode(http.StatusBadRequest).FailCode(ctx, 400, "api_key is required")
		return
	}
	err := repository.DeleteApiKeyByApiKey(apiKeyStr)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, gin.H{})

}
func convertUploadRequestToApiKey(uploadRequest *model.UploadApiKeyRequest) (*model.ApiKeyModel, error) {
	apikey := model.ApiKeyModel{
		ApiKey:  uploadRequest.ApiKey,
		KeyName: uploadRequest.ApiKeyName,
	}
	extraMap := make(map[string]any)

	extraMap["network_limit_enable"] = uploadRequest.NetWorkLimitEnable
	if len(uploadRequest.DomainWhitelist) > 0 {
		extraMap["domain_whitelist"] = uploadRequest.DomainWhitelist
	}

	if len(uploadRequest.IPWhiteList) > 0 {
		extraMap["ip_white_list"] = uploadRequest.IPWhiteList
	}
	extraMap["paymaster_enable"] = uploadRequest.PaymasterEnable
	extraMap["erc20_paymaster_enable"] = uploadRequest.Erc20PaymasterEnable
	extraMap["project_sponsor_paymaster_enable"] = uploadRequest.ProjectSponsorPaymasterEnable
	extraMap["user_pay_paymaster_enable"] = uploadRequest.UserPayPaymasterEnable
	extraMapJson, err := json.Marshal(extraMap)
	if err != nil {
		return nil, err
	}
	apikey.Extra = extraMapJson
	return &apikey, nil
}

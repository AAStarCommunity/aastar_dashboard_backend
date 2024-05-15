package model

type UploadStrategyRequest struct {
	StrategyCode       string            `json:"strategy_code"`
	ProjectCode        string            `json:"project_code"`
	StrategyName       string            `json:"strategy_name"`
	Status             string            `json:"status"`
	ExecuteRestriction map[string]string `json:"strategy_execute_restriction"`
	Extra              map[string]string `json:"extra"`
}

type UploadApiKeyRequest struct {
	ApiKey     string `json:"api_key"`
	ApiKeyName string `json:"api_key_name"`
}

type ApplyApiKeyRequest struct {
	ApiKeyName string `json:"api_key_name"`
}

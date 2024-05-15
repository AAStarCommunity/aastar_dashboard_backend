package model

type UploadStrategyRequest struct {
	StrategyCode       string            `json:"strategy_code"`
	ProjectCode        string            `json:"project_code"`
	StrategyName       string            `json:"strategy_name"`
	UserId             string            `json:"user_id"`
	Status             string            `json:"status"`
	ExecuteRestriction map[string]string `json:"strategy_execute_restriction"`
	Extra              map[string]string `json:"extra"`
}

type UploadApiKeyRequest struct {
	ApiKey     string `json:"api_key"`
	UserId     string `json:"user_id"`
	ApiKeyName string `json:"api_key_name"`
}

type ApplyApiKeyRequest struct {
	UserId     string `json:"user_id"`
	ApiKeyName string `json:"api_key_name"`
}

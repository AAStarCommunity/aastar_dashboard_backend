package model

type UploadStrategyRequest struct {
	StrategyCode       string         `json:"strategy_code"`
	ProjectCode        string         `json:"project_code"`
	StrategyName       string         `json:"strategy_name"`
	UserId             string         `json:"user_id"`
	Status             string         `json:"status"`
	ExecuteRestriction map[string]any `json:"strategy_execute_restriction"`
}

type UploadApiKeyRequest struct {
	ApiKey     string            `json:"api_key"`
	UserId     string            `json:"user_id"`
	ApiKeyName string            `json:"api_key_name"`
	ExtraInfo  map[string]string `json:"extra_info"`
}

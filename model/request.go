package model

type UploadStrategyRequest struct {
	StrategyCode     string   `json:"strategy_code"`
	ProjectCode      string   `json:"project_code"`
	StrategyName     string   `json:"strategy_name"`
	Description      string   `json:"description"`
	ChainIdWhitelist []string `json:"chain_id_whitelist"`
	AddressBlockList []string `json:"address_block_list"`
	StartTime        int64    `json:"start_time"`
	EndTime          int64    `json:"end_time"`
	GlobalMaxUSD     float32  `json:"global_max_usd"`
	DayMaxUSD        float32  `json:"day_max_usd"`
	PerUserMaxUsd    float32  `json:"per_user_max_usd"`
}

type UploadApiKeyRequest struct {
	ApiKey     string `json:"api_key"`
	ApiKeyName string `json:"api_key_name"`
}

type ApplyApiKeyRequest struct {
	ApiKeyName string `json:"api_key_name"`
}

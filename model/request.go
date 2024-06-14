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
	ApiKey                        string   `json:"api_key"`
	ApiKeyName                    string   `json:"api_key_name"`
	NetWorkLimitEnable            bool     `json:"network_limit_enable"`
	DomainWhitelist               []string `json:"domain_whitelist"`
	IPWhiteList                   []string `json:"ip_white_list"`
	PaymasterEnable               bool     `json:"paymaster_enable"`
	Erc20PaymasterEnable          bool     `json:"erc20_paymaster_enable"`
	ProjectSponsorPaymasterEnable bool     `json:"project_sponsor_paymaster_enable"`
	UserPayPaymasterEnable        bool     `json:"user_pay_paymaster_enable"`
}

type ApplyApiKeyRequest struct {
	ApiKeyName                    string   `json:"api_key_name"`
	NetWorkLimitEnable            bool     `json:"network_limit_enable"`
	DomainWhitelist               []string `json:"domain_whitelist"`
	IPWhiteList                   []string `json:"ip_white_list"`
	PaymasterEnable               bool     `json:"paymaster_enable"`
	Erc20PaymasterEnable          bool     `json:"erc20_paymaster_enable"`
	ProjectSponsorPaymasterEnable bool     `json:"project_sponsor_paymaster_enable"`
	UserPayPaymasterEnable        bool     `json:"user_pay_paymaster_enable"`
}

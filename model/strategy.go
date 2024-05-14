package model

import "math/big"

type PaymasterStrategy struct {
	Id                 string                     `gorm:"id"`
	StrategyCode       string                     `json:"strategy_code"`
	ProjectCode        string                     `json:"project_code"`
	StrategyName       string                     `json:"strategy_name"`
	ExecuteRestriction StrategyExecuteRestriction `gorm:"id" json:"execute_restriction"`
	Extra              Extra                      `json:"extra"`
}

type StrategyExecuteRestriction struct {
	BanSenderAddress   string   `json:"ban_sender_address"`
	EffectiveStartTime *big.Int `json:"effective_start_time"`
	EffectiveEndTime   *big.Int `json:"effective_end_time"`
	GlobalMaxUSD       int64    `json:"global_max_usd"`
	GlobalMaxOpCount   int64    `json:"global_max_op_count"`
	DayMaxUSD          int64    `json:"day_max_usd"`
	StartTime          int64    `json:"start_time"`
	EndTime            int64    `json:"end_time"`
}
type Extra struct {
}

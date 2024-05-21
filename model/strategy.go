package model

import (
	"aastar_dashboard_back/env"
	"aastar_dashboard_back/global_const"
	"gorm.io/datatypes"
	"math/big"
)

type PaymasterStrategy struct {
	BaseData
	Description        string                      `gorm:"type:varchar(500)" json:"description"`
	StrategyCode       string                      `gorm:"type:varchar(255)" json:"strategy_code"`
	ProjectCode        string                      `gorm:"type:varchar(255)" json:"project_code"`
	StrategyName       string                      `gorm:"type:varchar(255)" json:"strategy_name"`
	UserId             int64                       `gorm:"type:integer" json:"user_id"`
	Status             global_const.StrategyStatus `gorm:"type:varchar(20);default:disable" json:"status"`
	ExecuteRestriction datatypes.JSON              `gorm:"type:json" json:"execute_restriction"`
	Extra              datatypes.JSON              `gorm:"type:json" json:"extra"`
}

func (PaymasterStrategy) TableName() string {
	if env.Environment.IsProduction() {
		return "aastar_strategy_prod"
	}
	return "aastar_strategy_dev"
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

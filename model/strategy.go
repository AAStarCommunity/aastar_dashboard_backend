package model

import (
	"gorm.io/datatypes"
	"math/big"
	"time"
)

type PaymasterStrategy struct {
	Id                 int64          `gorm:"type:int;primary_key"`
	StrategyCode       string         `gorm:"type:varchar(255)" json:"strategy_code"`
	ProjectCode        string         `gorm:"type:varchar(255)" json:"project_code"`
	StrategyName       string         `gorm:"type:varchar(255)" json:"strategy_name"`
	ExecuteRestriction datatypes.JSON `gorm:"type:json" json:"execute_restriction"`
	Extra              Extra          `json:"extra"`
	CreateAt           time.Time      `gorm:"column:created_at"`
	UpdateAt           time.Time      `gorm:"column:update_at"`
}

func (PaymasterStrategy) TableName() string {
	return "paymaster_strategy"
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

package model

import (
	"gorm.io/datatypes"
	"time"
)

type ApiKeyModel struct {
	ID          int            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProjectCode string         `gorm:"column:project_code" json:"project_code"`
	UserId      string         `gorm:"column:user_id" json:"user_id"`
	ApiKey      string         `gorm:"column:api_key" json:"api_key"`
	KeyName     string         `gorm:"column:key_name" json:"key_name"`
	Extra       datatypes.JSON `gorm:"column:extra" json:"extra"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func (ApiKeyModel) TableName() string {
	return "paymaster_api_key"
}

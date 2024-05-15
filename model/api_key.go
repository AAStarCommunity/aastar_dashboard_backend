package model

import "gorm.io/datatypes"

type ApiKeyModel struct {
	ID        int            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	ProjectId string         `gorm:"column:project_id" json:"project_id"`
	UserId    string         `gorm:"column:user_id" json:"user_id"`
	ApiKey    string         `gorm:"column:api_key" json:"api_key"`
	KeyName   string         `gorm:"column:key_name" json:"key_name"`
	Extra     datatypes.JSON `gorm:"column:extra" json:"extra"`
	CreatedAt string         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string         `gorm:"column:updated_at" json:"updated_at"`
}

func (ApiKeyModel) TableName() string {
	return "paymaster_api_key"
}

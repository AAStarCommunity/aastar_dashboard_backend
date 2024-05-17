package model

import (
	"aastar_dashboard_back/config"
	"gorm.io/datatypes"
)

type ApiKeyModel struct {
	BaseData
	ProjectCode string         `gorm:"column:project_code;type:varchar(50)" json:"project_code"`
	UserId      string         `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	ApiKey      string         `gorm:"column:api_key;type:varchar(255)" json:"api_key"`
	KeyName     string         `gorm:"column:key_name;type:varchar(255)" json:"key_name"`
	Extra       datatypes.JSON `gorm:"column:extra" json:"extra"`
}

func (ApiKeyModel) TableName() string {
	if config.Environment.IsProduction() {
		return "aastar_api_key_prod"
	}
	return "aastar_api_key_dev"
}

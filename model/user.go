package model

import (
	"aastar_dashboard_back/env"
	"gorm.io/datatypes"
)

// User model
// UserId Is Unique Id
type User struct {
	BaseData
	Email           string         `gorm:"column:email;type:varchar(255)" json:"email"`
	PassWord        string         `gorm:"column:password;type:varchar(255)" json:"password"`
	GithubId        int            `gorm:"column:github_id;type:int" json:"github_id"`
	GithubAvatarUrl string         `gorm:"column:github_avatar_url;type:varchar(255)" json:"github_avatar_url"`
	GithubName      string         `gorm:"column:github_name;type:varchar(255)" json:"github_name"`
	GithubLogin     string         `gorm:"column:github_login;type:varchar(255)" json:"github_login"`
	Extra           datatypes.JSON `gorm:"column:extra" json:"extra"`
}

func (User) TableName() string {
	if env.Environment.IsProduction() {
		return "aastar_user_prod"
	}
	return "aastar_user_dev"
}

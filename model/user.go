package model

import "gorm.io/datatypes"

type User struct {
	BaseData
	UserId          string         `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	Email           string         `gorm:"column:email;type:varchar(255)" json:"email"`
	PassWord        string         `gorm:"column:password;type:varchar(255)" json:"password"`
	GithubId        int            `gorm:"column:github_id;type:int" json:"github_id"`
	GitHubAvatarUrl string         `gorm:"column:github_avatar_url;type:varchar(255)" json:"github_avatar_url"`
	GitHubName      string         `gorm:"column:github_name;type:varchar(255)" json:"github_name"`
	GitHubLogin     string         `gorm:"column:github_login;type:varchar(255)" json:"github_login"`
	Extra           datatypes.JSON `gorm:"column:extra" json:"extra"`
}

func (User) TableName() string {
	return "aastar_user"
}

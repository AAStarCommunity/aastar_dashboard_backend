package repository

import (
	"aastar_dashboard_back/model"
	"errors"
	"gorm.io/gorm"
)

func FindByUserId(userId string) (interface{}, error) {
	return nil, nil
}

func FindUserByGitHubId(githubId int) (user *model.User, err error) {
	user = &model.User{}
	tx := dataBase.Where("github_id = ?", githubId).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, tx.Error
		}
		return user, tx.Error
	}
	return nil, nil
}
func UpdateUserLatestLoginTime(user *model.User) error {
	return nil
}
func CreateUser(user *model.User) error {
	tx := dataBase.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

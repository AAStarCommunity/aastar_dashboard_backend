package repository

import (
	"aastar_dashboard_back/model"
	"golang.org/x/xerrors"
)

func SelectApiKeyListByUserId(userId string) (apikeys []model.ApiKeyModel, err error) {
	apikeys = make([]model.ApiKeyModel, 0)
	tx := DataBase.Model(&model.ApiKeyModel{}).Where("user_id = ?", userId).Find(&apikeys)
	if tx.Error != nil {
		return apikeys, xerrors.Errorf("error when finding apikeys: %w", tx.Error)
	}
	return apikeys, nil
}
func DeleteApiKeyByApiKey(apiKey string) (err error) {
	tx := DataBase.Where("api_key = ?", apiKey).Delete(&model.ApiKeyModel{})
	if tx.Error != nil {
		return xerrors.Errorf("error when deleting apikey: %w", tx.Error)
	}
	return nil
}
func SelectApiKeyByApiKey(apiKey string) (apikey model.ApiKeyModel, err error) {
	tx := DataBase.Where("api_key = ?", apiKey).First(&apikey)
	if tx.Error != nil {
		return apikey, xerrors.Errorf("error when finding apikey: %w", tx.Error)
	}
	return apikey, nil
}
func InsertApiKey(apikey model.ApiKeyModel) (err error) {
	tx := DataBase.Create(&apikey)
	if tx.Error != nil {
		return xerrors.Errorf("error when inserting apikey: %w", tx.Error)
	}
	return nil
}

func UpdateApiKey(apikey model.ApiKeyModel) (err error) {
	tx := DataBase.Model(&model.ApiKeyModel{}).Where("api_key = ?", apikey.ApiKey).Updates(apikey)
	if tx.Error != nil {
		return xerrors.Errorf("error when updating apikey: %w", tx.Error)
	}
	return nil
}

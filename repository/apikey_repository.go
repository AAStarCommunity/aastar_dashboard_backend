package repository

import (
	"aastar_dashboard_back/model"
	"golang.org/x/xerrors"
)

func SelectApiKeyListByUserId(userId string) (apikeys []model.ApiKeyModel, err error) {
	apikeys = make([]model.ApiKeyModel, 0)
	tx := dataBase.Model(&model.ApiKeyModel{}).Where("user_id = ?", userId).Find(&apikeys)
	if tx.Error != nil {
		return apikeys, xerrors.Errorf("error when finding apikeys: %w", tx.Error)
	}
	return apikeys, nil
}

// DeleteApiKeyByApiKey  TODO soft DELETE
func DeleteApiKeyByApiKey(apiKey string) (err error) {
	tx := dataBase.Where("api_key = ?", apiKey).Delete(&model.ApiKeyModel{})
	if tx.Error != nil {
		return xerrors.Errorf("error when deleting apikey: %w", tx.Error)
	}
	return nil
}
func FindApiKeyByApiKey(apiKey string) (apikey *model.ApiKeyModel, err error) {
	apikey = &model.ApiKeyModel{}
	tx := dataBase.Where("api_key = ?", apiKey).First(&apikey)
	if tx.Error != nil {
		return apikey, xerrors.Errorf("error when finding apikey: %w", tx.Error)
	}
	return apikey, nil
}
func CreateApiKey(apikey *model.ApiKeyModel) (err error) {
	tx := dataBase.Create(&apikey)
	if tx.Error != nil {
		return xerrors.Errorf("error when inserting apikey: %w", tx.Error)
	}
	return nil
}

func UpdateApiKey(apikey *model.ApiKeyModel) (err error) {
	tx := dataBase.Model(&model.ApiKeyModel{}).Where("api_key = ?", apikey.ApiKey).Updates(apikey)
	if tx.Error != nil {
		return xerrors.Errorf("error when updating apikey: %w", tx.Error)
	}
	return nil
}

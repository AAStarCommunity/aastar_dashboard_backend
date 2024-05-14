package repository

import (
	"aastar_dashboard_back/model"
	"golang.org/x/xerrors"
)

func SelectListByUserId(userId string) (strategies []model.PaymasterStrategy, err error) {
	strategies = make([]model.PaymasterStrategy, 0)
	tx := DataBase.Model(&model.PaymasterStrategy{}).Where("user_id = ?", userId).Find(&strategies)
	if tx.Error != nil {
		return strategies, xerrors.Errorf("error when finding strategies: %w", tx.Error)
	}
	return strategies, nil
}

func DeleteByStrategyCode(strategyCode string) (err error) {
	tx := DataBase.Where("strategy_code = ?", strategyCode).Delete(&model.PaymasterStrategy{})
	if tx.Error != nil {
		return xerrors.Errorf("error when deleting strategy: %w", tx.Error)
	}
	return nil
}
func SelectByStrategyCode(strategyCode string) (strategy model.PaymasterStrategy, err error) {
	tx := DataBase.Where("strategy_code = ?", strategyCode).First(&strategy)
	if tx.Error != nil {
		return strategy, xerrors.Errorf("error when finding strategy: %w", tx.Error)
	}
	return strategy, nil
}
func InsertStrategy(strategy model.PaymasterStrategy) (err error) {
	tx := DataBase.Create(&strategy)
	if tx.Error != nil {
		return xerrors.Errorf("error when inserting strategy: %w", tx.Error)
	}
	return nil
}
func UpdateStrategy(strategy model.PaymasterStrategy) (err error) {
	tx := DataBase.Save(&strategy)
	if tx.Error != nil {
		return xerrors.Errorf("error when updating strategy: %w", tx.Error)
	}
	return nil
}

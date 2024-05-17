package repository

import (
	"aastar_dashboard_back/model"
	"errors"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

func SelectListByUserId(userId string) (strategies []model.PaymasterStrategy, err error) {
	strategies = make([]model.PaymasterStrategy, 0)
	tx := dataBase.Model(&model.PaymasterStrategy{}).Where("user_id = ?", userId).Where("deleted_at IS NULL").Find(&strategies)
	if tx.Error != nil {
		return strategies, xerrors.Errorf("error when finding strategies: %w", tx.Error)
	}
	return strategies, nil
}

// DeleteByStrategyCode TODO soft DELETE
func DeleteByStrategyCode(strategyCode string) (err error) {
	tx := dataBase.Where("strategy_code = ?", strategyCode).Delete(&model.PaymasterStrategy{})
	if tx.Error != nil {
		return xerrors.Errorf("error when deleting strategy: %w", tx.Error)
	}
	return nil
}
func FindByStrategyCode(strategyCode string) (strategy *model.PaymasterStrategy, err error) {
	strategy = &model.PaymasterStrategy{}
	tx := dataBase.Where("strategy_code = ?", strategyCode).Where("deleted_at IS NULL").First(strategy)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, xerrors.Errorf("Can Not Find strategyCode %s", strategyCode)
		}
		return strategy, xerrors.Errorf("error when finding strategy: %w", tx.Error)

	}
	return strategy, nil
}
func CreateStrategy(strategy *model.PaymasterStrategy) (err error) {
	tx := dataBase.Create(&strategy)
	if tx.Error != nil {
		return xerrors.Errorf("error when inserting strategy: %w", tx.Error)
	}
	return nil
}
func UpdateStrategy(strategy *model.PaymasterStrategy) (err error) {
	tx := dataBase.Model(&model.PaymasterStrategy{}).Where("strategy_code = ?", strategy.StrategyCode).Updates(strategy)
	if tx.Error != nil {
		return xerrors.Errorf("error when updating strategy: %w", tx.Error)
	}
	return nil
}

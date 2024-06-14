package data_view_repository

import (
	"aastar_dashboard_back/global_const"
	"aastar_dashboard_back/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type UserSponsorBalanceDBModel struct {
	model.BaseData
	PayUserId        string         `gorm:"type:varchar(255);index" json:"pay_user_id"`
	SponsoredBalance model.BigFloat `gorm:"type:numeric(30,18)" json:"sponsored_balance"`
	AvailableBalance model.BigFloat `gorm:"type:numeric(30,18)" json:"available_balance"`
	LockBalance      model.BigFloat `gorm:"type:numeric(30,18)" json:"lock_balance"`
	Source           string         `gorm:"type:varchar(255)" json:"source"`
	SponsorAddress   string         `gorm:"type:varchar(255)" json:"sponsor_address"`
	IsTestNet        bool           `gorm:"type:boolean" json:"is_test_net"`
}

func (UserSponsorBalanceDBModel) TableName() string {
	return "relay_user_sponsor_balance"
}

type UserSponsorBalanceUpdateLogDBModel struct {
	model.BaseData
	PayUserId  string         `gorm:"type:varchar(255);index" json:"pay_user_id"`
	Amount     model.BigFloat `gorm:"type:numeric(30,18)" json:"amount"`
	UpdateType string         `gorm:"type:varchar(20)" json:"update_type"`
	UserOpHash string         `gorm:"type:varchar(255)" json:"user_op_hash"`
	TxHash     string         `gorm:"type:varchar(255)" json:"tx_hash"`
	TxInfo     datatypes.JSON `gorm:"type:json" json:"tx_info"`
	Source     string         `gorm:"type:varchar(255)" json:"source"`
	IsTestNet  bool           `gorm:"type:boolean" json:"is_test_net"`
}

func (UserSponsorBalanceUpdateLogDBModel) TableName() string {
	return "relay_user_sponsor_balance_update_log"
}

type PaymasterRecallLogDbModel struct {
	model.BaseData
	ProjectUserId   int64          `gorm:"column:project_user_id;type:integer" json:"project_user_id"`
	ProjectApikey   string         `gorm:"column:project_apikey;type:varchar(255)" json:"project_apikey"`
	PaymasterMethod string         `gorm:"column:paymaster_method;type:varchar(25)" json:"paymaster_method"`
	SendTime        string         `gorm:"column:send_time;type:varchar(50)" json:"send_time"`
	Latency         int64          `gorm:"column:latency;type:integer" json:"latency"`
	RequestBody     string         `gorm:"column:request_body;type:varchar(500)" json:"request_body"`
	ResponseBody    string         `gorm:"column:response_body;type:varchar(1000)" json:"response_body"`
	NetWork         string         `gorm:"column:network;type:varchar(25)" json:"network"`
	Status          int            `gorm:"column:status;type:integer" json:"status"`
	Extra           datatypes.JSON `gorm:"column:extra" json:"extra"`
}

func (*PaymasterRecallLogDbModel) TableName() string {
	return "paymaster_recall_log"
}

func FindUserSponsorModuleByUserid(userId string, isTestNet bool) (balanceModel *UserSponsorBalanceDBModel, err error) {
	balanceModel = &UserSponsorBalanceDBModel{}
	tx := dataVeiewDB.Where("pay_user_id = ?", userId).Where("is_test_net = ?", isTestNet).First(balanceModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return balanceModel, nil
}
func GetDepositAndWithDrawLog(userId string, IsTestNet bool) (models []*UserSponsorBalanceUpdateLogDBModel, err error) {
	tx := dataVeiewDB.Model(&UserSponsorBalanceUpdateLogDBModel{}).Where("pay_user_id = ?", userId).Where("is_test_net = ?", IsTestNet).Where("update_type in (?)", []global_const.UpdateType{global_const.UpdateTypeDeposit, global_const.UpdateTypeWithdraw}).Find(&models)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return models, nil
}
func FindPaymasterRecallLogDetailList(apiKey string) (models []*PaymasterRecallLogDbModel, err error) {
	tx := dataVeiewDB.Model(&PaymasterRecallLogDbModel{}).Where("project_apikey = ?", apiKey).Find(&models)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return models, nil
}

type GetApiKeysRequestDayRequestCountResult struct {
	ProjectApikey string `json:"project_apikey"`
	SuccessCount  int64  `json:"success_count"`
	FailureCount  int64  `json:"failure_count"`
}

func GetApiKeyRequestHealth(apikey string, startTime time.Time, endTime time.Time) (*GetApiKeysRequestDayRequestCountResult, error) {
	var result GetApiKeysRequestDayRequestCountResult
	tx := dataVeiewDB.Model(&PaymasterRecallLogDbModel{}).
		Select("project_apikey, COUNT(CASE WHEN status = 200 THEN 1 END) AS success_count, COUNT(CASE WHEN status != 200 THEN 1 END) AS failure_count").
		Where("created_at >= ? AND created_at < ?", startTime, endTime).
		Where("project_apikey = ?", apikey).
		Group("project_apikey").
		Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

// GetApiKeysRequestDayRequestCount  TODO Optimize (will use middle Table)
func GetApiKeysRequestDayRequestCount(apikeys []string) (*[]GetApiKeysRequestDayRequestCountResult, error) {
	var results []GetApiKeysRequestDayRequestCountResult
	err := dataVeiewDB.Model(&PaymasterRecallLogDbModel{}).
		Select("project_apikey, COUNT(CASE WHEN status = 200 THEN 1 END) AS success_count, COUNT(CASE WHEN status != 200 THEN 1 END) AS failure_count").
		Where("created_at >= ? AND project_apikey IN ?", time.Now().Add(-24*time.Hour), apikeys).
		Group("project_apikey").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return &results, nil
}

type RequestHealth struct {
	Time         time.Time `json:"time"`
	SuccessCount int64     `json:"success_count"`
	FailureCount int64     `json:"failure_count"`
}

func GetRequestHealthDay(userId string, apiKey string, startDay time.Time, endDay time.Time) ([]RequestHealth, error) {
	var results []RequestHealth
	tx := dataVeiewDB.Model(&PaymasterRecallLogDbModel{}).
		Select(`
			DATE(created_at) AS time ,
			COUNT(CASE WHEN status = 200 THEN 1 END) AS success_count,
			COUNT(CASE WHEN status != 200 THEN 1 END) AS failure_count
		`).
		Where("project_user_id = ?", userId).
		Where("created_at >= ? AND created_at < ?", startDay, endDay)
	if apiKey != "" {
		tx = tx.Where("project_apikey = ?", apiKey)
	}
	err := tx.Group("DATE(created_at)").
		Order("DATE(created_at)").
		Scan(&results).Error

	return results, err
}
func GetRequestHealthHour() {

}

type DayMetrics struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func GetSponsorDayMetrics(userId string, startDay time.Time, endDay time.Time) ([]DayMetrics, error) {
	var results []DayMetrics
	dailyUpdates := dataVeiewDB.Table("relay_user_sponsor_balance_update_log").
		Select("DATE(created_at) AS time, "+
			"SUM(CASE WHEN update_type = 'lock' THEN amount ELSE 0 END) AS total_lock, "+
			"SUM(CASE WHEN update_type = 'release' THEN amount ELSE 0 END) AS total_release").
		Where("created_at >= ? AND created_at < ?", startDay, endDay).
		Where("pay_user_id = ?", userId).
		Group("DATE(created_at)").
		Session(&gorm.Session{NewDB: true})

	err := dataVeiewDB.Table("(?) AS daily_updates", dailyUpdates).
		Select("time, SUM(total_lock - total_release) OVER (ORDER BY time) AS value").
		Order("time").
		Scan(&results).Error

	return results, err
}

package controller

import (
	"aastar_dashboard_back/data_view_repository"
	"aastar_dashboard_back/model"
	"github.com/gin-gonic/gin"
)

type transactionLog struct {
	UpdateType string         `gorm:"type:varchar(20)" json:"update_type"`
	IsTestNet  bool           `gorm:"type:boolean" json:"is_test_net"`
	Amount     model.BigFloat `gorm:"type:numeric(30,18)" json:"amount"`
	TxHash     string         `gorm:"type:varchar(255)" json:"tx_hash"`
	Time       string         `gorm:"type:varchar(255)" json:"time"`
}

// DataViewGetSponsorTransactionList
// @Tags DataViewGetSponsorTransactionList
// @Summary DataViewGetSponsorTransactionList
// @Description DataViewGetSponsorTransactionList
// @Accept json
// @Produce json
// @Router /api/v1/data_view/sponsor_transaction_list/[get]
// @Param is_test_net path bool true "is_test_net"
// @Success 200
func DataViewGetSponsorTransactionList(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	isTestNet := ctx.Param("is_test_net")
	isTestNetBool := isTestNet == "true"
	trasnLog, err := data_view_repository.GetDepositAndWithDrawLog(userId, isTestNetBool)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	transactionLogs := make([]transactionLog, 0)
	for _, log := range trasnLog {
		transactionLogs = append(transactionLogs, transactionLog{
			UpdateType: log.UpdateType,
			IsTestNet:  log.IsTestNet,
			Amount:     log.Amount,
			TxHash:     log.TxHash,
			Time:       log.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	response.WithDataSuccess(ctx, transactionLogs)

}

// DataViewGetSponsorTotalBalance
// @Tags DataViewGetSponsorTotalBalance
// @Summary DataViewGetSponsorTotalBalance
// @Description DataViewGetSponsorTotalBalance
// @Accept json
// @Produce json
// @Router /api/v1/data_view/sponsor_total_balance/[get]
// @Param is_test_net path bool true "is_test_net"
// @Success 200
// @Security JWT
func DataViewGetSponsorTotalBalance(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	isTestNet := ctx.Param("is_test_net")
	isTestNetBool := isTestNet == "true"
	sponsorModel, err := data_view_repository.FindUserSponsorModuleByUserid(userId, isTestNetBool)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, struct {
		AvailableBalance string `json:"available_balance"`
	}{
		AvailableBalance: sponsorModel.AvailableBalance.String(),
	})
}

// DataViewApiKeyPaymasterRecallDetailList
// @Tags DataViewApiKeyPaymasterRecallDetailList
// @Summary DataViewApiKeyPaymasterRecallDetailList
// @Description DataViewApiKeyPaymasterRecallDetailList
// @Accept json
// @Produce json
// @Router /api/v1/data/paymaster_requests/[get]
// @Param api_key query string true
// @Success 200
// @Security JWT
func DataViewApiKeyPaymasterRecallDetailList(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	apiKey := ctx.Query("api_key")
	res, err := data_view_repository.FindPaymasterRecallLogDetailList(apiKey)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, res)
}

func DataViewPaymasterMonthRecallAnalysis(ctx *gin.Context) {

}
func DataViewAllApikeyPaymasterDailySuccessRate(ctx *gin.Context) {
	//
}

func DataViewApiKeyPaymasterRecallSuccessRate(ctx *gin.Context) {
	//
}
func DataViewApiSuccessRate(ctx *gin.Context) {
	//
}

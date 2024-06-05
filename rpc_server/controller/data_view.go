package controller

import (
	"aastar_dashboard_back/data_view_repository"
	"aastar_dashboard_back/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

type healthData struct {
	Time       string `json:"time"`
	Successful int    `json:"successful"`
	Failed     int    `json:"failed"`
}

var mockHealData = []healthData{
	{Time: "05/07", Successful: 100, Failed: 10},
	{Time: "05/08", Successful: 200, Failed: 20},
	{Time: "05/09", Successful: 300, Failed: 30},
	{Time: "05/10", Successful: 300, Failed: 30},
	{Time: "05/11", Successful: 300, Failed: 30},
	{Time: "05/12", Successful: 300, Failed: 30},
}

// DataViewRequestHealth
// @Tags DataViewRequestHealth
// @Summary DataViewRequestHealth
// @Description DataViewRequestHealth
// @Accept json
// @Produce json
// @Router /api/v1/data/request_health_list [get]
// @Success 200
// @Security JWT
func DataViewRequestHealth(ctx *gin.Context) {
	response := model.GetResponse()
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, mockHealData)
}

type successRateDate struct {
	Time        string  `json:"time"`
	SuccessRate float32 `json:"success_rate"`
}

var mockSuccessRateData = []successRateDate{
	{Time: "05/07", SuccessRate: 100},
	{Time: "05/08", SuccessRate: 99.9},
	{Time: "05/09", SuccessRate: 99.8},
	{Time: "05/10", SuccessRate: 99.0},
	{Time: "05/11", SuccessRate: 81},
}

// DataViewSuccessRate
// @Tags DataViewSuccessRate
// @Summary DataViewSuccessRate
// @Description DataViewSuccessRate
// @Accept json
// @Produce json
// @Router /api/v1/data/success_rate_list [get]
// @Success 200
// @Security JWT
func DataViewSuccessRate(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, mockSuccessRateData)
}

// DataViewSuccessRateOne
// @Tags DataViewSuccessRateOne
// @Summary DataViewSuccessRateOne
// @Description DataViewSuccessRateOne
// @Accept json
// @Produce json
// @Router /api/v1/data/success_rate_one [get]
// @Success 200
// @Security JWT
func DataViewSuccessRateOne(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, 99.1)
}

// DataViewRequestCountOne
// @Tags DataViewRequestCountOne
// @Summary DataViewRequestCountOne
// @Description DataViewRequestCountOne
// @Accept json
// @Produce json
// @Router /api/v1/data/request_count_one [get]
// @Success 200
// @Security JWT
func DataViewRequestCountOne(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, 100)
}

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
// @Router /api/v1/data/sponsor_transaction_list [get]
// @Param is_test_net query bool true "is_test_net"
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

// DataViewGetBalance
// @Tags DataViewGetBalance
// @Summary DataViewGetBalance
// @Description DataViewGetBalance
// @Accept json
// @Produce json
// @Router /api/v1/data/balance [get]
// @Param is_test_net query bool true "is_test_net"
// @Success 200
// @Security JWT
func DataViewGetBalance(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}

	isTestNet := ctx.Query("is_test_net")
	logrus.Info("isTestNet: ", isTestNet)
	isTestNetBool, err := strconv.ParseBool(isTestNet)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}

	sponsorModel, err := data_view_repository.FindUserSponsorModuleByUserid(userId, isTestNetBool)
	if errors.Is(err, gorm.ErrRecordNotFound) {

		response.WithDataSuccess(ctx, 0)
		return

	}
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	balanceBigFloat := sponsorModel.AvailableBalance
	balanceFloat, _ := balanceBigFloat.Float64()
	response.WithDataSuccess(ctx, balanceFloat)
}

// DataViewApiKeyPaymasterRecallDetailList
// @Tags DataViewApiKeyPaymasterRecallDetailList
// @Summary DataViewApiKeyPaymasterRecallDetailList
// @Description DataViewApiKeyPaymasterRecallDetailList
// @Accept json
// @Produce json
// @Router /api/v1/data/paymaster_requests [get]
// @Param api_key query string true "Api Key"
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

type sponsorMetric struct {
	Time  string  `json:"time"`
	Value float32 `json:"value"`
}

var mockSponsorMetric = []sponsorMetric{

	{Time: "05/07", Value: 100},
	{Time: "05/08", Value: 99.9},
	{Time: "05/09", Value: 99.8},
	{Time: "05/10", Value: 99.0},
	{Time: "05/11", Value: 81},
}

// DataViewSponsoredMetrics
// @Tags DataViewSponsoredMetrics
// @Summary DataViewSponsoredMetrics
// @Description DataViewSponsoredMetrics
// @Accept json
// @Produce json
// @Router /api/v1/data/sponsored_metrics [get]
// @Success 200
// @Security JWT
func DataViewSponsoredMetrics(ctx *gin.Context) {
	response := model.GetResponse()
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, mockSponsorMetric)
}

type apiKeyDataView struct {
	ApiName         string  `json:"api_name"`
	RequestCount    int     `json:"request_count"`
	SuccessRateDate float32 `json:"success_rate"`
	ApiKey          string  `json:"api_key"`
}

var mockapiKeyDataView = []apiKeyDataView{
	{ApiName: "api1", RequestCount: 100, SuccessRateDate: 99, ApiKey: "8bced19b-505e-4d11-ae80-abbee3d3a38c"},
	{ApiName: "api2", RequestCount: 100, SuccessRateDate: 99, ApiKey: "8bced19b-505e-4d11-ae80-abbee3d3a38c"},
}

// DataViewApiKeysOverView
// @Tags DataViewApiKeysOverView
// @Summary DataViewApiKeysOverView
// @Description DataViewApiKeysOverView
// @Accept json
// @Produce json
// @Router /api/v1/data/api_keys_data_overview [get]
// @Success 200
// @Security JWT
func DataViewApiKeysOverView(ctx *gin.Context) {

	response := model.GetResponse()
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, mockapiKeyDataView)
}

type paymasterPayTypeMetric struct {
	Time           string `json:"time"`
	Erc20PayType   int    `json:"erc20_pay_type"`
	ProjectSponsor int    `json:"project_sponsor"`
	UserPay        int    `json:"user_pay"`
}

var mockpaymasterPayTypeMetric = []paymasterPayTypeMetric{
	{Time: "05/07", Erc20PayType: 100, ProjectSponsor: 99, UserPay: 99},
	{Time: "05/08", Erc20PayType: 100, ProjectSponsor: 99, UserPay: 99},
	{Time: "05/09", Erc20PayType: 189, ProjectSponsor: 231, UserPay: 99},
	{Time: "05/10", Erc20PayType: 100, ProjectSponsor: 99, UserPay: 99},
}

// DataViewPaymasterPayTypeMetrics
// @Tags DataViewPaymasterPayTypeMetrics
// @Summary DataViewPaymasterPayTypeMetrics
// @Description DataViewPaymasterPayTypeMetrics
// @Accept json
// @Produce json
// @Router /api/v1/data/paymaster_pay_type_metrics [get]
// @Success 200
// @Security JWT
func DataViewPaymasterPayTypeMetrics(ctx *gin.Context) {
	response := model.GetResponse()
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, mockpaymasterPayTypeMetric)

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

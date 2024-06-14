package controller

import (
	"aastar_dashboard_back/data_view_repository"
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

type healthData struct {
	Time        string  `json:"time"`
	Successful  int64   `json:"successful"`
	Failed      int64   `json:"failed"`
	SuccessRate float32 `json:"success_rate"`
}

// DataViewRequestHealthOneByApiKey
// @Tags DataViewRequestHealthOneByApiKey
// @Summary DataViewRequestHealthOneByApiKey
// @Description DataViewRequestHealthOneByApiKey
// @Accept json
// @Produce json
// @Router /api/v1/data/request_health_one [get]
// @Param api_key query string true "API Key"
// @Param time_type query string true "Time Type"
// @Success 200
// @Security JWT
func DataViewRequestHealthOneByApiKey(ctx *gin.Context) {
	response := model.GetResponse()

	apiKey := ctx.Query("api_key")
	timeType := ctx.Query("time_type")
	if apiKey == "" {
		response.FailCode(ctx, 400, "api_key is required")
		return
	}
	if timeType == "" {
		response.FailCode(ctx, 400, "time_type is required")
		return
	}
	endTime := time.Now()
	var startTime time.Time
	if timeType == "day" {
		startTime = endTime.Add(-24 * time.Hour)
	} else if timeType == "hour" {
		startTime = endTime.Add(-time.Hour)
	} else {
		response.FailCode(ctx, 400, "time_type is not support")
		return
	}
	requestHealth, err := data_view_repository.GetApiKeyRequestHealth(apiKey, startTime, endTime)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.WithDataSuccess(ctx, nil)
		return
	}
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}

	logrus.Debugf("requestHealth: %v", requestHealth)

	successRate := float32(requestHealth.SuccessCount) / float32(requestHealth.SuccessCount+requestHealth.FailureCount) * 100
	logrus.Debugf("successRate: %v", successRate)
	if math.IsNaN(float64(successRate)) {
		successRate = 0

	}
	response.WithDataSuccess(ctx, healthData{
		Successful:  requestHealth.SuccessCount,
		Failed:      requestHealth.FailureCount,
		SuccessRate: successRate,
	})

}

// DataViewRequestHealth
// @Tags DataViewRequestHealth
// @Summary DataViewRequestHealth
// @Description DataViewRequestHealth
// @Accept json
// @Produce json
// @Router /api/v1/data/request_health_list [get]
// @Param api_key query string false "Api Key"
// @Success 200
// @Security JWT
func DataViewRequestHealth(ctx *gin.Context) {
	response := model.GetResponse()

	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	apiKey := ctx.Query("api_key")
	endDay := time.Now().Truncate(24*time.Hour).AddDate(0, 0, 1).Add(-time.Nanosecond)
	startDay := endDay.AddDate(0, 0, -30)
	//TODO 1 allowUserSelect 2. AddHour
	requestHealthList, err := data_view_repository.GetRequestHealthDay(userId, apiKey, startDay, endDay)
	resultList := make([]healthData, 0)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.WithDataSuccess(ctx, resultList)
		return
	}
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	for _, requestHealth := range requestHealthList {
		successRate := float32(requestHealth.SuccessCount) / float32(requestHealth.SuccessCount+requestHealth.FailureCount) * 100
		resultList = append(resultList, healthData{
			Time:        requestHealth.Time.Format("2006-01-02"),
			Successful:  requestHealth.SuccessCount,
			Failed:      requestHealth.FailureCount,
			SuccessRate: successRate,
		})
	}
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, resultList)
}

type successRateDate struct {
	Time        string  `json:"time"`
	SuccessRate float32 `json:"success_rate"`
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

	balanceType := ctx.Query("balance_type")
	if balanceType == "" {
		response.FailCode(ctx, 400, "balance_type is required")
		return
	}
	var balanceRes float64
	if balanceType == "total_sponsored" {
		sponsorModel, err := data_view_repository.FindUserSponsorModuleByUserid(userId, isTestNetBool)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.WithDataSuccess(ctx, 0)
			return
		}
		if err != nil {
			response.FailCode(ctx, 500, err.Error())
			return
		}
		balanceBigFloat := sponsorModel.SponsoredBalance
		balanceRes, _ = balanceBigFloat.Float64()

	} else if balanceType == "sponsor_quota_balance" {
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
		balanceRes, _ = balanceBigFloat.Float64()
	} else {
		response.FailCode(ctx, 400, "balance_type is not support")
		return
	}
	response.WithDataSuccess(ctx, balanceRes)
	return

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
	endDay := time.Now().Truncate(24*time.Hour).AddDate(0, 0, 1).Add(-time.Nanosecond)
	startDay := endDay.AddDate(0, 0, -30)
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	sponsorDayMetrics, err := data_view_repository.GetSponsorDayMetrics(userId, startDay, endDay)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.WithDataSuccess(ctx, nil)
		return
	}
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	sponsorMetrics := make([]sponsorMetric, 0)
	for _, sponsorDayMetric := range sponsorDayMetrics {
		sponsorMetrics = append(sponsorMetrics, sponsorMetric{
			Time:  sponsorDayMetric.Time.Format("2006-01-02"),
			Value: float32(sponsorDayMetric.Value),
		})

	}
	//healthList := make([]healthData, 0)
	//response.WithDataSuccess(ctx, healthList)
	response.WithDataSuccess(ctx, sponsorMetrics)
}

type apiKeyDataView struct {
	ApiName         string  `json:"api_name"`
	RequestCount    int64   `json:"request_count"`
	SuccessRateDate float32 `json:"success_rate"`
	ApiKey          string  `json:"api_key"`
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

	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	achieveAPIkeyModels, err := repository.SelectApiKeyListByUserId(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.WithDataSuccess(ctx, nil)
		return
	}
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}

	apikeyMap := make(map[string]model.ApiKeyModel)
	apiKeys := make([]string, 0)
	for _, apiKeyModel := range achieveAPIkeyModels {
		apikeyMap[apiKeyModel.ApiKey] = apiKeyModel
		apiKeys = append(apiKeys, apiKeyModel.ApiKey)
	}
	logrus.Debugf("apiKeys: %v", apiKeys)

	viewResults, err := data_view_repository.GetApiKeysRequestDayRequestCount(apiKeys)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiKeyDataViews := make([]apiKeyDataView, 0)
		for key, value := range apikeyMap {
			apiKeyDataViews = append(apiKeyDataViews, apiKeyDataView{
				ApiName:         value.KeyName,
				RequestCount:    0,
				SuccessRateDate: 0,
				ApiKey:          key,
			})
		}
		response.WithDataSuccess(ctx, apiKeyDataViews)
		return
	}
	logrus.Debugf("viewResults: %v", viewResults)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	apiKeyDataViews := make([]apiKeyDataView, 0)
	for _, result := range *viewResults {
		apiKey := result.ProjectApikey
		apiKeyName := apikeyMap[apiKey].KeyName
		succesCount := result.SuccessCount
		failureCount := result.FailureCount
		totalCount := succesCount + failureCount
		successRate := (float32(succesCount) / float32(totalCount)) * 100
		apiKeyDataViews = append(apiKeyDataViews, apiKeyDataView{
			ApiName:         apiKeyName,
			RequestCount:    totalCount,
			SuccessRateDate: successRate,
			ApiKey:          apiKey,
		})
	}
	response.WithDataSuccess(ctx, apiKeyDataViews)
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

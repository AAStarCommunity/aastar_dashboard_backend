package controller

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// GetStrategy
// @Tags GetStrategy
// @Description GetStrategy
// @Accept json
// @Product json
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Router /api/v1/paymaster_strategy  [get]
// @Success 200
// @Security JWT
func GetStrategy(ctx *gin.Context) {
	strategyCode := ctx.Query("strategy_code")
	response := model.GetResponse()
	if strategyCode == "" {
		response.FailCode(ctx, 400, "strategy_code is required")
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	strategy, err := repository.FindByStrategyCode(strategyCode)
	if err != nil {
		logrus.Errorf("error when finding strategy: %v", err)
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, strategy)
}

// AddStrategy
// @Tags AddStrategy
// @Description AddStrategy
// @Accept json
// @Product json
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Router /api/v1/paymaster_strategy  [post]
// @Success 200
// @Security JWT
func AddStrategy(ctx *gin.Context) {
	request := model.UploadStrategyRequest{}
	err := ctx.ShouldBindJSON(&request)
	response := model.GetResponse()
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	strategy, err := convertUploadRequestToStrategy(request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	prefixLen := 5
	if len(userId) < prefixLen {
		prefixLen = len(userId)
	}
	if strategy.StrategyCode == "" {
		strategy.StrategyCode = strategy.StrategyName + "_" + userId[prefixLen:] + "_" + util.GenerateRandomString(5)
	}
	strategy.UserId = userId
	err = repository.CreateStrategy(&strategy)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, strategy)
}

// UpdateStrategy
// @Tags UpdateStrategy
// @Description UpdateStrategy
// @Accept json
// @Product json
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Router /api/v1/paymaster_strategy  [put]
// @Success 200
// @Security JWT
func UpdateStrategy(ctx *gin.Context) {
	response := model.GetResponse()
	request := model.UploadStrategyRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	strategy, err := convertUploadRequestToStrategy(request)
	if err != nil {
		response.FailCode(ctx, 400, err.Error())
		return
	}
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	strategy.UserId = userId
	err = repository.UpdateStrategy(&strategy)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, strategy)
}

// DeleteStrategy
// @Tags DeleteStrategy
// @Description DeleteStrategy
// @Accept json
// @Product json
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Router /api/v1/paymaster_strategy  [delete]
// @Success 200
// @Security JWT
func DeleteStrategy(ctx *gin.Context) {
	response := model.GetResponse()
	strategyCode := ctx.Query("strategy_code")
	if strategyCode == "" {
		response.FailCode(ctx, 400, "strategy_code is required")
		return
	}
	err := repository.DeleteByStrategyCode(strategyCode)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.Success(ctx)
}

// GetStrategyList
// @Tags GetStrategyList
// @Description GetStrategyList
// @Accept json
// @Product json
// @Router /api/v1/paymaster_strategy/list [get]
// @Success 200
// @Security JWT
func GetStrategyList(ctx *gin.Context) {
	response := model.GetResponse()

	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	strategies, err := repository.SelectListByUserId(userId)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, strategies)
}

func convertUploadRequestToStrategy(request model.UploadStrategyRequest) (model.PaymasterStrategy, error) {
	strategy := model.PaymasterStrategy{
		StrategyCode: request.StrategyCode,
		ProjectCode:  request.ProjectCode,
		StrategyName: request.StrategyName,
	}
	executeRestriction := make(map[string]any)
	if len(request.ChainIdWhitelist) != 0 {
		executeRestriction["chain_id_whitelist"] = request.ChainIdWhitelist
	}
	if len(request.AddressBlockList) != 0 {
		executeRestriction["address_block_list"] = request.AddressBlockList
	}
	if request.StartTime > 0 {
		executeRestriction["start_time"] = request.StartTime
	}
	if request.EndTime > 0 {
		executeRestriction["end_time"] = request.EndTime
	}
	if request.GlobalMaxUSD > 0 {
		executeRestriction["global_max_usd"] = request.GlobalMaxUSD
	}
	if request.DayMaxUSD > 0 {
		executeRestriction["day_max_usd"] = request.DayMaxUSD
	}
	if request.PerUserMaxUsd > 0 {
		executeRestriction["per_user_max_usd"] = request.PerUserMaxUsd
	}

	executeRestrictionJson, err := json.Marshal(executeRestriction)
	if err != nil {
		return strategy, xerrors.Errorf("error when marshal execute restriction: %w", err)
	}
	strategy.ExecuteRestriction = executeRestrictionJson
	//if len(request.Extra) != 0 {
	//	extraJson, err := json.Marshal(request.Extra)
	//	if err != nil {
	//		return strategy, xerrors.Errorf("error when marshal extra: %w", err)
	//	}
	//	strategy.Extra = extraJson
	//}
	return strategy, nil
}

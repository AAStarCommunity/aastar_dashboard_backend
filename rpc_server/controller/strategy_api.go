package controller

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

// GetStrategy
// @Tags GetStrategy
// @Description GetStrategy
// @Accept json
// @Product json
// @Param user_id header string true "User ID"
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Router /api/v1/paymaster_strategy  [get]
// @Success 200
func GetStrategy(ctx *gin.Context) {
	strategyCode := ctx.Query("strategy_code")
	response := model.GetResponse()
	if strategyCode == "" {
		response.FailCode(ctx, 400, "strategy_code is required")
		return
	}
	strategy, err := repository.SelectByStrategyCode(strategyCode)
	if err != nil {
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
// @Param user_id header string true "User ID"
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Router /api/v1/paymaster_strategy  [post]
// @Success 200
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
	userId := ctx.GetHeader("user_id")
	prefixLen := 10
	if len(userId) < prefixLen {
		prefixLen = len(userId)
	}

	strategy.StrategyCode = strategy.StrategyName + "_" + userId[prefixLen:] + "_" + util.GenerateRandomString(10)
	strategy.UserId = ctx.GetHeader("user_id")
	err = repository.InsertStrategy(strategy)
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
// @Param user_id header string true "User ID"
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Router /api/v1/paymaster_strategy  [put]
// @Success 200
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
	strategy.UserId = ctx.GetHeader("user_id")
	err = repository.UpdateStrategy(strategy)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, gin.H{})
}

// DeleteStrategy
// @Tags DeleteStrategy
// @Description DeleteStrategy
// @Accept json
// @Product json
// @Param user_id header string true "User ID"
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Router /api/v1/paymaster_strategy  [delete]
// @Success 200
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
// @Param user_id header string true "User ID"
// @Router /api/v1/paymaster_strategy/list [get]
// @Success 200
func GetStrategyList(ctx *gin.Context) {
	response := model.GetResponse()

	userId := ctx.GetHeader("user_id")
	if userId == "" {
		response := model.GetResponse()
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

package controller

import (
	"aastar_dashboard_back/model"
	"github.com/gin-gonic/gin"
)

// GetStrategy
// @Tags GetStrategy
// @Description GetStrategy
// @Accept json
// @Product json
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Param user_id query string true "User ID"
// @Router /api/v1/paymaster_strategy  [get]
// @Success 200
func GetStrategy(ctx *gin.Context) {

	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// AddStrategy
// @Tags AddStrategy
// @Description AddStrategy
// @Accept json
// @Product json
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Param user_id query string true "User ID"
// @Router /api/v1/paymaster_strategy  [post]
// @Success 200
func AddStrategy(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// UpdateStrategy
// @Tags UpdateStrategy
// @Description UpdateStrategy
// @Accept json
// @Product json
// @Param uploadStrategyRequest  body  model.UploadStrategyRequest true "UploadStrategyRequest Model"
// @Param user_id query string true "User ID"
// @Router /api/v1/paymaster_strategy  [put]
// @Success 200
func UpdateStrategy(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// DeleteStrategy
// @Tags DeleteStrategy
// @Description DeleteStrategy
// @Accept json
// @Product json
// @Param strategy_code query string true "PaymasterStrategy Code"
// @Param user_id query string true "User ID"
// @Router /api/v1/paymaster_strategy  [delete]
// @Success 200
func DeleteStrategy(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

// GetStrategyList
// @Tags GetStrategyList
// @Description GetStrategyList
// @Accept json
// @Product json
// @Param user_id query string true "User ID"
// @Router /api/v1/paymaster_strategy/list [get]
// @Success 200
func GetStrategyList(ctx *gin.Context) {
	response := model.GetResponse()
	response.WithDataSuccess(ctx, gin.H{})
}

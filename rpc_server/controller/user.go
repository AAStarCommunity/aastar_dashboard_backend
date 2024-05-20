package controller

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"github.com/gin-gonic/gin"
)

// GetUserInfo
// @Tags GetUserInfo
// @Description GetUserInfo
// @Accept json
// @Product json
// @Router /api/v1/user  [get]
// @Success 200
// @Security JWT
func GetUserInfo(ctx *gin.Context) {
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}

	user, err := repository.FindUserByUserId(userId)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, user)
}

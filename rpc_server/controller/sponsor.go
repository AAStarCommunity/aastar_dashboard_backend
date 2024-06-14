package controller

import (
	"aastar_dashboard_back/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

type DepositSponsorRequest struct {
	Amount           *big.Float
	TxHash           string
	SponsorSignature string
	SponsorAddress   string
	IsTestNet        bool
}
type WithdrawSponsorRequest struct {
	Amount           *big.Float
	SponsorSignature string
	SponsorAddress   string
	IsTestNet        bool
}

func SponsorDeposit(ctx *gin.Context) {

	request := DepositSponsorRequest{}
	response := model.GetResponse()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errStr := fmt.Sprintf("Request Error [%v]", err)
		response.SetHttpCode(http.StatusBadRequest).FailCode(ctx, http.StatusBadRequest, errStr)
		return
	}

}
func SponsorWithdraw(ctx *gin.Context) {
	request := WithdrawSponsorRequest{}
	response := model.GetResponse()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errStr := fmt.Sprintf("Request Error [%v]", err)
		response.SetHttpCode(http.StatusBadRequest).FailCode(ctx, http.StatusBadRequest, errStr)
		return
	}
}

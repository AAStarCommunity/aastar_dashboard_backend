package controller

import (
	"aastar_dashboard_back/config"
	"aastar_dashboard_back/model"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"io"
	"math/big"
	"net/http"
	"time"
)

type DepositSponsorRequest struct {
	TxHash    string `json:"tx_hash"`
	IsTestNet bool   `json:"is_test_net"`
}
type WithdrawSponsorRequest struct {
	Amount    *big.Float `json:"amount"`
	IsTestNet bool       `json:"is_test_net"`
}

type RelaySponsorRequest struct {
	TimeStamp     int64  `json:"time_stamp"`
	TxHash        string `json:"tx_hash"`
	IsTestNet     bool   `json:"is_test_net"`
	PayUserId     string `json:"pay_user_id"`
	DepositSource string `json:"deposit_source"`
}

// SponsorDeposit
// @Summary SponsorDeposit
// @Description SponsorDeposit
// @Tags Sponsor
// @Accept json
// @Product json
// @Router  /api/v1/sponsor/deposit  [post]
// @Param request body DepositSponsorRequest true "DepositSponsorRequest Model"
// @Success 200
// @Security JWT
func SponsorDeposit(ctx *gin.Context) {

	request := DepositSponsorRequest{}
	response := model.GetResponse()
	userId := ctx.GetString("user_id")
	if userId == "" {
		response.FailCode(ctx, 400, "user_id is required")
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errStr := fmt.Sprintf("Request Error [%v]", err)
		response.SetHttpCode(http.StatusBadRequest).FailCode(ctx, http.StatusBadRequest, errStr)
		return
	}
	relayRequest := RelaySponsorRequest{
		TimeStamp:     time.Now().Unix(),
		TxHash:        request.TxHash,
		IsTestNet:     request.IsTestNet,
		PayUserId:     userId,
		DepositSource: "dashboard",
	}
	jsonData, err := json.Marshal(relayRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	hash := sha256.New()
	hash.Write(jsonData)
	hashBytes := hash.Sum(nil)
	signatureByte, err := crypto.Sign(accounts.TextHash(hashBytes), config.GetSignerEoa().PrivateKey)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}
	hashBytesHex := hex.EncodeToString(hashBytes)
	signatureHex := hex.EncodeToString(signatureByte)

	// Sign
	req, err := http.NewRequest("POST",
		config.GetRelayUrl()+"/api/v1/paymaster_sponsor/deposit",
		bytes.NewBuffer(jsonData))
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	req.Header.Set("relay_hash", hashBytesHex)
	req.Header.Set("relay_signature", signatureHex)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.FailCode(ctx, 500, fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Relay Server Error: [%v] - %s", resp.StatusCode, string(body))
		response.FailCode(ctx, 500, msg)
		return
	}

	resp.Body.Close()

	response.WithDataSuccess(ctx, string(body))
}

// SponsorWithdraw
// @Summary SponsorWithdraw
// @Description SponsorWithdraw
// @Tags Sponsor
// @Accept json
// @Product json
// @Router  /api/v1/sponsor/withdraw  [post]
// @Param request body WithdrawSponsorRequest true "WithdrawSponsorRequest Model"
// @Success 200
// @Security JWT
func SponsorWithdraw(ctx *gin.Context) {
	request := WithdrawSponsorRequest{}
	response := model.GetResponse()
	if err := ctx.ShouldBindJSON(&request); err != nil {
		errStr := fmt.Sprintf("Request Error [%v]", err)
		response.SetHttpCode(http.StatusBadRequest).FailCode(ctx, http.StatusBadRequest, errStr)
		return
	}
}
func GetInfoByHash(txHash string, client *ethclient.Client) (*types.Transaction, error) {
	txHashHex := common.HexToHash(txHash)
	//TODO consider about pending
	tx, _, err := client.TransactionByHash(context.Background(), txHashHex)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return nil, xerrors.Errorf("Transaction [%s] not found", txHash)
		}
		return nil, err
	}
	return tx, nil

}

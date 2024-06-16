package controller

import (
	"aastar_dashboard_back/config"
	"aastar_dashboard_back/model"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

type DepositSponsorRequest struct {
	DepositAddress   string `json:"deposit_address"`
	TxHash           string `json:"tx_hash"`
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

type RelaySponsorRequest struct {
	TimeStamp      int64  `json:"time_stamp"`
	DepositAddress string `json:"deposit_address"`
	TxHash         string `json:"tx_hash"`
	IsTestNet      bool   `json:"is_test_net"`
	PayUserId      string `json:"pay_user_id"`
	DepositSource  string `json:"deposit_source"`
}

// SponsorDeposit
// @Summary SponsorDeposit
// @Description SponsorDeposit
// @Tags Sponsor
// @Accept json
// @Product json
// @Param request body DepositSponsorRequest true "DepositSponsorRequest Model"
func SponsorDeposit(ctx *gin.Context) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
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
	//TODO Add Signature Verification
	relayRequest := RelaySponsorRequest{
		TimeStamp:      time.Now().Unix(),
		DepositAddress: request.DepositAddress,
		TxHash:         request.TxHash,
		IsTestNet:      request.IsTestNet,
		PayUserId:      userId,
	}
	jsonData, err := json.Marshal(relayRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	hash := sha256.New()
	hash.Write(jsonData)
	hashBytes := hash.Sum(nil)
	signatureByte, err := crypto.Sign(hashBytes, config.GetSignerEoa().PrivateKey)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}
	hashBytesHex := hex.EncodeToString(hashBytes)
	signatureHex := hex.EncodeToString(signatureByte)

	// Sign
	req, err := http.NewRequest(config.GetRelayUrl(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	req.Header.Set("relay_hash", hashBytesHex)
	req.Header.Set("relay_signature", signatureHex)

	resp, err := client.Do(req)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			response.FailCode(ctx, 500, err.Error())
			return
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.FailCode(ctx, 500, err.Error())
		return
	}
	response.WithDataSuccess(ctx, body)
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
func GetInfoByHash(txHash string, client *ethclient.Client) (*types.Transaction, error) {
	txHashHex := common.HexToHash(txHash)
	//TODO consider about pending
	tx, _, err := client.TransactionByHash(context.Background(), txHashHex)
	if err != nil {
		if err.Error() == "not found" {
			return nil, xerrors.Errorf("Transaction [%s] not found", txHash)
		}
		return nil, err
	}
	return tx, nil

}

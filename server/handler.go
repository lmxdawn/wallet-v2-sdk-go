package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github/lmxdawn/wallet-example/pkg/util"
	"github/lmxdawn/wallet-example/wallet"
	"strconv"
)

func RechargeAddress(c *gin.Context) {

	var q RechargeAddressReq

	if err := c.ShouldBindQuery(&q); err != nil {
		HandleValidatorError(c, err)
		return
	}

	walletV, ok := c.Get("wallet")
	if !ok {
		APIResponse(c, InternalServerError, nil)
		return
	}

	walletClient := walletV.(*wallet.Client)

	callUrl := util.GetUrl(c.Request, "/recharge/call")
	// 获取
	createWallet, err := walletClient.CreateWallet(q.MemberId, q.NetworkName, q.CoinSymbol, callUrl)
	if err != nil {
		fmt.Println("创建地址失败", err)
		APIResponse(c, InternalServerError, nil)
		return
	}

	APIResponse(c, nil, createWallet)

}

// RechargeCall 充值回调
func RechargeCall(c *gin.Context) {

	var q RechargeCallReq

	if err := c.ShouldBindJSON(&q); err != nil {
		APIResponse(c, err, nil)
		return
	}

	// 验签
	signData := make(map[string]string)
	signData["network_name"] = q.NetworkName
	signData["coin_symbol"] = q.CoinSymbol
	signData["decimals"] = q.Decimals
	signData["address"] = q.Address
	signData["amount"] = q.Amount
	signData["business_id"] = q.BusinessId
	signData["block_high"] = q.BlockHigh
	signData["block_hash"] = q.BlockHash
	signData["txid"] = q.Txid
	signData["status"] = strconv.Itoa(q.Status)
	signBool := wallet.VerifySign(signData)
	if !signBool {
		APIResponse(c, InternalServerError, nil)
		return
	}

	// 充值成功
	fmt.Println(q)

}

// WithdrawCreate 获取充值地址
func WithdrawCreate(c *gin.Context) {

	var q WithdrawCreateReq

	if err := c.ShouldBindJSON(&q); err != nil {
		APIResponse(c, err, nil)
		return
	}

	businessId, _ := uuid.NewUUID()

	walletV, ok := c.Get("wallet")
	if !ok {
		APIResponse(c, InternalServerError, nil)
		return
	}

	walletClient := walletV.(*wallet.Client)

	callUrl := util.GetUrl(c.Request, "/withdraw/call")
	withdrawRes, err := walletClient.CreateWithdraw(q.NetworkName, q.CoinSymbol, q.Address, q.Amount, businessId.String(), callUrl)
	if err != nil {
		fmt.Println("网络错误", err)
		APIResponse(c, InternalServerError, nil)
	}

	fmt.Println("请求结果", withdrawRes)

}

// WithdrawCall 提现回调
func WithdrawCall(c *gin.Context) {

	var q WithdrawCallReq

	if err := c.ShouldBindJSON(&q); err != nil {
		APIResponse(c, err, nil)
		return
	}

	// 验签
	signData := make(map[string]string)
	signData["network_name"] = q.NetworkName
	signData["coin_symbol"] = q.CoinSymbol
	signData["decimals"] = q.Decimals
	signData["address"] = q.Address
	signData["amount"] = q.Amount
	signData["business_id"] = q.BusinessId
	signData["block_high"] = q.BlockHigh
	signData["block_hash"] = q.BlockHash
	signData["txid"] = q.Txid
	signData["status"] = strconv.Itoa(q.Status)
	signBool := wallet.VerifySign(signData)
	if !signBool {
		APIResponse(c, InternalServerError, nil)
		return
	}

	fmt.Println("提现回调", q)

}

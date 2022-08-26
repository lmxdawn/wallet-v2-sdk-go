package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/lmxdawn/wallet-example/config"
	"github/lmxdawn/wallet-example/wallet"
)

// Start 启动服务
func Start(isSwag bool, configPath string) {

	conf, err := config.NewConfig(configPath)

	walletClient := wallet.NewWalletClient(conf.Wallet.Appid, conf.Wallet.SecretKey, conf.Wallet.Url)

	server := gin.Default()

	// 中间件
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(SetWallet(walletClient))

	server.GET("/recharge/call", RechargeCall)
	server.GET("/withdraw/call", WithdrawCall)
	server.GET("/recharge/address", RechargeAddress)
	server.POST("/withdraw/create", WithdrawCreate)

	err = server.Run(fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic("start error")
	}

}

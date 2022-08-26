package server

import (
	"github.com/gin-gonic/gin"
	"github/lmxdawn/wallet-example/wallet"
)

// AuthRequired 认证中间件
func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		if token == "" {
			c.Abort()
			APIResponse(c, ErrToken, nil)
		}

	}

}

// SetWallet Wallet
func SetWallet(wallet *wallet.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("wallet", wallet)
	}

}

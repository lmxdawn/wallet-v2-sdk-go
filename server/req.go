package server

type RechargeAddressReq struct {
	MemberId    string `form:"member_id"`    // 网络名称
	NetworkName string `form:"network_name"` // 网络名称
	CoinSymbol  string `form:"coin_symbol"`  // 币种符号
}

type RechargeCallReq struct {
	NetworkName string `json:"network_name"` // 网络名称
	CoinSymbol  string `json:"coin_symbol"`  // 币种符号
	Address     string `json:"address"`      // 地址
	Amount      string `json:"amount"`       // 充值数量
	BusinessId  string `json:"business_id"`  // 业务ID
	BlockHigh   string `json:"block_high"`   // 区块高度
	Txid        string `json:"txid"`         // 区块链交易哈希
	Status      int    `json:"status"`       // 状态（0: 链上打包中，1：充值成功，2：充值失败）
}

type WithdrawCreateReq struct {
	NetworkName string `json:"network_name"` // 网络名称
	CoinSymbol  string `json:"coin_symbol"`  // 币种符号
	Address     string `json:"address"`      // 地址
	Amount      string `json:"amount"`       // 提币数量
	PayPassword string `json:"pay_password"` // 支付密码
}

type WithdrawCallReq struct {
	NetworkName string `json:"network_name"` // 网络名称
	CoinSymbol  string `json:"coin_symbol"`  // 币种符号
	Address     string `json:"address"`      // 地址
	Amount      string `json:"amount"`       // 充值数量
	BusinessId  string `json:"business_id"`  // 业务ID
	BlockHigh   string `json:"block_high"`   // 区块高度
	Txid        string `json:"txid"`         // 区块链交易哈希
	Status      int    `json:"status"`       // 状态（0: 链上打包中，1：提币成功，2：提币失败）
}

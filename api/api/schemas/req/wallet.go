package req

//WalletCreateReq 钱包创建参数
type WalletCreateReq struct {
	Count uint `form:"count,default=1" binding:"min=1,max=100"` // 创建数量
}

//WalletPrivateKeyToAddressReq 私钥转钱包地址参数
type WalletPrivateKeyToAddressReq struct {
	PrivateKey string `form:"privateKey" binding:"required"` // 私钥
}

//WalletMnemonicToAddressAndPrivateKeyReq 助记词转钱包地址和私钥参数
type WalletMnemonicToAddressAndPrivateKeyReq struct {
	Mnemonic string `form:"mnemonic" binding:"required"` // 助记词
}

// WalletGetBalanceReq 获取余额参数
type WalletGetBalanceReq struct {
	Address string `form:"address" binding:"required"` // 钱包地址
}

// WalletTransferReq 转账参数
type WalletTransferReq struct {
	To     string  `form:"to" binding:"required"` // 收款方钱包地址
	Amount float64 `form:"amount" binding:"gt=0"` // 转账金额
}

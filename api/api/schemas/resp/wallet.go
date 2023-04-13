package resp

import "github.com/shopspring/decimal"

// WalletCreateResp 钱包创建返回信息
type WalletCreateResp struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	Mnemonic   string `json:"mnemonic"`
}

// WalletPrivateKeyToAddressResp 私钥解锁返回信息
type WalletPrivateKeyToAddressResp struct {
	Address string `json:"address"`
}

// WalletMnemonicToAddressAndPrivateKeyResp 助记词解锁返回信息
type WalletMnemonicToAddressAndPrivateKeyResp struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

// WalletGetBalanceResp 查询余额返回信息
type WalletGetBalanceResp struct {
	Balance decimal.Decimal `json:"balance"`
}

// WalletGetUsdtBalanceResp 查询USDT余额返回信息
type WalletGetUsdtBalanceResp struct {
	Balance decimal.Decimal `json:"balance"`
}

// WalletGetTokenBalanceResp 查询Token余额返回信息
type WalletGetTokenBalanceResp struct {
	Balance decimal.Decimal `json:"balance"`
}

// WalletTransferResp 主网币转账返回信息
type WalletTransferResp struct {
	Tx string `json:"tx"`
}

// WalletTransferResp USDT转账返回信息
type WalletUsdtTransferResp struct {
	Tx string `json:"tx"`
}

// WalletTokenferResp Token转账返回信息
type WalletTokenTransferResp struct {
	Tx string `json:"tx"`
}

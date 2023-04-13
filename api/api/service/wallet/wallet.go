package wallet

import (
	"server/api/logic"
	"server/api/schemas/resp"
)

var WalletService = walletService{}

// walletService 钱包服务实现类
type walletService struct{}

// Create 生成钱包
// Count: 生成数量，不填默认为1，最大100 (验证器中有限制)
func (ws walletService) Create(Count uint) (data []resp.WalletCreateResp) {
	data = logic.WalletLogic.BatchCreateWallet(Count)
	return
}

// PrivateKeyToAddress 私钥解锁钱包
func (ws walletService) PrivateKeyToAddress(privateKey string) resp.WalletPrivateKeyToAddressResp {
	address := logic.WalletLogic.PrivateKeyToAddress(privateKey)
	return resp.WalletPrivateKeyToAddressResp{
		Address: address,
	}
}

// MnemonicToAddressAndPrivateKey 助记词解锁钱包
func (ws walletService) MnemonicToAddressAndPrivateKey(mnemonic string) resp.WalletMnemonicToAddressAndPrivateKeyResp {
	address, privateKey := logic.WalletLogic.MnemonicToAddressAndPrivateKey(mnemonic)
	return resp.WalletMnemonicToAddressAndPrivateKeyResp{
		Address:    address,
		PrivateKey: privateKey,
	}
}

// GetBalance 查询余额
func (ws walletService) GetBalance(address string) resp.WalletGetBalanceResp {
	balance := logic.CommonLogic.GetBalance(address)
	return resp.WalletGetBalanceResp{
		Balance: balance,
	}
}

// GetUsdtBalance 查询USDT余额
func (ws walletService) GetUsdtBalance(address string) resp.WalletGetUsdtBalanceResp {
	balance := logic.CommonLogic.GetUsdtBalance(address)
	return resp.WalletGetUsdtBalanceResp{
		Balance: balance,
	}
}

// GetTokenBalance 查询Token余额
func (ws walletService) GetTokenBalance(address string) resp.WalletGetTokenBalanceResp {
	balance := logic.CommonLogic.GetTokenBalance(address)
	return resp.WalletGetTokenBalanceResp{
		Balance: balance,
	}
}

// Transfer 主网币转账
func (ws walletService) Transfer(to string, amount float64) resp.WalletTransferResp {
	tx := logic.WalletLogic.Transfer(to, amount)
	return resp.WalletTransferResp{
		Tx: tx,
	}
}

// UsdtTransfer USDT转账
func (ws walletService) UsdtTransfer(to string, amount float64) resp.WalletUsdtTransferResp {
	tx := logic.WalletLogic.UsdtTransfer(to, amount)
	return resp.WalletUsdtTransferResp{
		Tx: tx,
	}
}

// TokenTransfer Token转账
func (ws walletService) TokenTransfer(to string, amount float64) resp.WalletTokenTransferResp {
	tx := logic.WalletLogic.TokenTransfer(to, amount)
	return resp.WalletTokenTransferResp{
		Tx: tx,
	}
}

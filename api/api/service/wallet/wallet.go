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
func (ws walletService) GetBalance(address string) resp.WalletBalanceResp {
	balance := logic.CommonLogic.GetBalance(address)
	return resp.WalletBalanceResp{
		Balance: balance,
	}
}

// GetUsdtBalance 查询USDT余额
func (ws walletService) GetUsdtBalance(address string) resp.WalletBalanceResp {
	balance := logic.CommonLogic.GetUsdtBalance(address)
	return resp.WalletBalanceResp{
		Balance: balance,
	}
}

// GetTokenBalance 查询Token余额
func (ws walletService) GetTokenBalance(address string) resp.WalletBalanceResp {
	balance := logic.CommonLogic.GetTokenBalance(address)
	return resp.WalletBalanceResp{
		Balance: balance,
	}
}

// Transfer 主网币转账
func (ws walletService) Transfer(to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.Transfer(to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// UsdtTransfer USDT转账
func (ws walletService) UsdtTransfer(to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.UsdtTransfer(to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// TokenTransfer Token转账
func (ws walletService) TokenTransfer(to string, amount float64) resp.WalletTxResp {
	tx := logic.WalletLogic.TokenTransfer(to, amount)
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// Unlock TokenLock解锁合约代币
func (ws walletService) Unlock() resp.WalletTxResp {
	tx := logic.WalletLogic.Unlock()
	return resp.WalletTxResp{
		Tx: tx,
	}
}

// GetUnlockToken 获取TokenLock已解锁代币数量
func (ws walletService) GetUnlockToken() resp.WalletNumResp {
	num := logic.WalletLogic.GetUnlockToken()
	return resp.WalletNumResp{
		Num: num,
	}
}

// GetLastUnlockTimestamp 获取上次解锁时间戳
func (ws walletService) GetLastUnlockTimestamp() resp.WalletTimestampResp {
	timestamp := logic.WalletLogic.GetLastUnlockTimestamp()
	return resp.WalletTimestampResp{
		Timestamp: timestamp,
	}
}

// GetStartTimestamp 获取合约启动时间戳
func (ws walletService) GetStartTimestamp() resp.WalletTimestampResp {
	startTimestamp := logic.WalletLogic.GetStartTimestamp()
	return resp.WalletTimestampResp{
		Timestamp: startTimestamp,
	}
}

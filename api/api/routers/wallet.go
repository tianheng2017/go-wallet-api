package routers

import (
	"server/api/schemas/req"
	"server/api/service/wallet"
	"server/core"
	"server/core/response"
	"server/util"

	"github.com/gin-gonic/gin"
)

var WalletGroup = core.Group("/wallet")

func init() {
	group := WalletGroup
	group.AddGET("/create", create)
	group.AddPOST("/privateKeyUnlock", privateKeyUnlock)
	group.AddPOST("/mnemonicUnlock", mnemonicUnlock)
	group.AddPOST("/getBalance", getBalance)
	group.AddPOST("/getUsdtBalance", getUsdtBalance)
	group.AddPOST("/getTokenBalance", getTokenBalance)
	group.AddPOST("/transfer", transfer)
	group.AddPOST("/usdtTransfer", usdtTransfer)
	group.AddPOST("/tokenTransfer", tokenTransfer)
	group.AddPOST("/unlock", unlock)
	group.AddGET("/getUnlockToken", getUnlockToken)
	group.AddGET("/getLastUnlockTimestamp", getLastUnlockTimestamp)
	group.AddGET("/getStartTimestamp", getStartTimestamp)
}

// 生成钱包
func create(c *gin.Context) {
	var createReq req.WalletCreateReq
	util.VerifyUtil.VerifyQuery(c, &createReq)
	response.OkWithData(c, wallet.WalletService.Create(createReq.Count))
}

// 私钥解锁
func privateKeyUnlock(c *gin.Context) {
	var privateKeyReq req.WalletPrivateKeyUnlockReq
	util.VerifyUtil.VerifyBody(c, &privateKeyReq)
	response.OkWithData(c, wallet.WalletService.PrivateKeyUnlock(privateKeyReq.PrivateKey))
}

// 助记词解锁
func mnemonicUnlock(c *gin.Context) {
	var mnemonicReq req.WalletMnemonicUnlockReq
	util.VerifyUtil.VerifyBody(c, &mnemonicReq)
	response.OkWithData(c, wallet.WalletService.MnemonicUnlock(mnemonicReq.Mnemonic))
}

// 查询主网币余额
func getBalance(c *gin.Context) {
	var getBalanceReq req.WalletGetBalanceReq
	util.VerifyUtil.VerifyBody(c, &getBalanceReq)
	response.OkWithData(c, wallet.WalletService.GetBalance(getBalanceReq.Address))
}

// 查询Usdt余额
func getUsdtBalance(c *gin.Context) {
	var getUsdtBalanceReq req.WalletGetBalanceReq
	util.VerifyUtil.VerifyBody(c, &getUsdtBalanceReq)
	response.OkWithData(c, wallet.WalletService.GetUsdtBalance(getUsdtBalanceReq.Address))
}

// 查询Token余额
func getTokenBalance(c *gin.Context) {
	var getTokenBalanceReq req.WalletGetBalanceReq
	util.VerifyUtil.VerifyBody(c, &getTokenBalanceReq)
	response.OkWithData(c, wallet.WalletService.GetTokenBalance(getTokenBalanceReq.Address))
}

// 主网币转账(ETH/BNB等)
func transfer(c *gin.Context) {
	var transferReq req.WalletTransferReq
	util.VerifyUtil.VerifyBody(c, &transferReq)
	response.OkWithData(c, wallet.WalletService.Transfer(transferReq.To, transferReq.Amount))
}

// Usdt转账
func usdtTransfer(c *gin.Context) {
	var usdtTransferReq req.WalletTransferReq
	util.VerifyUtil.VerifyBody(c, &usdtTransferReq)
	response.OkWithData(c, wallet.WalletService.UsdtTransfer(usdtTransferReq.To, usdtTransferReq.Amount))
}

// Token转账
func tokenTransfer(c *gin.Context) {
	var tokenTransferReq req.WalletTransferReq
	util.VerifyUtil.VerifyBody(c, &tokenTransferReq)
	response.OkWithData(c, wallet.WalletService.TokenTransfer(tokenTransferReq.To, tokenTransferReq.Amount))
}

// TokenLock合约代币解锁
func unlock(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.Unlock())
}

// 获取TokenLock合约已解锁代币数量
func getUnlockToken(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetUnlockToken())
}

// 获取TokenLock上次解锁时间
func getLastUnlockTimestamp(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetLastUnlockTimestamp())
}

// 获取TokenLock合约启动时间
func getStartTimestamp(c *gin.Context) {
	response.OkWithData(c, wallet.WalletService.GetStartTimestamp())
}

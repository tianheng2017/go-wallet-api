package logic

import (
	"math/big"
	"server/config"
	"server/contracts"
	"server/util"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var NftLogic = nftLogic{}

// nftLogic NFT逻辑层
type nftLogic struct{}

// MintSeaDrop 铸造NFT
func (nl nftLogic) MintSeaDrop(minter string, quantity uint) (txHash string) {
	// 铸造人私钥解锁，获取钱包
	fromAddress := WalletLogic.PrivateKeyToAddress(config.Config.NftFromPrivateKey)
	// 校验验证接收人钱包格式
	CommonLogic.CheckAddress(minter, "接收人")
	// 获取NFT实例
	nftInstance, err := contracts.NewNftTransactor(
		common.HexToAddress(config.Config.NftContractAddress),
		CommonLogic.NewClient(),
	)
	util.CheckUtil.CheckApiErr(err, "获取NFT实例失败")
	// 创建签名选项
	auth, err := bind.NewKeyedTransactorWithChainID(
		CommonLogic.GetTruePrivateKey(config.Config.TokenFromPrivateKey),
		CommonLogic.GetChainID(),
	)
	util.CheckUtil.CheckApiErr(err, "创建签名选项失败")
	auth.Nonce = big.NewInt(int64(CommonLogic.GetNonce(fromAddress)))
	auth.GasLimit = uint64(600000)
	auth.GasPrice = CommonLogic.GetGasPrice()
	// 生成未签名事务
	tx, err := nftInstance.MintSeaDrop(
		auth,
		common.HexToAddress(minter),
		big.NewInt(int64(quantity)),
	)
	util.CheckUtil.CheckApiErr(err, "铸造NFT失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}

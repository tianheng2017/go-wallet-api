package logic

import (
	"errors"
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

// GetNftInstance 获取NFT实例
func (nl nftLogic) GetNftInstance() (nftInstance *contracts.Nft) {
	nftInstance, err := contracts.NewNft(common.HexToAddress(config.Config.NftContractAddress), CommonLogic.NewClient())
	util.CheckUtil.CheckApiErr(err, "获取NFT实例失败")
	return
}

// MintSeaDrop 铸造NFT
func (nl nftLogic) MintSeaDrop(minter string, quantity uint) (txHash string) {
	// 铸造人私钥解锁，获取钱包
	fromAddress := WalletLogic.PrivateKeyToAddress(config.Config.NftFromPrivateKey)
	// 校验验证接收人钱包格式
	CommonLogic.CheckAddress(minter, "接收人")
	// 获取NFT实例
	nftInstance := nl.GetNftInstance()
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

// GetApproved 获取NFT授权地址
func (nl nftLogic) GetApproved(tokenId uint) (address string) {
	// 获取NFT实例
	nftInstance := nl.GetNftInstance()
	// 获取授权地址
	approved, err := nftInstance.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	util.CheckUtil.CheckApiErr(err, "获取NFT授权地址失败")
	address = approved.Hex()
	return
}

// GetOwnerOf 获取NFT持有人
func (nl nftLogic) GetOwnerOf(tokenId uint) (address string) {
	// 获取NFT实例
	nftInstance := nl.GetNftInstance()
	// 获取授权地址
	approved, err := nftInstance.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	util.CheckUtil.CheckApiErr(err, "获取NFT持有人失败")
	address = approved.Hex()
	return
}

// TransferFrom 授权转出NFT
func (nl nftLogic) TransferFrom(from string, tokenId uint) (txHash string) {
	// 调用人私钥解锁，获取钱包
	toAddress := WalletLogic.PrivateKeyToAddress(config.Config.NftFromPrivateKey)
	// 校验持有人钱包格式
	CommonLogic.CheckAddress(from, "持有人")
	// 校验tokenId的授权地址
	approved := nl.GetApproved(tokenId)
	if approved != toAddress {
		util.CheckUtil.CheckApiErr(errors.New(""), "NFT未授权给铸造人")
	}
	// 获取NFT实例
	nftInstance := nl.GetNftInstance()
	// 创建签名选项
	auth, err := bind.NewKeyedTransactorWithChainID(
		CommonLogic.GetTruePrivateKey(config.Config.NftFromPrivateKey),
		CommonLogic.GetChainID(),
	)
	util.CheckUtil.CheckApiErr(err, "创建签名选项失败")
	auth.Nonce = big.NewInt(int64(CommonLogic.GetNonce(toAddress)))
	auth.GasLimit = uint64(600000)
	auth.GasPrice = CommonLogic.GetGasPrice()
	// 生成未签名事务
	tx, err := nftInstance.TransferFrom(
		auth,
		common.HexToAddress(from),
		common.HexToAddress(toAddress),
		big.NewInt(int64(tokenId)),
	)
	util.CheckUtil.CheckApiErr(err, "NFT转出失败")
	// 返回txHash
	txHash = tx.Hash().Hex()
	return
}

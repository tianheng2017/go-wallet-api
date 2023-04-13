package logic

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"server/config"
	"server/contracts"
	"server/util"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var CommonLogic = commonLogic{}

// commonLogic 公共逻辑层
type commonLogic struct{}

// NewClient 返回客户端实例
func (cm commonLogic) NewClient() *ethclient.Client {
	var currentNodeUrl string
	// 读取配置文件当前节点
	switch config.Config.CurrentNodeUrl {
	case 1:
		currentNodeUrl = config.Config.EthereumNodeUrl
	case 2:
		currentNodeUrl = config.Config.BscNodeUrl
	case 3:
		currentNodeUrl = config.Config.BscTestNodeUrl
	default:
		currentNodeUrl = config.Config.LocalNodeUrl
	}
	client, err := ethclient.Dial(currentNodeUrl)
	util.CheckUtil.CheckApiErr(err, "节点通信异常")
	return client
}

// CheckAddress 检查钱包地址
func (cm commonLogic) CheckAddress(address string, extra string) (result bool) {
	result = true
	reg := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if result = reg.MatchString(address); !result {
		util.CheckUtil.CheckApiErr(errors.New(""), extra+"钱包格式无效")
	}
	return
}

// IsZeroAddress 是否是零地址
func (cm commonLogic) IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}
	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// IsContractAddress 是否是合约地址
func (cm commonLogic) IsContractAddress(iaddress string) bool {
	cm.CheckAddress(iaddress, "")
	address := common.HexToAddress(iaddress)
	bytecode, err := CommonLogic.NewClient().CodeAt(context.Background(), address, nil)
	util.CheckUtil.CheckApiErr(err, "获取合约代码失败")
	isContract := len(bytecode) > 0
	return isContract
}

// GetTruePrivateKey 获取真正的私钥
func (cm commonLogic) GetTruePrivateKey(privKey string) (privateKey *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(privKey)
	util.CheckUtil.CheckApiErr(err, "私钥无效")
	return
}

// ToDecimal wei转小数
func (cm commonLogic) ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	return result
}

// ToWei 小数转换为wei
func (cm commonLogic) ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

// CalcGasCost 根据燃气上限和燃气价格计算燃气花费
func (cm commonLogic) CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV 从签名中提取R，S和V值
func (cm commonLogic) SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}
	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)
	return R, S, V
}

// ToCommonAddress 字符串转commonAddress
func (cm commonLogic) ToCommonAddress(address string) (commonAddress string) {
	commonAddress = common.HexToAddress(address).Hex()
	return
}

// GetBalance 获取钱包余额
func (cm commonLogic) GetBalance(address string) (result decimal.Decimal) {
	cm.CheckAddress(address, "")
	balance, err := cm.NewClient().BalanceAt(context.Background(), common.HexToAddress(address), nil)
	util.CheckUtil.CheckApiErr(err, "查询余额失败")
	result = cm.ToDecimal(balance, 18)
	return
}

// CheckBalance 校验钱包余额
func (cm commonLogic) CheckBalance(address string, amount float64) (result bool) {
	result = true
	balance := cm.GetBalance(address)
	value, _ := strconv.ParseFloat(balance.String(), 64)
	if value < amount {
		result = false
		tip := "钱包余额不足, 当前余额: " + fmt.Sprintf("%f", value)
		util.CheckUtil.CheckApiErr(errors.New(""), tip)
	}
	return
}

// GetUsdtBalance 获取USDT余额
func (cm commonLogic) GetUsdtBalance(address string) (result decimal.Decimal) {
	cm.CheckAddress(address, "")
	accountAddress := common.HexToAddress(address)
	contractAddress := common.HexToAddress(config.Config.UsdtContractAddress)
	usdtInstance, err := contracts.NewUsdtCaller(contractAddress, cm.NewClient())
	util.CheckUtil.CheckApiErr(err, "获取USDT实例失败")
	balance, err := usdtInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	util.CheckUtil.CheckApiErr(err, "查询USDT余额失败")
	result = cm.ToDecimal(balance, 18)
	return
}

// CheckUsdtBalance 校验钱包USDT余额
func (cm commonLogic) CheckUsdtBalance(address string, amount float64) (result bool) {
	result = true
	balance := cm.GetUsdtBalance(address)
	value, _ := strconv.ParseFloat(balance.String(), 64)
	if value < amount {
		result = false
		tip := "钱包USDT余额不足, 当前余额: " + fmt.Sprintf("%f", value)
		util.CheckUtil.CheckApiErr(errors.New(""), tip)
	}
	return
}

// GetTokenBalance 获取Token余额
func (cm commonLogic) GetTokenBalance(address string) (result decimal.Decimal) {
	cm.CheckAddress(address, "")
	accountAddress := common.HexToAddress(address)
	contractAddress := common.HexToAddress(config.Config.TokenContractAddress)
	tokenInstance, err := contracts.NewTokenCaller(contractAddress, cm.NewClient())
	util.CheckUtil.CheckApiErr(err, "获取Token实例失败")
	balance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, accountAddress)
	util.CheckUtil.CheckApiErr(err, "查询Token余额失败")
	result = cm.ToDecimal(balance, 18)
	return
}

// CheckTokenBalance 校验钱包Token余额
func (cm commonLogic) CheckTokenBalance(address string, amount float64) (result bool) {
	result = true
	balance := cm.GetTokenBalance(address)
	value, _ := strconv.ParseFloat(balance.String(), 64)
	if value < amount {
		result = false
		tip := "钱包代币余额不足, 当前余额: " + fmt.Sprintf("%f", value)
		util.CheckUtil.CheckApiErr(errors.New(""), tip)
	}
	return
}

// GetNonce 获取钱包nonce
func (cm commonLogic) GetNonce(fromAddress string) (nonce uint64) {
	cm.CheckAddress(fromAddress, "")
	address := common.HexToAddress(fromAddress)
	nonce, err := cm.NewClient().PendingNonceAt(context.Background(), address)
	util.CheckUtil.CheckApiErr(err, "获取nonce失败")
	return
}

// GetGasPrice 获取gasPrice
func (cm commonLogic) GetGasPrice() (gasPrice *big.Int) {
	gasPrice, err := cm.NewClient().SuggestGasPrice(context.Background())
	util.CheckUtil.CheckApiErr(err, "获取gasPrice失败")
	return
}

// GetChainID 获取chainID
func (cm commonLogic) GetChainID() (chainID *big.Int) {
	chainID, err := cm.NewClient().NetworkID(context.Background())
	util.CheckUtil.CheckApiErr(err, "获取chainID失败")
	return
}

// GetNewTransaction 生成未签名事务
func (cm commonLogic) GetNewTransaction(from string, to string, number float64, input []byte) (tx *types.Transaction) {
	nonce := cm.GetNonce(from)
	toAddress := common.HexToAddress(to)
	value := cm.ToWei(number, 18)
	gasLimit := uint64(600000)
	gasPrice := cm.GetGasPrice()
	tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, input)
	return
}

// GetSign 交易签名
func (cm commonLogic) GetSign(tx *types.Transaction, privKey string) (signedTx *types.Transaction) {
	result := cm.GetTruePrivateKey(privKey)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(cm.GetChainID()), result)
	util.CheckUtil.CheckApiErr(err, "签名失败")
	return
}

// Send 广播数据到链上
func (cm commonLogic) Send(signedTx *types.Transaction) (err error) {
	err = cm.NewClient().SendTransaction(context.Background(), signedTx)
	util.CheckUtil.CheckApiErr(err, "广播数据失败")
	return
}
package config

// config 环境配置
var Config = config{
	// -----------------------------------Gin配置-----------------------------------
	// gin运行模式 debug 或 release
	GinMode: "debug",

	// -----------------------------------API配置-----------------------------------
	// 服务端口，要改
	ServerPort: 8000,
	// API Key，要改
	ApiKey: "TzIyVyUC",

	// -----------------------------------节点配置-----------------------------------
	// 当前启用的节点，4选1，要改
	CurrentNodeUrl: 0,
	// 以下都可以从metamask网络里面复制，我这里从TokenPocket里面复制的
	// 0 - 本地测试节点
	LocalNodeUrl: "http://127.0.0.1:8545",
	// 1 - 以太坊节点
	EthereumNodeUrl: "https://web3.mytokenpocket.vip",
	// 2 - 币安链节点
	BscNodeUrl: "https://bsc-dataseed.binance.org",
	// 3 - 币安链测试网节点
	BscTestNodeUrl: "https://data-seed-prebsc-1-s3.binance.org:8545",

	// -----------------------------------NFT相关配置----------------------------------------
	// NFT铸造人私钥(去掉0x)，要改
	NftFromPrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
	// NFT合约地址，要改
	NftContractAddress: "0x1fDEdE3eBe5384776E4c2BA46a5c4b703daD83a6",

	// -----------------------------------主网币转账配置---------------------------------
	// 主网币发送人私钥(去掉0x)，要改
	PrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",

	// -----------------------------------USDT相关配置-----------------------------------
	// USDT发送人私钥(去掉0x)，要改
	UsdtFromPrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
	// USDT合约地址，要改
	UsdtContractAddress: "0xF40B98bab30Da9C583Bf4762dA927c51afEE2c9b",

	// -----------------------------------代币相关配置-----------------------------------
	// 这不是一个任意的代币，是babytoken，与contracts/token.abi对应
	// 代币发送人私钥(去掉0x)，要改
	TokenFromPrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
	// 代币合约地址，要改
	TokenContractAddress: "0x1eD10a20F94F56C103305f61FA1aAaf023F51eB6",

	// -----------------------------------TokenLock解锁配置------------------------------
	// TokenLock合约所有人私钥(去掉0x)，要改
	TokenLockPrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
	// TokenLock合约地址，要改
	TokenLockContractAddress: "0x825A62186a9386Ec3EAb3CC5C442ce51BE8Fb321",
}

type config struct {
	GinMode    string
	ServerPort int
	ApiKey     string

	CurrentNodeUrl  uint8
	EthereumNodeUrl string
	BscTestNodeUrl  string
	BscNodeUrl      string
	LocalNodeUrl    string

	NftFromPrivateKey  string
	NftContractAddress string

	PrivateKey string

	UsdtFromPrivateKey  string
	UsdtContractAddress string

	TokenFromPrivateKey  string
	TokenContractAddress string

	TokenLockPrivateKey      string
	TokenLockContractAddress string
}

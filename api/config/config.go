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
	// 以下都可以从metamask网络里面或TokenPocket里面复制
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
	NftContractAddress: "0x2bC3Ed169511Cc923338754DA100547fEdF838Cd",

	// -----------------------------------主网币转账配置---------------------------------
	// 主网币发送人私钥(去掉0x)，要改
	PrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",

	// -----------------------------------ERC20代币配置---------------------------------
	// 填写代币名称、合约地址、私钥、代币精度
	Erc20: map[string]Info{
		"AAA": {
			Contract:   "0xB577B45CAC1229a846B6250C8dBD964C8998be59",
			PrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
			Decimal:    18,
		},
		"BBB": {
			Contract:   "0x33578316FE7e9B48A0Ff0B4db811878B811F1843",
			PrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
			Decimal:    8,
		},
	},

	// -----------------------------------TokenLock解锁配置------------------------------
	// TokenLock合约所有人私钥(去掉0x)，要改
	TokenLockPrivateKey: "c9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed",
	// TokenLock合约地址，要改
	TokenLockContractAddress: "0x7BB86F63C687A1d5C963b53d93fB385de7661609",
	// TokenLock解锁的代币精度
	TokenLockDecimal: 18,
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

	Erc20 map[string]Info

	TokenLockPrivateKey      string
	TokenLockContractAddress string
	TokenLockDecimal         uint8
}

type Info struct {
	Contract   string
	PrivateKey string
	Decimal    uint8
}

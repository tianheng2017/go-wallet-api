require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    defaultNetwork: "ganache",
    networks: {
        // ganache本地测试链
        // hardhat和go全局测试用助记词：spoil test knock border hybrid drop remove hope banana face model seat
        ganache: {
            chainId: 1337,
            url: "http://127.0.0.1:8545",
            accounts: [
                '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
                '0x0575153404b945c98e6f7eecfade303890e00b4431aee66c0c2972354979afee',
            ],
            // 部署人私钥
            defaultAccount: '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
        },
        // bsc测试链
        bsctestnet: {
            url: "https://data-seed-prebsc-1-s3.binance.org:8545",
            accounts: [
                '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
                '0x0575153404b945c98e6f7eecfade303890e00b4431aee66c0c2972354979afee',
            ],
            defaultAccount: '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
        },
        // bsc主链
        bscmainnet: {
            url: "https://bsc-dataseed.binance.org/",
            accounts: [
                '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
                '0x0575153404b945c98e6f7eecfade303890e00b4431aee66c0c2972354979afee',
            ],
            defaultAccount: '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
        }
    },
    // 编译器配置，支持同时配置多个版本
    solidity: {
        compilers: [
            {
                version: "0.8.17",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
            {
                version: "0.5.16",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200
                    }
                }
            },
        ],
    },
};

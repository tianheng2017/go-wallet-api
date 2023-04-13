require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    defaultNetwork: "hardhat",
    networks: {
        // 本地测试链
        hardhat: {},
        // bsc测试链
        bsctestnet: {
            url: "https://data-seed-prebsc-1-s3.binance.org:8545",
            accounts: [
                '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
            ]
        },
        // bsc主链
        bscmainnet: {
            url: "https://bsc-dataseed.binance.org/",
            accounts: [
                '0xc9cab9af5f2ddf72c0a7ad70a313745939a00919dfc71f3b4badf71c88f5d9ed',
            ]
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

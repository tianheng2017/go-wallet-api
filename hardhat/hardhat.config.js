require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    defaultNetwork: "hardhat",
    networks: {
        // 本地测试链
        hardhat: {
            chainId: 1337,
        },
        // bsc测试链
        bsctestnet: {
            url: "https://data-seed-prebsc-1-s3.binance.org:8545",
            // 私钥组，代表accounts[0]....accounts[n]，默认accounts[0]为部署合约账号，其他测试用
            accounts: [
                // 对应钱包：0xe7c23A4c7a4771A79284797eD35598b32aB5E44a
                '0x4b5d78940ea5782ff0e2de98aee5b748bb7e4f57cc09f052a3f7e7bf7dd7ea24',
            ]
        },
        // bsc主链
        bscmainnet: {
            url: "https://bsc-dataseed.binance.org/",
            // 私钥组，代表accounts[0]....accounts[n]，默认accounts[0]为部署合约账号，其他测试用
            accounts: [
                // 对应钱包：0xe7c23A4c7a4771A79284797eD35598b32aB5E44a
                '0x4b5d78940ea5782ff0e2de98aee5b748bb7e4f57cc09f052a3f7e7bf7dd7ea24',
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
                        // 测试环境关闭优化，否者可能会导致堆栈跟踪的行数有些偏差
                        enabled: false,
                        runs: 200
                    }
                }
            },
            {
                version: "0.5.16",
                settings: {
                    optimizer: {
                        // 测试环境关闭优化，否者可能会导致堆栈跟踪的行数有些偏差
                        enabled: false,
                        runs: 200
                    }
                }
            },
        ],
    },
};

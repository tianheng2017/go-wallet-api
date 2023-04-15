const hre = require("hardhat");
const ethers = hre.ethers;

// 正式部署上链脚本demo
// 部署方法：npx hardhat run scripts/deploy.js --network bsc_testnet
// 要求：部署前先在hardhat.config.js中配置好accounts[0]，这是合约部署人
async function main() {
    const accounts = await ethers.getSigners();

    console.log("\n");

    //---------------------------------NFT---------------------------------
    // 部署NFT合约，设置NFT name、symbol、拥有铸造权的地址组
    const ERC721SeaDrop = await ethers.getContractFactory("ERC721SeaDrop");
    const nft = await ERC721SeaDrop.deploy(
        "MyNft", 
        "MNFT", 
        [accounts[0].address]   //[铸造人1,铸造人2]
    );
    console.log("1、NFT合约部署成功: ", nft.address);

    // 设置NFT最大供应量，不设置无法铸造，切记
    await nft.setMaxSupply(70000);
    console.log("1.1、设置NFT合约的最大供应量为70000", "\n", "\n");
    //---------------------------------NFT---------------------------------

    //--------------Token（请修改为BABYTOKEN合约，我这里用Token演示）--------
    // 部署Token.sol合约
    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy(
        "AAA",
        "AAA",
    );
    console.log("3、Token合约部署成功: ", token.address, "\n", "\n");
    //---------------------------------Token---------------------------------


    //---------------------------------TokenLock---------------------------------
    // 部署TokenLock合约
    const TokenLock = await ethers.getContractFactory("TokenLock");
    // 传入Token合约地址
    const tokenLock = await TokenLock.deploy(token.address);
    console.log("4、TokenLock合约部署成功: ", tokenLock.address);

    // 部署人转账给TokenLock合约1000000个Token用于锁仓
    console.log("4.1、部署人转账1000000个Token给TokenLock合约，模拟待解锁总额，方便测试解锁API");
    await token.transfer(
        tokenLock.address, 
        ethers.utils.parseEther("1000000"),
    );

    // 验证TokenLock合约的余额
    const tokenLockBalance = await token.balanceOf(tokenLock.address);
    console.log("4.2、验证TokenLock合约的余额: ", ethers.utils.formatEther(tokenLockBalance), "\n", "\n");
    //---------------------------------TokenLock---------------------------------
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
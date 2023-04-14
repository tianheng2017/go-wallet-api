const hre = require("hardhat");
const ethers = hre.ethers;

// 部署思路演示
async function main() {
    const accounts = await ethers.getSigners();

    // 部署USDT合约
    const BEP20USDT = await ethers.getContractFactory("BEP20USDT");
    const usdt = await BEP20USDT.deploy();
    console.log("USDT合约部署成功: ", usdt.address, "\n");

    // 部署NFT合约
    const ERC721SeaDrop = await ethers.getContractFactory("ERC721SeaDrop");
    // 部署时，设置NFT name、symbol、拥有铸造权的地址组
    const nft = await ERC721SeaDrop.deploy(
        "MyNft", 
        "MNFT", 
        [accounts[0].address]
    );
    console.log("NFT合约部署成功: ", nft.address, "\n");
    // 设置NFT最大供应量，不设置无法铸造，切记
    await nft.setMaxSupply(100);
    console.log("设置NFT合约得最大铸造数量为100", "\n");
    // 铸造1个NFT给accounts[1] (首次tokenId默认为1)
    await nft.mintSeaDrop(
        accounts[1].address, 
        1
    );
    console.log("铸造1个NFT给accounts[1]", "\n");
    // account[1]将NFT授权给account[0]，方便测试授权转账API
    await nft.approve(
        accounts[0].address, 
        1,
        { from: accounts[1].address }
    );
    console.log("account[1]将tokenId 1的NFT授权给account[0]，方便等下测试授权转账API", "\n");

    // 这个BABYTOKEN合约参数较复杂，没具体研究，未部署，我用Token.sol替代了
    // 部署Token.sol合约
    const Token = await ethers.getContractFactory("Token");
    const token = await Token.deploy(
        "AAA",
        "AAA",
    );
    console.log("Token合约部署成功: ", token.address, "\n");

    // 部署TokenLock合约
    const TokenLock = await ethers.getContractFactory("TokenLock");
    // 传入Token合约地址
    const tokenLock = await TokenLock.deploy(token.address);
    console.log("TokenLock合约部署成功: ", tokenLock.address, "\n");

    // 部署人转账给TokenLock合约1000000个Token用于锁仓
    console.log("转账1000000个Token给TokenLock合约，模拟待解锁总额", "\n");
    await token.transfer(
        tokenLock.address, ethers.utils.parseEther("1000000"),
    );

    // 验证TokenLock合约的余额
    const tokenLockBalance = await token.balanceOf(tokenLock.address);
    console.log("验证TokenLock合约的余额: ", ethers.utils.formatEther(tokenLockBalance), "\n");
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
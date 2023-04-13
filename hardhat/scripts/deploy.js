const hre = require("hardhat");
const ethers = hre.ethers;

async function main() {
    // 测试内容
    const accounts = await ethers.getSigners();
    console.log("account[0]地址: ", accounts[0].address);
    console.log("account[1]地址: ", accounts[1].address);
    
    console.log("account[0]开始部署USDT合约");
    const BEP20USDT = await ethers.getContractFactory("BEP20USDT");
    const usdt = await BEP20USDT.deploy();
    console.log("USDT合约部署成功: ", usdt.address);
    console.log("Account balance:", ethers.utils.formatEther((await accounts[0].getBalance())));

    console.log("account[0]转账20000USDT给account[1]");
    await usdt.transfer(accounts[1].address, ethers.utils.parseEther("20000"), {from: accounts[0].address});
    console.log("account[1]的余额: ", ethers.utils.formatEther((await accounts[1].getBalance())));

    console.log("account[1]授权 10000USDT 给account[2]");
    await usdt.approve(accounts[2].address, ethers.utils.parseEther("10000"), {from: accounts[1].address});
    const allowance = await usdt.allowance(accounts[2].address, accounts[1].address, { from: accounts[2].address });
    console.log("account[1]授权给account[2]的额度: ", ethers.utils.formatEther(allowance));

    console.log("account[2]通过授权转账将account[1]的10000USDT转给account[3]");
    await usdt.transferFrom(accounts[1].address, accounts[3].address, ethers.utils.parseEther("10000"), {from: accounts[2].address});
    console.log("account[3]的余额: ", ethers.utils.formatEther((await accounts[3].getBalance())));
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
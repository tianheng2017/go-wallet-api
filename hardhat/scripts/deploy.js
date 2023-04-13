async function main() {
    // 所有的合约都是hardhat.config.js中的accounts[0]账户部署的
    const BEP20USDT = await ethers.getContractFactory("BEP20USDT");
    const usdt = await BEP20USDT.deploy();
    console.log("USDT部署成功: ", usdt.address);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
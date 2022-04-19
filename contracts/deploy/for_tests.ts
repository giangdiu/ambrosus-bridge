import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";
import {ethers} from "hardhat";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (hre.network.name !== "hardhat") return;
  const {owner} = await hre.getNamedAccounts();


  const {address: mockAddr} = await hre.deployments.deploy("BridgeERC20Test", {
    contract: "BridgeERC20Test",
    from: owner,
    args: [
      "Mock", "Mock", 18,
      [ethers.constants.AddressZero], // bridgeAddresses
    ],
    log: true,
  });

  const {address: wrapperAddr} = await hre.deployments.deploy("sAMB", {
    contract: "sAMB",
    from: owner,
    args: ["sAMB", "sAMB"],
    log: true,
  });


  const commonArgs = {
    sideBridgeAddress: ethers.constants.AddressZero,
    adminAddress: ethers.constants.AddressZero,
    relayAddress: ethers.constants.AddressZero,
    wrappingTokenAddress: wrapperAddr,
    tokenThisAddresses: [],
    tokenSideAddresses: [],
    fee: 1000,
    feeRecipient: ethers.constants.AddressZero,
    timeframeSeconds: 14400,
    lockTime: 1000,
    minSafetyBlocks: 10,
  };

  await hre.deployments.deploy("CommonBridge", {
    from: owner,
    args: [commonArgs],
    log: true,
  });

  await hre.deployments.deploy("AmbBridgeTest", {
    from: owner,
    args: [commonArgs],
    log: true,
  });

  await hre.deployments.deploy("EthBridgeTest", {
    contract: "EthBridgeTest",
    from: owner,
    args: [commonArgs,
      [
        "0x4c9785451bb2CA3E91B350C06bcB5f974cA33F79",
        "0x90B2Ce3741188bCFCe25822113e93983ecacfcA0",
        "0xAccdb7a2268BC4Af0a1898e725138888ba1Ca6Fc"
      ],
      ethers.constants.AddressZero, // validatorSetAddress_
      ethers.constants.HashZero, // lastProcessedBlock
    ],
    log: true,
  });

};

export default func;

func.tags = ["for_tests"];

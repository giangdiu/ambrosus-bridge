import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DeployFunction} from "hardhat-deploy/types";
import {ethers} from "ethers";
import {addNewTokensToBridge, getAmbValidators, options, parseNet, readConfig_} from "./utils/utils";

const BRIDGE_NAME = "ETH_EthBridge";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const isMainNet = parseNet(hre.network).stage === 'main'

  let configFile = readConfig_(hre.network);
  const tokenPairs = configFile.getTokenPairs("eth", "amb")

  const ambNet = hre.companionNetworks['amb']
  const ambBridge = await ambNet.deployments.get('ETH_AmbBridge');

  const optionsWithOnUpgrade: any = await options(hre, BRIDGE_NAME, tokenPairs,
    {
      sideBridgeAddress: ambBridge.address,
      wrappingTokenAddress: configFile.tokens.WETH.addresses.eth,
      timeframeSeconds: isMainNet ? 60 * 60 * 4 : 60,
      lockTime: isMainNet ? 60 * 10 : 60,
      minSafetyBlocks: isMainNet ? 10 : 2,
    },
    [
      ...(await getAmbValidators(ambNet, isMainNet)),
      isMainNet ? 10 : 2, // minSafetyBlocksValidators
    ]
  )

  // Uncomment when updateLastProcessedBlock is needed
  // optionsWithOnUpgrade.proxy.execute.onUpgrade = {
  //   methodName: "updateLastProcessedBlock",
  //   args: ["0x1111111111111111111111111111111111111111111111111111111111111111"]
  // };

  if (isMainNet) {
    console.log("To update prod contract remove this if statement :)");

  }  else {

    const deployResult = await hre.deployments.deploy(BRIDGE_NAME, {
      contract: BRIDGE_NAME,
      ...optionsWithOnUpgrade
    });


    configFile.bridges.eth.side = deployResult.address;
    configFile.save()

    if (deployResult.newlyDeployed) {
      console.log('Call this cmd second time to update tokens')
      return;
    }
  }

  // add new tokens

  // DISABLE WETH TOKEN ( will be along with enabling USDC coin for gas economy :) )
  // todo remove this after call
  if (parseNet(hre.network).stage === "main") {
    tokenPairs["0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"] = "0x0000000000000000000000000000000000000000"
  }
  console.log(tokenPairs);
  // END

  await addNewTokensToBridge(tokenPairs, hre, BRIDGE_NAME);
};


export default func;
func.tags = ["bridges_eth"];
func.skip = async (hre: HardhatRuntimeEnvironment) => !hre.network.tags["eth"]; // only eth

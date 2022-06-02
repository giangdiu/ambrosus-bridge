import "@nomiclabs/hardhat-waffle";
import "hardhat-deploy";
import "hardhat-gas-reporter";
import "solidity-coverage";
import "hardhat-abi-exporter";
import * as dotenv from "dotenv";
import {HardhatUserConfig, TaskArguments} from "hardhat/types";
import {ethers} from "ethers";
import {task} from "hardhat/config";

dotenv.config();
// todo add other roles
const PK = [
  process.env.PRIVATEKEY || ethers.constants.HashZero,
  process.env.SECONDPRIVATEKEY || ethers.constants.HashZero,
];

const config: HardhatUserConfig = {

  networks: {
    hardhat: {
      blockGasLimit: 40000000, // amb value
      hardfork: "byzantium",
      companionNetworks: {amb: 'hardhat'},
      initialDate: "13 May 2022 18:10:36 GMT",

      // for debugging: run local node with state from some network
      forking: {
        enabled: false,
        url: "https://network.ambrosus-dev.io",
        // url: "https://eth-rinkeby.alchemyapi.io/v2/e1F5R9XuWDU2-zCtzaMDg4Ybb5SuoEDA"
        blockNumber: 0xb00ba,
      }
    },

    "dev/eth": {
      url: "https://nunki.htznr.fault.dev/rpc",
      accounts: PK,
      tags: ["eth", "devnet"],
      companionNetworks: {amb: 'dev/amb'},
      gasPrice: 9000000000
    },
    "test/eth": {
      url: "https://ropsten.infura.io/v3/" + process.env.INFURA_KEY,
      accounts: PK,
      tags: ["eth", "testnet"],
      companionNetworks: {amb: 'test/amb'},
    },
    "main/eth": {
      url: "https://mainnet.infura.io/v3/" + process.env.INFURA_KEY,
      accounts: PK,
      tags: ["eth", "mainnet"],
      companionNetworks: {amb: 'main/amb'},
    },
    "integr/eth": {
      url: "http://127.0.0.1:8502",
      accounts: ["0x51d098d8aee092622149d8f3a79cc7b1ce36ff97fadaa2fbd623c65badeefadc", "e7420b6492b8c876d23cd8a1156e35af4bc5dc5703fb4b79b376cb268a718e2e"],
      tags: ["eth", "integr"],
      companionNetworks: {amb: 'integr/amb'},
    },

    "dev/amb": {
      url: "https://network.ambrosus-dev.io",
      accounts: PK,
      tags: ["amb", "devnet"],
      hardfork: "byzantium",
    },
    "test/amb": {
      url: "https://network.ambrosus-test.io",
      accounts: PK,
      tags: ["amb", "testnet"],
      hardfork: "byzantium",
    },
    "main/amb": {
      url: "https://network.ambrosus.io",
      accounts: PK,
      tags: ["amb", "mainnet"],
      hardfork: "byzantium",
    },
    "integr/amb": {
      url: "http://127.0.0.1:8545",
      accounts: ["0x80f702eb861f36fe8fbbe1a7ccceb04ef7ddef714604010501a5f67c8065d446", "0x5b18f0adcca221f65373b20158f95313ecd51bde42b96a4c16f5eb851576bc06"],
      tags: ["amb", "integr"],
      hardfork: "byzantium",
    },
  },

  namedAccounts: {
    owner: 0,
    proxyAdmin: 1,
    admin: 2,
    relay: 3,
    bridge: 4,
    user: 5,
  },

  verify: {
    etherscan: {
      apiKey: "DY4Z86MQ2D9E24C6HB98PTA79EKJ5TQIFX",
    },
  },

  solidity: {
    compilers: [
      {
        version: "0.8.6",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,  // todo bigger
          },
          // Note: for amb deploy
          evmVersion: "byzantium"
        },
      }, {
        version: "0.4.22",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
          // Note: for amb deploy
          evmVersion: "byzantium"
        },
      },
    ],
  },

  abiExporter: {
    runOnCompile: true,
    clear: true,
    flat: true,
    only: [
      "Bridge$",
      ":ValidatorSet$",
      "sAMB",
      "IWrapper",
      ":ERC20$",
    ]
  }
};

task("confirmUpgrade", "Confirms last transaction", async (args: TaskArguments, hre) => {
  const {proxyAdmin} = await hre.getNamedAccounts();
  const proxyAdminS = await hre.ethers.getSigner(proxyAdmin);

  const contract = await hre.ethers.getContract(args.bridgename);
  const Factory = await hre.ethers.getContractFactory("MultiSigWallet");
  const proxy = await Factory.attach(contract.address);

  let transactionCount = await proxy.getTransactionCount(true, true);
  if (transactionCount === 0) {
    console.log("There is no transactions to confirm");
  } else {
    await (await proxy.connect(proxyAdminS).confirmTransaction(transactionCount.toNumber() - 1)).wait();
  }
}).addParam("bridgename");

export default config;

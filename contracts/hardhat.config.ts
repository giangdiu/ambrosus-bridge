import "@nomiclabs/hardhat-waffle";
import "hardhat-deploy";
import "hardhat-gas-reporter";
import "solidity-coverage";
import "hardhat-abi-exporter";


import * as dotenv from "dotenv";

dotenv.config();
const PK = [process.env.PRIVATEKEY || "00000000"];

module.exports = {
    networks: {
        hardhat: {
            // hardfork: "byzantium"

        },
        rinkeby: {
            url: "https://rinkeby.infura.io/v3/" + process.env.INFURA_KEY,
            accounts: PK,
        },
        mainnet: {
            url: "https://mainnet.infura.io/v3/" + process.env.INFURA_KEY,
            accounts: PK,
        },
        amb: {
            url: "https://network.ambrosus-dev.io",
            accounts: PK,
            hardfork: "byzantium"
        },
        oelocal: {
            url: "http://127.0.0.1:8545",
            accounts: PK,
        }
    },
    namedAccounts: {
        owner: 0,
        user: 1,
    },
    etherscan: {
        apiKey: "DY4Z86MQ2D9E24C6HB98PTA79EKJ5TQIFX",
    },
    solidity: {
        compilers: [
            {
                version: "0.8.6",
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200,
                    },
                    evmVersion: "byzantium"
                },
            },
        ],
    },
    abiExporter: {
        clear: true,
        flat: true,
        only: [
          "AmbBridge",
          "EthBridge",
          "ValidatorSet$",
        ]
    }
};

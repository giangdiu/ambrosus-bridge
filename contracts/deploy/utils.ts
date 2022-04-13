import path from "path";
import fs from "fs";


interface Token {
  name: string;
  symbol: string;
  denomination: number;
  addresses: { [net: string]: string }
  primaryNet: string;
}

interface Config {
  tokens: { [symb: string]: Token };
  bridges: { [net: string]: { amb: string, side: string } };
}


export function networkName(network: any): string {
  for (const networkName of ['amb', 'eth'])
    if (network.tags[networkName])
      return networkName;
  throw "Network missing networkName tag";
}

export function networkType(network: any): string {
  for (const networkType of ['testnet', 'mainnet'])
    if (network.tags[networkType])
      return networkType;
  throw "Network missing networkType tag";
}


export function getTokensPair(thisNet: string, sideNet: string, network: any): any {
  return _getTokensPair(thisNet, sideNet, readConfig(configPath(network)));
}

function _getTokensPair(thisNet: string, sideNet: string, configFile: Config): any {
  const tokensPair: { [k: string]: string } = {};

  for (const tokenThis of Object.values(configFile.tokens)) {
    if (!tokenThis.addresses[thisNet] || !tokenThis.addresses[sideNet]) continue;
    tokensPair[tokenThis.addresses[thisNet]] = tokenThis.addresses[sideNet];
  }
  return [Object.keys(tokensPair), Object.values(tokensPair)];
}

// get all deployed bridges in `net` network;
// for amb it's array of amb addresses for each network pair (such "amb-eth" or "amb-bsc")
// for other networks is array of one address
export function bridgesInNet(net: string, configFile: Config): string[] {
  const bridges = (net == "amb") ?
    Object.values(configFile.bridges).map(i => i.amb) :
    [configFile.bridges[net].side];
  return bridges.filter(i => !!i);  // filter out empty strings
}

export function configPath(network: any): string {
  return path.resolve(__dirname, `../config-${networkType(network)}.json`);
}

export function writeConfig(path: string, config: Config) {
  fs.writeFileSync(path, JSON.stringify(config, null, 2));
}

export function readConfig(tokenPath: string): Config {
  return require(tokenPath);
}


// :(((
export function urlFromHHProvider(provider: any): string {
  while (provider && !provider.url) provider = provider._wrapped;
  return provider.url
}


# Download abigen:
# https://geth.ethereum.org/downloads/

python merge_abis.py ETH_AmbBridge.json ETH_AmbBridge.json ETH_AmbBridge_PoW.json ETH_EthBridge.json BSC_AmbBridge.json BSC_BscBridge.json | abigen --abi - --pkg=bindings --type=bridge --out=../internal/bindings/bridge.go
abigen --abi=../../contracts/abi/ValidatorSet.json --pkg=bindings --type=vs --out=../internal/bindings/validatorSet.go
abigen --abi=../../contracts/abi/ERC20.json --pkg=bindings --type=token --out=../internal/bindings/erc20.go

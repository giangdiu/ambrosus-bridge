// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import "../checks/CheckPoSA.sol";


contract CheckPoSATest is CheckPoSA {
    constructor(
        bytes1 chainId_
    ) {
        chainId = chainId_;
    }

    function checkPoSATest(PoSAProof calldata posaProof, uint minSafetyBlocks, address sideBridgeAddress,
        address[] memory _initialValidators, uint _initialEpoch, bytes1 _chainId) public {
        // TODO: we can't use the `__CheckPoSA_init` twice, but copy-paste is also not good
        chainId = _chainId;
        currentEpoch = _initialEpoch;
        currentValidatorSetSize = _initialValidators.length;

        for (uint i = 0; i < _initialValidators.length; i++) {
            allValidators[currentEpoch][_initialValidators[i]] = true;
        }

        checkPoSA_(posaProof, minSafetyBlocks, sideBridgeAddress);
    }

    function blockHashTest(BlockPoSA calldata block_) public view returns (bytes32, bytes32) {
        return calcBlockHash(block_);
    }

    function blockHashTestPaid(BlockPoSA calldata block_) public returns (bytes32, bytes32) {
        return calcBlockHash(block_);
    }

}

// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import "./CommonStructs.sol";


contract CheckReceiptsProof {
    // check readme for focs
    function CalcReceiptsHash(
        bytes[] memory proof, address eventContractAddress, bytes memory topic, bytes memory data
    ) public view returns (bytes32) {
        bytes32 el = keccak256(abi.encodePacked(proof[0], eventContractAddress, proof[1], topic, proof[2], data, proof[3]));
        bytes memory s;

        for (uint i = 4; i < proof.length; i += 2) {
            s = abi.encodePacked(proof[i], el, proof[i + 1]);
            el = (s.length > 32) ? keccak256(s) : bytes32(s);
        }

        return el;
    }

    function Test(bytes32 receiptsRoot, bytes[] memory proof,
        address eventContractAddress, bytes memory topic, bytes memory data
    ) public {
        require(CalcReceiptsHash(proof, eventContractAddress, topic,  data) == receiptsRoot, "Failed to verify receipts proof");
    }

}
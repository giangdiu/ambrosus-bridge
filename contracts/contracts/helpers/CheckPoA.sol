// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import "./CommonStructs.sol";


contract CheckPoA {

    struct BlockPoA {
        bytes p0_seal;
        bytes p0_bare;

        bytes p1;
        bytes32 prevHashOrReceiptRoot;  // receipt for main block, prevHash for safety blocks
        bytes p2;
        bytes timestamp;
        bytes p3;

        bytes s1;
        bytes signature;
        bytes s2;
    }


    function CheckPoA(BlockPoA[] memory blocks, CommonStructs.Transfer[] memory events, bytes[] memory proof) public {
        bytes32 hash = calcReceiptsRoot(proof, abi.encode(events));

        for (uint i = 0; i < blocks.length; i++) {
            require(blocks[i].prevHashOrReceiptRoot == hash, "prevHash or receiptRoot wrong");

            bytes memory rlp = abi.encodePacked(blocks[i].p1, blocks[i].prevHashOrReceiptRoot, blocks[i].p2, blocks[i].timestamp, blocks[i].p3);

            // hash without seal for signature check
            bytes32 bare_hash = keccak256(abi.encodePacked(blocks[i].p0_bare, rlp));
            address validator = GetValidator(bytesToUint(blocks[i].timestamp));
            CheckSignature(validator, bare_hash, blocks[i].signature);

            // hash with seal, for prev_hash check
            hash = keccak256(abi.encodePacked(blocks[i].p0_seal, rlp, blocks[i].s1, blocks[i].signature, blocks[i].s2));
        }
    }


    function GetValidator(uint timestamp) internal view returns (address) {
        // todo
        return address(this);
    }



    function CheckReceiptsProof(bytes[] memory proof, bytes memory eventToSearch, bytes32 receiptsRoot) public {
        require(calcReceiptsRoot(proof, eventToSearch) == receiptsRoot, "Failed to verify receipts proof");
    }

    function calcReceiptsRoot(bytes[] memory proof, bytes memory eventToSearch) public view returns (bytes32){
        bytes32 el = keccak256(abi.encodePacked(proof[0], eventToSearch, proof[1]));
        bytes memory s;

        for (uint i = 2; i < proof.length - 1; i += 2) {
            s = abi.encodePacked(proof[i], el, proof[i + 1]);
            el = (s.length > 32) ? keccak256(s) : bytes32(s);
        }

        return el;
    }


    function CheckSignature(address signer, bytes32 message, bytes memory signature) internal view {
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }
        require(
            ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", message)), v, r, s) == signer,
            "Failed to verify sign");
    }


    function bytesToUint(bytes memory b) public view returns (uint){
        return uint(bytes32(b)) >> (256 - b.length * 8);
    }
}
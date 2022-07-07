// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import "./CheckReceiptsProof.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./SignatureCheck.sol";


contract CheckAura is Initializable {
    bytes1 constant PARENT_HASH_PREFIX = 0xA0;
    bytes1 constant STEP_PREFIX = 0x84;
    bytes2 constant SIGNATURE_PREFIX = 0xB841;

    address[] public validatorSet;
    address validatorSetAddress;
    bytes32 public lastProcessedBlock; // actually latest processed (finalized) *vs change event* block
    uint public minSafetyBlocksValidators;


    struct BlockAura {
        bytes3 p0Seal;
        bytes3 p0Bare;

        bytes32 parentHash;
        bytes p2;
        bytes32 receiptHash;
        bytes p3;

        bytes4 step;
        bytes signature;  // todo maybe pass s r v values?

        uint64 finalizedVs;
    }


    struct ValidatorSetChange {
        address deltaAddress;
        int64 deltaIndex; // < 0 ? remove : add
    }

    struct ValidatorSetProof {
        bytes[] receiptProof;
        ValidatorSetChange[] changes;
        uint eventBlock;
    }

    struct AuraProof {
        BlockAura[] blocks;
        CommonStructs.TransferProof transfer;
        ValidatorSetProof[] vsChanges;
        uint64 transferEventBlock;
    }


    function __CheckAura_init(
        address[] calldata initialValidators_,
        address validatorSetAddress_,
        bytes32 lastProcessedBlock_,
        uint minSafetyBlocksValidators_
    ) internal initializer {
        require(initialValidators_.length > 0, "Length of _initialValidators must be bigger than 0");

        validatorSet = initialValidators_;
        validatorSetAddress = validatorSetAddress_;
        lastProcessedBlock = lastProcessedBlock_;
        minSafetyBlocksValidators = minSafetyBlocksValidators_;
    }



    /*
     AuraProof.blocks contains:
      - blocks for validate validatorSet change event, in that order:
        - block with `InitiateChange` events (contains list of all validators)
        - some blocks, that need for validation (the amount depends on the length of the current validator set)
        - block, when validators finalize; have `finalizedVs` != 0
        * repeated for each vs change event; all events must go in order, without omissions *

      - block with transfer event;
      - safety blocks for transfer event

      AuraProof.vsChanges contains changes in validator set and receiptProof for validation.
      block.finalizedVs-1 is index in AuraProof.vsChanges array

      Function will check all blocks, processing vs change events if needed.
      Each block parentHash must be equal to the seal hash of the previous block, except for gaps between vsChange events
      If there are no errors, the transfer is considered valid
    */
    function checkAura_(AuraProof calldata auraProof, uint minSafetyBlocks, address sideBridgeAddress) internal {

        bytes32 parentHash;
        bytes32 receiptHash;

        // auraProof can be without transfer event when we have to many vsChanges and transfer doesn't fit into proof
        if (auraProof.transferEventBlock != 0) {
            receiptHash = calcTransferReceiptsHash(auraProof.transfer, sideBridgeAddress);
            require(auraProof.blocks[auraProof.transferEventBlock].receiptHash == receiptHash, "Transfer event validation failed");
            require(auraProof.blocks.length - auraProof.transferEventBlock >= minSafetyBlocks, "Not enough safety blocks");
        }

        for (uint i = 0; i < auraProof.blocks.length; i++) {
            BlockAura calldata block_ = auraProof.blocks[i];

            // 0 means no events should be finalized, so indexes are shifted by 1
            if (block_.finalizedVs != 0) {
                // vs changes in that block
                ValidatorSetProof calldata vsProof = auraProof.vsChanges[block_.finalizedVs - 1];

                // apply vs changes
                for (uint k = 0; k < vsProof.changes.length; k++)
                    applyVsChange(vsProof.changes[k]);

                // check proof
                receiptHash = calcValidatorSetReceiptHash(vsProof.receiptProof, validatorSetAddress, validatorSet);
                require(auraProof.blocks[vsProof.eventBlock].receiptHash == receiptHash, "Wrong VS receipt hash");
                require(i - vsProof.eventBlock >= minSafetyBlocksValidators, "Few safety blocks validators");

                // there is gap BEFORE finalizing block, so disable parentHash check for it
                // but only if it's not the safety blocks for transfer event
                if (i < auraProof.transferEventBlock)
                    parentHash = bytes32(0);
            }

            // check that block.parentHash == parentHash (hash of prev block)
            // don't check parentHash for first block and for blocks before and after finalizing vs
            if (parentHash != bytes32(0))
                require(block_.parentHash == parentHash, "Wrong parent hash");

            // calc block hash for this block
            parentHash = checkBlock(block_);

            // and again, if this block is finalizing block
            if (block_.finalizedVs != 0) {
                // save it hash to `lastProcessedBlock`
                // todo saving to temp memory variable may reduce gas cost
                lastProcessedBlock = parentHash;

                // there is gap AFTER finalizing block, so disable parentHash check for it
                // but only if it's not the safety blocks for transfer event
                if (i < auraProof.transferEventBlock)
                    parentHash = bytes32(0);
            }

        }
    }

    function getValidatorSet() public view returns (address[] memory) {
        return validatorSet;
    }

    function applyVsChange(ValidatorSetChange calldata vsEvent) internal {
        // todo invert indexes sign meaning
        if (vsEvent.deltaIndex < 0) { // delete validator
            uint index = uint(int(vsEvent.deltaIndex * (- 1) - 1));
            validatorSet[index] = validatorSet[validatorSet.length - 1];
            validatorSet.pop();
        } else { // add validator
            validatorSet.push(vsEvent.deltaAddress);
        }
    }

    function checkBlock(BlockAura calldata block_) internal view returns (bytes32) {
        (bytes32 bareHash, bytes32 sealHash) = calcBlockHash(block_);

        address validator = validatorSet[bytesToUint(block_.step) % validatorSet.length];
        require(ecdsaRecover(bareHash, block_.signature) == validator, "Failed to verify sign");

        return sealHash;
    }

    function calcBlockHash(BlockAura calldata block_) internal pure returns (bytes32, bytes32) {
        bytes memory commonRlp = abi.encodePacked(PARENT_HASH_PREFIX, block_.parentHash, block_.p2, block_.receiptHash, block_.p3);
        return (
        // hash without seal (bare), for signature check
        keccak256(abi.encodePacked(block_.p0Bare, commonRlp)),
        // hash with seal, for prevHash check
        keccak256(abi.encodePacked(block_.p0Seal, commonRlp, STEP_PREFIX, block_.step, SIGNATURE_PREFIX, block_.signature))
        );
    }


    function calcValidatorSetReceiptHash(bytes[] calldata receiptProof, address validatorSetAddress_, address[] storage vSet) private pure returns (bytes32) {
        bytes32 el = keccak256(abi.encodePacked(
                receiptProof[0],
                validatorSetAddress_,
                receiptProof[1],
                abi.encode(vSet),
                receiptProof[2]
            ));
        return calcReceiptsHash(receiptProof, el, 3);
    }

    function bytesToUint(bytes4 b) internal pure returns (uint){
        return uint(uint32(b));
    }

    uint256[15] private ___gap;
}

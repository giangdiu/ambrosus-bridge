// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import "./helpers/CommonBridge.sol";
import "./helpers/CheckPoW.sol";
import "./helpers/CommonStructs.sol";

contract AmbBridge is CommonBridge, CheckPoW {
    constructor(
        address _sideBridgeAddress, address relayAddress,
        address[] memory tokenThisAddresses, address[] memory tokenSideAddresses,
        uint fee_, uint timeframeSeconds_, uint lockTime_, uint minSafetyBlocks_)
    CommonBridge(_sideBridgeAddress, relayAddress,
                 tokenThisAddresses, tokenSideAddresses,
                 fee_, timeframeSeconds_, lockTime_, minSafetyBlocks_) { emitTestEvent(address(this), msg.sender, 10); }

    function submitTransfer(
        uint event_id,
        BlockPoW[] memory blocks,
        CommonStructs.Transfer[] memory events,
        bytes[] memory proof,
        uint passedBlocks) public onlyRole(RELAY_ROLE) {

        require(event_id == inputEventId + 1);
        inputEventId++;

        require(passedBlocks > minSafetyBlocks, "passedBlocks must be larger than minSafetyBlocks");

        CheckPoW_(blocks, events, proof);

        lockTransfers(events, event_id);
    }
}

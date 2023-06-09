// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;
//https://github.com/hackbg/chainlink-makerdao-automation/tree/main/vendor
interface IUpkeepRefunder {
    function refundUpkeep() external;

    function shouldRefundUpkeep() external view returns (bool);

    function setUpkeepId(uint256 _upkeepId) external;
}
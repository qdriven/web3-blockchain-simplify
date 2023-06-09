
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;
import "hardhat/console.sol";
import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";


/**
 * If you are reading data feeds on L2 networks, you must
 * check the latest answer from the L2 Sequencer Uptime
 * Feed to ensure that the data is accurate in the event
 * of an L2 sequencer outage. See the
 * https://docs.chain.link/data-feeds/l2-sequencer-feeds
 * page for details.
 */

contract DataFeedConsumerV1 {
    //0xF6b9663D91157198a26F37009e85E8666992bEAc
    AggregatorV3Interface internal dataFeed;

    /**
     * Network: Sepolia
     * Aggregator: BTC/USD
     * Address: 0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
     */
    constructor() {
        dataFeed = AggregatorV3Interface(
            0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
        );
    }

    /**
     * Returns the latest answer.
     */
    function getLatestData() public view returns (int) {
        // prettier-ignore
        (
            uint80 roundID,
            int answer,
            uint startedAt,
            uint timeStamp,
            uint80 answeredInRound
        ) = dataFeed.latestRoundData();
        // console.log(roundID,answer,startedAt,timeStamp,answeredInRound);
        return answer;
    } 
}

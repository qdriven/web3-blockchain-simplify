// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;


contract PublishProxy{
    struct MsgWithStatus {
        string content;
        bool stauts;
    }
    event PublishEvent(address sender, string content);
    mapping (string => mapping (address => string[])) topicMapping;
    constructor(){
        
    }
    function publish(string memory topic, 
                    string memory message) public {
        topicMapping[topic][msg.sender].push(message);
        emit PublishEvent(msg.sender, message);
    }

    function GetMyTopics(string memory topic) public view returns (string[] memory) {
        return topicMapping[topic][msg.sender];
    }

    function GetAddressTopics(string memory topic,address otherAddress) public view returns (string[] memory) {
        return topicMapping[topic][otherAddress];
    }
}
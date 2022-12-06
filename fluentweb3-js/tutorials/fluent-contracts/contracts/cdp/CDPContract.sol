pragma solidity >=0.8.0 <0.9.0;
//SPDX-License-Identifier: MIT
import "./EmissionDataAuth.sol";


contract CDPContract is EmissionDataAuth {
  struct EmissionData {
    address owner;
    // uint256 timestamp;
    string data; //string data structure
    bytes referenceFile;
  }
  mapping(address => EmissionData[]) private emissionStore;

  event AddEmissionDataEvent(address owner, string data, bytes fileData);

  function addEmissionData(string memory emissonData, bytes memory referenceFile) public auth {
    emissionStore[msg.sender].push(EmissionData(msg.sender, emissonData, referenceFile));
    emit AddEmissionDataEvent(msg.sender, emissonData, referenceFile);
  }

  function getEmisionData() public view auth returns (EmissionData[] memory) {
    return emissionStore[msg.sender];
  }
}
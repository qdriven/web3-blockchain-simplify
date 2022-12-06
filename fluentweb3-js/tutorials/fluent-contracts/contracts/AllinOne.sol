pragma solidity >=0.8.0 <0.9.0;

//SPDX-License-Identifier: MIT
/**
 * Only For Manual Deploy Contract in Quorum Explorer
 */
contract EmissionAppContract {
  struct EmissionData {
    address owner;
    // uint256 timestamp;
    string data;
    // bytes32[] referenceFile;
  }
  mapping(address => EmissionData[]) private emissionStore;

  event AddEmissionDataEvent(address owner, string data);

  address public _owner;
  mapping(address => bool) private _auth;
  event AuthChangeEvent(address user, bool targetVal);

  constructor() {
    _owner = msg.sender;
  }

  modifier onlyOwner() {
    require(msg.sender == _owner, "Not Admin");
    _;
  }

  modifier auth() {
    require(_auth[msg.sender] == true, "Not authenticated");
    _;
  }

  function addToWhiteList(address user) public onlyOwner {
    _auth[user] = true;
    emit AuthChangeEvent(user, true);
  }

  function removeFromWhiteList(address user) public onlyOwner {
    _auth[user] = false;
    emit AuthChangeEvent(user, false);
  }

  function addEmissionData(string memory emissonData) public auth {
    emissionStore[msg.sender].push(EmissionData( msg.sender, emissonData));
    emit AddEmissionDataEvent(msg.sender, emissonData);
  }

  function getEmissionData() public view auth returns (EmissionData[] memory) {
    return emissionStore[msg.sender];
  }
}

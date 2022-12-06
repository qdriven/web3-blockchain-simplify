pragma solidity >=0.8.0 <0.9.0;

//SPDX-License-Identifier: MIT

contract EmissionDataAuth {
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
    require(msg.sender != _owner && _auth[msg.sender] == true, "Not authenticated");
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
}

// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;
import "hardhat/console.sol";

error HelloError();

contract Hello {
    string public world;

    constructor(string memory _world) {
        console.log("Deploying a Hello Contract:", _world);
        world = _world;
    }

    function get() public view returns (string memory) {
        return world;
    }

    function setWorld(string memory _world) public {
        console.log("Set a Hello Contract from %s to %s", world, _world);
        world = _world;
    }

    function _throwError() internal pure {
        revert HelloError();
    }

    function setWorldButFail(string memory _world) public {
        console.log("Set a Hello Contract from %s to %s", world, _world);
        world = _world;
        _throwError();
    }
}

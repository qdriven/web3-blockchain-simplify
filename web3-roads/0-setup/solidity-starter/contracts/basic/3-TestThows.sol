// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.7;

contract TestThrows {
    function testAssert() public pure {
        assert(1 == 2);
    }
    
    function testRequire() public pure {
        require(2 == 1,"not equal");
    }
    
    function testRevert() public pure {
        revert("revert it");
    }

}
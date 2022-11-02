// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
// // https://github.com/owner/repo/blob/branch/path/to/Contract.sol
// import "https://github.com/owner/repo/blob/branch/path/to/Contract.sol";

// // Example import ECDSA.sol from openzeppelin-contract repo, release-v4.5 branch
// // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/contracts/utils/cryptography/ECDSA.sol
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/contracts/utils/cryptography/ECDSA.sol";

// // import Foo.sol from current directory
// import "./Foo.sol";

struct Point {
    uint x;
    uint y;
}

error Unauthorized(address caller);

function add(uint x, uint y) pure returns (uint) {
    return x + y;
}

contract Foo {
    string public name = "Foo";
}

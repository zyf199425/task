// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

contract Counter {

    event CountChangeEvent(address indexed operator, uint256 newCount, string operation);

    uint256 private count;

    function increment() public {
        count = count + 1;
        emit CountChangeEvent(msg.sender, count, "increment");
    }

    function decrement() public {
        require(count > 0, "count must be > 0");
        count = count - 1;
        emit CountChangeEvent(msg.sender, count, "decrement");
    }

     function reset() public {
        count = 0;
        emit CountChangeEvent(msg.sender, count, "reset");
    }

    function getCount() public view returns (uint256){
        return count;
    }

}
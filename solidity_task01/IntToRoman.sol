// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目4
contract IntToRoman {
    uint256[] values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    bytes[] symbols = [bytes('M'), bytes('CM'), bytes('D'), bytes('CD'), bytes('C'), bytes('XC'), bytes('L'), bytes('XL'), bytes('X'), bytes('IX'), bytes('V'), bytes('IV'), bytes('I')];
    
    
    function intToRoman(uint256 num) public view returns (string memory) {
        bytes memory result = "";
        uint i = 0;
        while(num > 0) {
            while(values[i] > num) {
                i += 1;
            }
            num -= values[i];
            result = bytes.concat(result, symbols[i]);
        }
        return string(result);
    }
}
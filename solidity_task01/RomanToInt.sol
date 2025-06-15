// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目3
contract Roman {

    mapping(bytes1 => uint256) public m;

    constructor() {
        m[bytes1('I')] = 1;
        m[bytes1('V')] = 5;
        m[bytes1('X')] = 10;
        m[bytes1('L')] = 50;
        m[bytes1('C')] = 100;
        m[bytes1('D')] = 500;
        m[bytes1('M')] = 1000;
    }

    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory strBytes = bytes(s);
        uint256 result = 0;
        uint256 length = strBytes.length;
        uint256 prev = 0;
        for(uint256 i = 0;i < length;i++){
            uint256 currentNum = getNum(strBytes[i]);
            if( prev != 0 && currentNum > prev) {
                result -= prev;
                result += (currentNum - prev);
            }else{
                result += currentNum;
            }
            prev = currentNum;
        }
        return result;
    }

    function getNum(bytes1  b ) private view returns (uint256) {
        return m[b];
    }
}
// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

contract BinarySearch {
    
    function search(uint[] memory nums, uint target) public pure returns (int256) {
        uint left = 0;
        uint right = nums.length;
        uint mid = 0;
        while(left < right) {
            mid = (left + right ) / 2;
            if(nums[mid] > target){
                right = mid;
            }else if(nums[mid] < target){
                left = mid + 1;
            }else{
                return int256(mid);
            }
        }
        return -1;
    }}
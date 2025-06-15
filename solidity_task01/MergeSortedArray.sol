// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目5
contract MergeSortedArray {

    function mergeArray(uint[] memory nums1, uint[] memory nums2) public pure returns (uint[] memory) {
        uint[] memory result = new uint[](nums1.length + nums2.length);
        uint i1 = 0;
        uint i2 = 0;
        uint index = 0;
        while(i1 < nums1.length && i2 < nums2.length) {
            if(nums1[i1] < nums2[i2]) {
                result[index] = nums1[i1];
                i1++;
            }else{
                result[index] = nums2[i2];
                i2++;
            }
            index++;
        }
        while(i1 < nums1.length) {
            result[index] = nums1[i1];
            i1++;
            index++;
        }
        while(i2 < nums2.length) {
            result[index] = nums2[i2];
            i2++;
            index++;
        }
        return result;
    }
}
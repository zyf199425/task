// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目1
contract Voting {
    
    // 候选人的得票数
    mapping(address => uint256) public candidateVotes;
    
    // 候选人列表
    address[] public candidates;

    // 投票
    function vote(address _vandidate) public {
        require(_vandidate != address(0), "Invalid candidate address");
        candidates.push(_vandidate);
        candidateVotes[_vandidate]++;
    }

    // 获取候选人得票数
    function getVotes(address _vandidate) public view returns (uint256) {
        return candidateVotes[_vandidate];
    }

    // 重置投票结果
    function resetVotes() public {
        for(uint256 i = 0; i < candidates.length; i++) {
            candidateVotes[candidates[i]] = 0;
        }
    }
}
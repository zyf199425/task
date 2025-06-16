// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目3
// 创建一个名为 BeggingContract 的合约。
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。
contract BeggingContract {
    // 记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    // 合约所有者
    address payable public owner;
    // 总捐赠金额
    uint256 public totalDonations;
    // 开始时间
    uint256 public startTime;
    // 持续时间
    uint256 public duration;

    event Donation(address indexed donor, uint256 amount);

    constructor(uint256 _duration) {
        startTime = block.timestamp;
        duration = _duration;
        owner = payable(msg.sender);
    }

     // 仅所有者可调用的修饰器
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this");
        _;
    }
    // 捐赠金额
    function donate() public payable {
        // 校验时间
        require(block.timestamp < startTime, "Donation period have not started");
        require(startTime + duration < block.timestamp, "Donation period has ended");
        // 校验金额
        require(msg.value > 0, "Donation amount must be greater than 0");

        donations[msg.sender] += msg.value;
        totalDonations += msg.value;
        emit Donation(msg.sender, msg.value);
    }
    // 提取所有资金，只有所有者可以调用
    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        owner.transfer(balance);
    }
    // 查询某个地址的捐赠金额
    function getDonation(address addr) public view returns (uint256) {
        return  donations[addr];
    }
    // 查询捐赠总金额
    function getTotalDonations() public view returns (uint256) {
        return totalDonations;
    }
}
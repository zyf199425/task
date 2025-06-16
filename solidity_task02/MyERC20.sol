// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

// 题目一：实现一个ERC20代币合约
contract MyERC20 {
    // 代币基本信息
    uint256 private _totalSupply;
    string private _name;
    string private _symbol;
    uint8 public _decimals;
    // 合约所有者
    address public _owner;

    mapping(address account => uint256) private _balances;

    mapping(address account => mapping(address spender => uint256)) private _allowances;

    event Transfer(address indexed from, address indexed to, uint256 value);

    event Approve(address indexed owner, address indexed spender, uint256 value);

    // 仅所有者可调用的修饰器
    modifier onlyOwner() {
        require(msg.sender == _owner, "Only owner can call this");
        _;
    }

    constructor(string memory name_, string memory symbol_, uint8 decimals_, uint256 initialSupply) {
        _name = name_;
        _symbol = symbol_;
        _decimals = decimals_;
        _totalSupply = initialSupply;
        _owner = msg.sender;
        // 初始铸造给合约创建者
        _mint(msg.sender, initialSupply * (10 ** uint256(_decimals)));
    }
    function name() public view virtual returns (string memory) {
        return _name;
    }
    function symbol() public view virtual returns (string memory) {
        return _symbol;
    }
    function decimals() public view virtual returns (uint8) {
        return _decimals;
    }
    // 查询总供应量
    function totalSupply() public view virtual returns (uint256) {
        return _totalSupply;
    }
    // 查询余额
    function balanceOf(address account) external view returns (uint256){
        return _balances[account];
    }
    // 查询授权额度
    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    // 转账
    function transfer(address to, uint256 value) external returns (bool){
        address from = msg.sender;
        _transfer(from, to, value);
        return true;
    }

    function _transfer(address from, address to, uint256 value) internal {
        require(from != address(0), "ERC20: transfer to the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");

        uint256 fromBalance = _balances[from];
        require(fromBalance >= value, "ERC20: transfer amount exceeds balance");
        
        // 更新余额
        _balances[from] = fromBalance - value;
        _balances[to] += value;

        // 触发转账事件
        emit Transfer(from, to, value);
    }
    // 授权
    function approve(address spender, uint256 amount) external returns (bool){
        address owner = msg.sender;
        _approve(owner, spender, amount);
        return true;
    }

    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");
        _allowances[owner][spender] = amount;

        emit Approve(owner, spender, amount);
    }

    function transferFrom(address from, address to, uint256 value) external returns (bool){
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        // 检查授权额度
        uint256 allowanceAmount = _allowances[from][msg.sender];
        require(allowanceAmount >= value, "ERC20: transfer amount exceeds allowance");
        // 更新授权额度
        _approve(from, msg.sender, value);
        // 转账
        _transfer(from, to, value);
        return true;
    }
    // 铸币功能（仅所有者）
    function mint(address account, uint256 value) public onlyOwner {
        _mint(account, value);
    }
    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "ERC20: mint to the zero address");
        
        _totalSupply += amount;
        _balances[account] += amount;
        
        emit Transfer(address(0), account, amount);
    }
    // 销毁功能
    function  burn(address account, uint256 value) public {
        _burn(account, value);
    }
    function _burn(address account, uint256 amount) internal {
        require(account != address(0), "ERC20: burn from the zero address");
        require(_balances[account] >= amount, "ERC20: burn amount exceeds balance");
        _balances[account] -= amount;
        _totalSupply -= amount;
        emit Transfer(account, address(0), amount);
    }
}
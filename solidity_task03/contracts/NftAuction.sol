// SPDX-License-Identifier: MIT 
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract AuctionNft is Initializable, UUPSUpgradeable {

    struct Auction {
        // 卖家
        address seller;
        // 开始时间
        uint256 startTime;
        // 持续时间
        uint256 duration;
        // 是否结束
        bool ended;
        // 起始价格
        uint256 startPrice;
        // 最高价格
        uint256 highestBid;
        // 最高价格者
        address highestBidder;
        // nft 合约地址
        address nftAddress;
        // token ID 
        uint256 tokenId;
        // 参与竞价的资产类型
        address tokenAddress;
    }    

    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can call this function");
        _;
    }
    // 状态变量
    mapping(uint256 => Auction) public auctions;

    // 下一个拍卖ID
    uint256 public nextAuctionId;

    // 管理员
    address public admin;

    // 喂价
    mapping(address => AggregatorV3Interface) public priceFeeds;

    // 初始化函数
    function initialize() public initializer {
        admin = msg.sender;
    }
    // 设置价格
    function setPriceFeed(address _tokenAddress, address _priceFeed) public {
        priceFeeds[_tokenAddress] = AggregatorV3Interface(_priceFeed);
    }
    function getLatestPrice(address _tokenAddress) public view returns (int256) {
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
        ) = priceFeeds[_tokenAddress].latestRoundData();
        return answer;
    }

    // 创建拍卖
    function createNftAuction(uint256 _startPrice, uint256 _duration, address _nftAddress, uint256 tokenId) public onlyAdmin {
        require(_startPrice > 0, "Start price must be greater than 0");
        require(_duration > 10, "Duration must be greater than 10s");

        // 转移NFT 到合约
        IERC721(_nftAddress).safeTransferFrom(msg.sender, address(this), tokenId);

        auctions[nextAuctionId] = Auction({
            seller: msg.sender,
            startTime: block.timestamp,
            duration: _duration,
            ended: false,
            startPrice: _startPrice,
            highestBid: 0,
            highestBidder: address(0),
            nftAddress: _nftAddress,
            tokenId: tokenId,
            tokenAddress: address(0)
        });

        nextAuctionId++;
    }

    // 参与竞拍
    function placeBid(uint256 _auctionId, uint256 _amount, address _tokenAddress) external payable {
        // 检查拍卖是否结束
        Auction storage auction = auctions[_auctionId];
        require(!auction.ended && (auction.startTime + auction.duration) > block.timestamp, "Auction has ended");

        // 统一价格
        uint256 payAmount;
        if (_tokenAddress == address(0)) {
            // ERC20
            payAmount = _amount * uint256(getLatestPrice(_tokenAddress));
        }else {
            // ETH
            _amount = msg.value;
            payAmount = _amount * uint256(getLatestPrice(address(0)));
        }
        uint256 startPrice = auction.startPrice * uint256(getLatestPrice(auction.tokenAddress));
        uint256 heigthBid = auction.highestBid * uint256(getLatestPrice(auction.tokenAddress));

        // 校验价格
        require(payAmount > startPrice && payAmount > heigthBid, "Bid amount must be greater than the current highest bid");

        if(_tokenAddress == address(0)) {
            IERC20(_tokenAddress).transferFrom(msg.sender, address(this), _amount);
        }
        // 退还之前的最高价
        if(auction.highestBid > 0) {
            if(auction.tokenAddress == address(0)) {
                payable(auction.highestBidder).transfer(auction.highestBid);
            }else{
                IERC20(auction.tokenAddress).transfer(auction.highestBidder, auction.highestBid);
            }
        }

        // 更新最高价格
        auction.highestBidder = msg.sender;
        auction.tokenAddress = _tokenAddress;
        auction.highestBid = _amount;
    } 

    // 结束拍卖
    function endAuction(uint256 _auctionId) external onlyAdmin {
        Auction storage auction = auctions[_auctionId];
        require(!auction.ended && (auction.startTime + auction.duration) <= block.timestamp, "Auction has not ended");

        // 转移NFT 到最高价者
        IERC721(auction.tokenAddress).safeTransferFrom(address(this), auction.highestBidder, auction.tokenId);

        // 转移资金到卖家
        if(auction.tokenAddress == address(0)) {
            payable(auction.seller).transfer(auction.highestBid);
        }else{
            IERC20(auction.tokenAddress).transfer(auction.seller, auction.highestBid);
        }
        auction.ended = true;
    }

    // UUPS 升级
    function _authorizeUpgrade(address) internal view override {
        require(msg.sender == admin, "Only admin can upgrade");
    }

}
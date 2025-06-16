// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol"; 

// 题目2
contract MyNFT is ERC721URIStorage {
     uint256 private _nextTokenId; 

     constructor(string memory name_, string memory symbol_) ERC721(name_, symbol_) {

     }

    function mintNFT(address recipient , string memory tokenURI) public returns (uint256){
        require(recipient  != address(0), "Invalid address");
        uint256 tokenId = _nextTokenId;
        _nextTokenId++;
        _mint(recipient, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return tokenId;
    }

}
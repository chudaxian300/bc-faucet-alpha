// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Faucet {
    
    mapping(address => uint) public timeLimt;

    event Withdraw(address sender,uint award,uint userLimit);

    function withdrawOfOneMin() public payable{
        require(timeLimt[msg.sender] < block.timestamp);

        uint award =  1 ether;
        uint userLimit = block.timestamp + 60;
        timeLimt[msg.sender] = userLimit;
        emit Withdraw(msg.sender, award, userLimit);
    }
    
    function withdrawOfThreeMin() public {
        require(timeLimt[msg.sender] < block.timestamp);

        uint award =  2.5 ether;
        uint userLimit = block.timestamp + 180;
        timeLimt[msg.sender] = userLimit;
        emit Withdraw(msg.sender, award, userLimit);
    }

    function withdrawOfNineMin() public {
        require(timeLimt[msg.sender] < block.timestamp);

        uint award = 6 ether;  
        uint userLimit = block.timestamp + 540;
        timeLimt[msg.sender] = userLimit;
        emit Withdraw(msg.sender, award, userLimit);
    }
}
pragma solidity ~0.8.0

contract Voting{
    address[] condidates;
    mapping(address => uint) votes;
}
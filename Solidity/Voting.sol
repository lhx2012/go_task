// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract Voting{
    address[] private condidates;
    mapping(address => uint) votes;


    function vote(address _candidate)  public {
        votes[_candidate] += 1;
        condidates.push(_candidate);
    }

    function getVotes(address _candidate) public view returns( uint){
        return votes[_candidate];
    }

    function resetVotes() public {
        for(uint i = 0;i < condidates.length;i++){
            votes[condidates[i]] = 0;
        }
        delete condidates;
    }
}
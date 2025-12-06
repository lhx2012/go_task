// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract  ReverseString{

    function reverseString(string memory input)public  pure returns (string memory){
        bytes memory inBytes = bytes(input);
        uint length = inBytes.length;
        bytes memory reverseBytes = new bytes(length);
        for (uint i = 0; i < length; i++) 
        {
            reverseBytes[length - i - 1] = inBytes[i];
        }
        return string(reverseBytes);
    }
}
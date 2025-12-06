// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract IntConvertRoman{
    string[] romStrs = ["I", "V", "X", "L", "C", "D", "M"];

    function ConvertToRoman(int num) public view returns(string memory){

    string memory roman ="";
    uint index = 0;
    int div = num;
    int remain = 0;

    while(div >0)
    {
        remain = div%10;
        if(remain == 0) 
        {
            break;
        }

        if (remain < 4)
        {

            {
                roman = string.concat(romStrs[2 * index],roman);
            }
        }
        else if (remain == 4)
        {

            {
                roman = string.concat(romStrs[2 * index],roman);
            }
        }
        else if (remain == 4)
        {
            roman = string.concat(romStrs[2 * index] , romStrs[2 * index + 1] , roman);
        }
        else if (remain == 9)
        {
            roman = string.concat(romStrs[2 * index] , romStrs[2 * index + 2] , roman);
        }
        else
        {
            for (int i = 0; i < (remain - 5); i++)
            {
                roman = string.concat( romStrs[2 * index] , roman);
            }
            roman = string.concat(romStrs[2 * index + 1] ,roman);
        }

        index++;
        div = div / 10;

    }

    return roman;
}
}
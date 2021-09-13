pragma solidity >=0.8.0 <0.9.0;

contract SmartContractStart {
    function Hello() public view returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public view returns (string memory) {
        return str;
    }
}

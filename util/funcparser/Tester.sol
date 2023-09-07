pragma solidity ^0.8.20;
pragma experimental ABIEncoderV2;

contract Tester {
    function test() public pure {}

    function testUint8(uint8 arg1) public pure returns (uint8) {
        return arg1;
    }

    function testUint8Array(uint8[] memory arg1) public pure returns (uint8[] memory) {
        return arg1;
    }

    function testUint82DArray(uint8[][] memory arg1) public pure returns (uint8[][] memory) {
        return arg1;
    }

    function testUint16(uint16 arg1) public pure returns (uint16) {
        return arg1;
    }

    function testUint16Array(uint16[] memory arg1) public pure returns (uint16[] memory) {
        return arg1;
    }

    function testUint162DArray(uint16[][] memory arg1) public pure returns (uint16[][] memory) {
        return arg1;
    }

    function testUint32(uint32 arg1) public pure returns (uint32) {
        return arg1;
    }

    function testUint32Array(uint32[] memory arg1) public pure returns (uint32[] memory) {
        return arg1;
    }

    function testUint322DArray(uint32[][] memory arg1) public pure returns (uint32[][] memory) {
        return arg1;
    }

    function testUint64(uint64 arg1) public pure returns (uint64) {
        return arg1;
    }

    function testUint64Array(uint64[] memory arg1) public pure returns (uint64[] memory) {
        return arg1;
    }

    function testUint642DArray(uint64[][] memory arg1) public pure returns (uint64[][] memory) {
        return arg1;
    }

    function testUint128(uint128 arg1) public pure returns (uint128) {
        return arg1;
    }

    function testUint128Array(uint128[] memory arg1) public pure returns (uint128[] memory) {
        return arg1;
    }

    function testUint1282DArray(uint128[][] memory arg1) public pure returns (uint128[][] memory) {
        return arg1;
    }

    function testUint256(uint256 arg1) public pure returns (uint256) {
        return arg1;
    }

    function testUint256Array(uint256[] memory arg1) public pure returns (uint256[] memory) {
        return arg1;
    }

    function testUint2562DArray(uint256[][] memory arg1) public pure returns (uint256[][] memory) {
        return arg1;
    }

    function testInt8(int8 arg1) public pure returns (int8) {
        return arg1;
    }

    function testInt8Array(int8[] memory arg1) public pure returns (int8[] memory) {
        return arg1;
    }

    function testInt82DArray(int8[][] memory arg1) public pure returns (int8[][] memory) {
        return arg1;
    }

    function testInt16(int16 arg1) public pure returns (int16) {
        return arg1;
    }

    function testInt16Array(int16[] memory arg1) public pure returns (int16[] memory) {
        return arg1;
    }

    function testInt162DArray(int16[][] memory arg1) public pure returns (int16[][] memory) {
        return arg1;
    }

    function testInt32(int32 arg1) public pure returns (int32) {
        return arg1;
    }

    function testInt32Array(int32[] memory arg1) public pure returns (int32[] memory) {
        return arg1;
    }

    function testInt322DArray(int32[][] memory arg1) public pure returns (int32[][] memory) {
        return arg1;
    }

    function testInt64(int64 arg1) public pure returns (int64) {
        return arg1;
    }

    function testInt64Array(int64[] memory arg1) public pure returns (int64[] memory) {
        return arg1;
    }

    function testInt642DArray(int64[][] memory arg1) public pure returns (int64[][] memory) {
        return arg1;
    }

    function testInt128(int128 arg1) public pure returns (int128) {
        return arg1;
    }

    function testInt128Array(int128[] memory arg1) public pure returns (int128[] memory) {
        return arg1;
    }

    function testInt1282DArray(int128[][] memory arg1) public pure returns (int128[][] memory) {
        return arg1;
    }

    function testInt256(int256 arg1) public pure returns (int256) {
        return arg1;
    }

    function testInt256Array(int256[] memory arg1) public pure returns (int256[] memory) {
        return arg1;
    }

    function testInt2562DArray(int256[][] memory arg1) public pure returns (int256[][] memory) {
        return arg1;
    }

    function testString(string memory arg1) public pure returns (string memory) {
        return arg1;
    }

    function testStringArray(string[] memory arg1) public pure returns (string[] memory) {
        return arg1;
    }

    function testString2DArray(string[][] memory arg1) public pure returns (string[][] memory) {
        return arg1;
    }

    function testBool(bool arg1) public pure returns (bool) {
        return arg1;
    }

    function testBoolArray(bool[] memory arg1) public pure returns (bool[] memory) {
        return arg1;
    }

    function testBool2DArray(bool[][] memory arg1) public pure returns (bool[][] memory) {
        return arg1;
    }

    function testAddress(address arg1) public pure returns (address) {
        return arg1;
    }

    function testAddressArray(address[] memory arg1) public pure returns (address[] memory) {
        return arg1;
    }

    function testAddress2DArray(address[][] memory arg1) public pure returns (address[][] memory) {
        return arg1;
    }

    function testBytes(bytes memory arg1) public pure returns (bytes memory) {
        return arg1;
    }

    function testBytesArray(bytes[] memory arg1) public pure returns (bytes[] memory) {
        return arg1;
    }

    function testBytes2DArray(bytes[][] memory arg1) public pure returns (bytes[][] memory) {
        return arg1;
    }

    struct TestTuple {
      uint32 field1;
      uint64 field2;
      bool field3;
    }

    function testTuple(TestTuple memory arg1) public pure returns (TestTuple memory) {
        return arg1;
    }
}

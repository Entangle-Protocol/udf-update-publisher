//SPDX-License-Identifier: BSL 1.1
pragma solidity ^0.8.19;

contract MockEndPoint {
    /// @dev Represents a digital signature
    struct Signature {
        uint8 v; // The recovery byte
        bytes32 r; // The first 32 bytes of the signature
        bytes32 s; // The second 32 bytes of the signature
    }

    /// @notice protocol info struct
    struct AllowedProtocolInfo {
        bool isCreated;
        uint256 consensusTargetRate; // percentage of proofs div numberOfAllowedTransmitters which should be reached to approve operation. Scaled with 10000 decimals, e.g. 6000 is 60%
    }

    mapping(bytes32 protocolId => uint256) public numberOfAllowedTransmitters;
    mapping(bytes32 proocolId => mapping(address => bool))
        public allowedTransmitters;

    /// @notice protocol info map
    mapping(bytes32 protocolId => AllowedProtocolInfo)
        public allowedProtocolInfo;

    /// @notice Adding protocol to whitelist
    /// @param _protocolId protocol Id
    /// @param _consensusTargetRate consensus target rate
    /// @param _transmitters initial array of transmitters
    function addAllowedProtocol(
        bytes32 _protocolId,
        uint256 _consensusTargetRate,
        address[] calldata _transmitters
    ) external {
        allowedProtocolInfo[_protocolId].isCreated = true;
        allowedProtocolInfo[_protocolId]
            .consensusTargetRate = _consensusTargetRate;

        for (uint i; i < _transmitters.length; ) {
            allowedTransmitters[_protocolId][_transmitters[i]] = true;
            numberOfAllowedTransmitters[_protocolId]++;
            unchecked {
                ++i;
            }
        }
    }
}

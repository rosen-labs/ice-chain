package types

// ValidateBasic is used for validating the packet
func (p MsgMintRequestPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p MsgMintRequestPacketData) GetBytes() ([]byte, error) {
	var modulePacket XchainPacketData

	modulePacket.Packet = &XchainPacketData_MsgMintRequestPacket{&p}

	return modulePacket.Marshal()
}

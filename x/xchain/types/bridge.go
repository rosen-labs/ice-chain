package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgBridgeRequest(
	destChainId string,
	amount float32,
	fee float32,
	reciever string,
	Signer string,
) *MsgBridgeRequest {
	return &MsgBridgeRequest{
		Reciever:    reciever,
		Amount:      amount,
		Fee:         fee,
		DestChainId: destChainId,
	}
}

// Route ...
func (m *MsgBridgeRequest) Route() string {
	return "Bridge"
}

// Type ...
func (m *MsgBridgeRequest) Type() string {
	return "BridgeRequest"
}

func (m *MsgBridgeRequest) ValidateBasic() error {
	return nil
}

func (m *MsgBridgeRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgBridgeRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

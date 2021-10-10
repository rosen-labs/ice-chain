package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/rosen-labs/xchain/x/xchain/types"
)

// TransmitMsgMintRequestPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitMsgMintRequestPacket(
	ctx sdk.Context,
	packetData types.MsgMintRequestPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvMsgMintRequestPacket processes packet reception
func (k Keeper) OnRecvMsgMintRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgMintRequestPacketData) (packetAck types.MsgMintRequestPacketAck, err error) {
	fmt.Println("DEBUG : start mint proccess")
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}
	revAddress, uid, err := types.GetRevAddressAndUID(data.Reciever)
	if err != nil {
		return packetAck, err
	}
	fmt.Printf("DEBUG : get rev = %s | uid = %s\n", revAddress, uid)

	totalCoins := sdk.NewCoins(sdk.NewCoin("token", sdk.NewIntFromUint64(data.Amount)))

	fmt.Println("DEBUG : create reciever account")
	recieverAccount, err := sdk.AccAddressFromBech32(revAddress)
	if err != nil {
		return packetAck, err
	}

	fmt.Println("DEBUG : mint coin to module account")
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, totalCoins); err != nil {
		return packetAck, err
	}

	fmt.Println("DEBUG : send module to account")
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recieverAccount, totalCoins); err != nil {
		return packetAck, err
	}

	ctx.EventManager().EmitEvents([]sdk.Event{
		sdk.NewEvent(
			types.EventTypeBridgingMint,
			sdk.NewAttribute(types.AttributeKeyUID, uid),
			sdk.NewAttribute(types.AttributeKeyEventName, types.EventTypeBridgingMint),
			sdk.NewAttribute(types.AttributeKeyReciever, revAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", data.Amount)),
			sdk.NewAttribute(types.AttributeKeyFee, fmt.Sprintf("%d", data.Fee)),
			sdk.NewAttribute(types.AttributeKeySrcChainId, fmt.Sprintf("%d", data.SrcChainId)),
			sdk.NewAttribute(types.AttributeKeyDestChainId, fmt.Sprintf("%d", data.SrcChainId)),
			sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%d", data.TokenId)),
			sdk.NewAttribute(types.AttributeKeyDenom, "token"),
		),
	})

	return packetAck, nil
}

// OnAcknowledgementMsgMintRequestPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementMsgMintRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgMintRequestPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.MsgMintRequestPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutMsgMintRequestPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutMsgMintRequestPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgMintRequestPacketData) error {

	// TODO: packet timeout logic

	return nil
}

package xchain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rosen-labs/xchain/x/xchain/keeper"
	"github.com/rosen-labs/xchain/x/xchain/types"
)

func handleMsgBridgeRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBridgeRequest) (*sdk.Result, error) {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}
	totalAmount := sdk.NewCoin("token", sdk.NewIntFromUint64(msg.Amount))
	totalAmount = totalAmount.Add(sdk.NewCoin("token", sdk.NewIntFromUint64(msg.Fee)))
	totalInVouchers := sdk.Coins{totalAmount}

	if err := k.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, totalInVouchers); err != nil {
		return nil, err
	}

	if err := k.BurnCoin(ctx, types.ModuleName, totalInVouchers); err != nil {
		return nil, err
	}

	if err := k.SendMintRequest(
		ctx,
		msg.Reciever,
		msg.Amount,
		msg.Fee,
		0,
		1,
		msg.DestChainId,
	); err != nil {
		fmt.Println("DEBUG : send mint request fail", err)
		return nil, err
	}

	revAddress, uid, err := types.GetRevAddressAndUID(msg.Reciever)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DEBUG : get rev = %s | uid = %s\n", revAddress, uid)
	ctx.EventManager().EmitEvents([]sdk.Event{
		sdk.NewEvent(
			types.EventTypeBridgingMint,
			sdk.NewAttribute(types.AttributeKeyUID, uid),
			sdk.NewAttribute(types.AttributeKeyEventName, types.EventTypeBridgingMint),
			sdk.NewAttribute(types.AttributeKeyReciever, revAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", msg.Amount)),
			sdk.NewAttribute(types.AttributeKeyFee, fmt.Sprintf("%d", msg.Fee)),
			sdk.NewAttribute(types.AttributeKeySrcChainId, fmt.Sprintf("%d", 1)),
			sdk.NewAttribute(types.AttributeKeyDestChainId, fmt.Sprintf("%d", msg.DestChainId)),
			sdk.NewAttribute(types.AttributeKeyTokenID, fmt.Sprintf("%d", 0)),
			sdk.NewAttribute(types.AttributeKeyDenom, "token"),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

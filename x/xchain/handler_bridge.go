package xchain

import (
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
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

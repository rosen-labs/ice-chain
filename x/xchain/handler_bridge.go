package xchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rosen-labs/xchain/x/xchain/keeper"
	"github.com/rosen-labs/xchain/x/xchain/types"
)

func handleMsgBridgeRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBridgeRequest) (*sdk.Result, error) {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	totalInVouchers := sdk.Coins{*msg.Amount}
	totalInVouchers = totalInVouchers.Add(*msg.Fee)

	//TODO : fix panic bug from bank keeper
	// balance := k.GetAllBalances(ctx, signer)
	// ctx.Logger().Info("DEBUG : ", balance)
	// if err := k.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, totalInVouchers); err != nil {
	// 	return nil, err
	// }

	// if err := k.BurnCoin(ctx, types.ModuleName, totalInVouchers); err != nil {
	// 	return nil, err
	// }

	//TODO : send ibc protocal

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

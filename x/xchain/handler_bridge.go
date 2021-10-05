package xchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rosen-labs/xchain/x/xchain/keeper"
	"github.com/rosen-labs/xchain/x/xchain/types"
)

func handleMsgBridgeRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBridgeRequest) (*sdk.Result, error) {

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

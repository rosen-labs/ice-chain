package keeper

import (
	"github.com/rosen-labs/xchain/x/xchain/types"
)

var _ types.QueryServer = Keeper{}

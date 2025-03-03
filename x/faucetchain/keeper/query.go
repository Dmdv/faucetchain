package keeper

import (
	"faucetchain/x/faucetchain/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"faucetchain/x/faucet/types"
)

var _ types.QueryServer = Keeper{}

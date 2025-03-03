package keeper

import (
	"context"

	"faucetchain/x/faucet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestTokens(goCtx context.Context, msg *types.MsgRequestTokens) (*types.MsgRequestTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRequestTokensResponse{}, nil
}

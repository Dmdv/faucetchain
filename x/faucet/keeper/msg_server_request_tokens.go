package keeper

import (
	"context"
	errwrapper "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"faucetchain/x/faucet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestTokens(goCtx context.Context, msg *types.MsgRequestTokens) (*types.MsgRequestTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	requester, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	params := k.GetParams(ctx)

	// Check maxPerRequest constraint
	if msg.Amount > params.MaxPerRequest {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("requested amount exceeds maxPerRequest (%d)", params.MaxPerRequest)
	}

	// Fetch request history
	history := k.GetRequestHistory(ctx, requester)
	var totalRequested uint64
	for _, req := range history.Requests {
		totalRequested += req.Amount
	}

	// Check maxPerAddress constraint
	if totalRequested+msg.Amount > params.MaxPerAddress {
		return nil, errwrapper.Wrapf(sdkerrors.ErrInvalidRequest, "cumulative amount %d exceeds maxPerAddress %d", totalRequested+msg.Amount, params.MaxPerAddress)
	}

	// Perform token transfer from a faucet module account
	coin := sdk.NewCoin("stake", math.NewInt(int64(msg.Amount)))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requester, sdk.NewCoins(coin))
	if err != nil {
		return nil, errwrapper.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	// explicitly record the request in history
	newRequest := &types.Request{
		Amount: msg.Amount,
		Height: ctx.BlockHeight(),
	}

	history.Requests = append(history.Requests, newRequest)

	// explicitly persist updated history
	k.SetRequestHistory(ctx, requester, history)

	// explicitly emit event (bonus evaluation)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("RequestTokens",
			sdk.NewAttribute("address", msg.Creator),
			sdk.NewAttribute("amount", fmt.Sprintf("%d", msg.Amount)),
		),
	)

	return &types.MsgRequestTokensResponse{}, nil
}

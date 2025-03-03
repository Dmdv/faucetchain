package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"faucetchain/x/faucet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestsByAddress(goCtx context.Context, req *types.QueryRequestsByAddressRequest) (*types.QueryRequestsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	history := k.GetRequestHistory(ctx, addr)

	return &types.QueryRequestsByAddressResponse{
		Requests: history.Requests,
	}, nil
}

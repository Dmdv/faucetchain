package keeper

import (
	"faucetchain/x/faucet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRequestHistory(ctx sdk.Context, addr sdk.AccAddress, history types.QueryRequestsByAddressResponse) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&history)
	store.Set(requestHistoryKey(addr), bz)
}

func (k Keeper) GetRequestHistory(ctx sdk.Context, addr sdk.AccAddress) types.QueryRequestsByAddressResponse {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(requestHistoryKey(addr))
	if bz == nil {
		return types.QueryRequestsByAddressResponse{}
	}
	var history types.QueryRequestsByAddressResponse
	k.cdc.MustUnmarshal(bz, &history)
	return history
}

func requestHistoryKey(addr sdk.AccAddress) []byte {
	return append([]byte("RequestHistory-"), addr.Bytes()...)
}

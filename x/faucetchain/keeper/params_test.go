package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "faucetchain/testutil/keeper"
	"faucetchain/x/faucetchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FaucetchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

package faucetchain_test

import (
	"testing"

	keepertest "faucetchain/testutil/keeper"
	"faucetchain/testutil/nullify"
	faucetchain "faucetchain/x/faucetchain/module"
	"faucetchain/x/faucetchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FaucetchainKeeper(t)
	faucetchain.InitGenesis(ctx, k, genesisState)
	got := faucetchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

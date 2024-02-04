package mydist_test

import (
	"testing"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	mydist "example/x/mydist/module"
	"example/x/mydist/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MydistKeeper(t)
	mydist.InitGenesis(ctx, k, genesisState)
	got := mydist.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

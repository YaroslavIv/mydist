package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "example/testutil/keeper"
	"example/x/mydist/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MydistKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

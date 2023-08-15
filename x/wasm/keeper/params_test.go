package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "testWasm/testutil/keeper"
	"testWasm/x/wasm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WasmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

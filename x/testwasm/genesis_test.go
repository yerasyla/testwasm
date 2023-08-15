package testwasm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testWasm/testutil/keeper"
	"testWasm/testutil/nullify"
	"testWasm/x/testwasm"
	"testWasm/x/testwasm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestwasmKeeper(t)
	testwasm.InitGenesis(ctx, *k, genesisState)
	got := testwasm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

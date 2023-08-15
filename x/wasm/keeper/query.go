package keeper

import (
	"testWasm/x/wasm/types"
)

var _ types.QueryServer = Keeper{}

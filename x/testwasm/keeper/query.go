package keeper

import (
	"testWasm/x/testwasm/types"
)

var _ types.QueryServer = Keeper{}

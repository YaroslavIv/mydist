package keeper

import (
	"example/x/mydist/types"
)

var _ types.QueryServer = Keeper{}

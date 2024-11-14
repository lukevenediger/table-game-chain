package keeper

import (
	"table-game-chain/x/checkers/types"
)

var _ types.QueryServer = Keeper{}

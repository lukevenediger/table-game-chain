package types_test

import (
	"table-game-chain/testutil/sample"
	"table-game-chain/x/checkers/types"
	"time"
)

// NewValidGame returns a valid game with sample addresses and start time set to now
func NewValidGame() *types.Game {
	addresses := sample.AccAddresses(2)
	return types.NewGame(addresses[0],
		addresses[0],
		addresses[1],
		time.Now(),
	)
}

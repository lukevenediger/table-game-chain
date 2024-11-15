package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "table-game-chain/testutil/keeper"
	"table-game-chain/x/checkers/keeper"
	"table-game-chain/x/checkers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGameMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.CheckersKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	index := "B"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateGame{Creator: creator, Index: index}
		_, err := srv.CreateGame(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetGame(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

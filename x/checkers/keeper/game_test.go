package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "table-game-chain/testutil/keeper"
	"table-game-chain/testutil/nullify"
	"table-game-chain/x/checkers/keeper"
	"table-game-chain/x/checkers/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGame(keeper keeper.Keeper, ctx context.Context, n int) []types.Game {
	items := make([]types.Game, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGame(ctx, items[i])
	}
	return items
}

func TestGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGame(ctx,
			item.Index,
		)
		_, found := keeper.GetGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGame(ctx)),
	)
}

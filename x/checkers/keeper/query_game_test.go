package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "table-game-chain/testutil/keeper"
	"table-game-chain/testutil/nullify"
	"table-game-chain/x/checkers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGameQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	msgs := createNGame(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGameRequest
		response *types.QueryGetGameResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGameRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetGameResponse{Game: *msgs[0].Game},
		},
		{
			desc: "Second",
			request: &types.QueryGetGameRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetGameResponse{Game: *msgs[1].Game},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGameRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Game(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestGameQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	msgs := createNGame(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGameRequest {
		return &types.QueryAllGameRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GameAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Games), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Games),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GameAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Games), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Games),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GameAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Games),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GameAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

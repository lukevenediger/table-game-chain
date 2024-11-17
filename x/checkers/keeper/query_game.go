package keeper

import (
	"context"

	"table-game-chain/x/checkers/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GameAll returns a page of indexed games
func (k Keeper) GameAll(ctx context.Context, req *types.QueryAllGameRequest) (*types.QueryAllGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var indexedGames []types.IndexedGame

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	gameStore := prefix.NewStore(store, types.KeyPrefix(types.GameKeyPrefix))

	pageRes, err := query.Paginate(gameStore, req.Pagination, func(key []byte, value []byte) error {
		var game types.Game
		if err := k.cdc.Unmarshal(value, &game); err != nil {
			return err
		}

		indexedGames = append(indexedGames, types.IndexedGame{
			// TODO: super sus about whether this conversion will work :/
			Index: string(key),
			Game:  &game,
		})
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGameResponse{Games: indexedGames, Pagination: pageRes}, nil
}

// Game returns a game from its index
func (k Keeper) Game(ctx context.Context, req *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetGame(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGameResponse{Game: val}, nil
}

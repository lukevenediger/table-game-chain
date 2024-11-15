package keeper

import (
	"context"

	"table-game-chain/x/checkers/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetGame set a specific game in the store from its index
func (k Keeper) SetGame(ctx context.Context, game types.IndexedGame) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKeyPrefix))
	b := k.cdc.MustMarshal(&game)
	store.Set(types.GameKey(
		game.Index,
	), b)
}

// GetGame returns a game from its index
func (k Keeper) GetGame(
	ctx context.Context,
	index string,

) (val types.Game, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKeyPrefix))

	b := store.Get(types.GameKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGame removes a game from the store
func (k Keeper) RemoveGame(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKeyPrefix))
	store.Delete(types.GameKey(
		index,
	))
}

// GetAllGame returns all game
func (k Keeper) GetAllGame(ctx context.Context) (list []types.Game) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GameKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Game
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

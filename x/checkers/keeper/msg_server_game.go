package keeper

import (
	"context"

	"table-game-chain/x/checkers/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetGame(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var game = types.NewGame(
		msg.Creator,
		msg.Black,
		msg.Red,
		ctx.BlockTime(),
	)

	k.SetGame(
		ctx,
		types.NewIndexedGame(
			msg.Index,
			game,
		),
	)
	return &types.MsgCreateGameResponse{}, nil
}

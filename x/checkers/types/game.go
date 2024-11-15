package types

import (
	"table-game-chain/x/checkers/rules"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGame(creator, index, black, red string, startTime int64) Game {

	newBoard := rules.New()

	return Game{
		Creator:   black,
		Black:     black,
		Red:       red,
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		State:     "",
		StartTime: startTime,
		EndTime:   0,
	}
}

// GetBlackAddress validates and returns the black player's address
func (g Game) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, parseErr := sdk.AccAddressFromBech32(g.Black)
	return black, errors.Wrapf(parseErr, ErrInvalidBlack.Error(), g.Black)
}

// GetRedAddress validates and returns the red player's address
func (g Game) GetRedAddress() (red sdk.AccAddress, err error) {
	red, parseErr := sdk.AccAddressFromBech32(g.Red)
	return red, errors.Wrapf(parseErr, ErrInvalidRed.Error(), g.Red)
}

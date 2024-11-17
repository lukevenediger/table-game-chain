package types

import (
	"table-game-chain/x/checkers/rules"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karrick/tparse"
)

// NewGame creates a new checkers game
func NewGame(creator, black, red string, startTime time.Time) *Game {

	newBoard := rules.New()

	return &Game{
		Creator:   creator,
		Black:     black,
		Red:       red,
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		State:     "",
		StartTime: FormatWithGameTimeLayout(startTime),
		EndTime:   "",
	}
}

// GetBlackAddress validates and returns the black player's address
func (g *Game) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, parseErr := sdk.AccAddressFromBech32(g.Black)
	return black, errors.Wrapf(parseErr, ErrInvalidBlack.Error(), g.Black)
}

// GetRedAddress validates and returns the red player's address
func (g *Game) GetRedAddress() (red sdk.AccAddress, err error) {
	red, parseErr := sdk.AccAddressFromBech32(g.Red)
	return red, errors.Wrapf(parseErr, ErrInvalidRed.Error(), g.Red)
}

// GetDeadlineAsTime returns the deadline of the game as a time.Time object
func (g *Game) GetDeadlineAsTime() (deadline time.Time, err error) {
	// Parse the start time
	startTime, parseErr := time.Parse(GameTimeLayout, g.StartTime)
	if parseErr != nil {
		return time.Time{},
			errors.Wrapf(parseErr, ErrCannotParseGameTime.Error(), g.StartTime)
	}

	deadline, parseErr = tparse.AddDuration(startTime, MaxTurnDuration)
	if parseErr != nil {
		return time.Time{},
			errors.Wrapf(parseErr, ErrCannotParseDeadline.Error(), MaxTurnDuration)
	}

	return deadline, nil
}

// FormatWithGameTimeLayout returns the game start time as a formatted string
func FormatWithGameTimeLayout(t time.Time) string {
	return t.Format(GameTimeLayout)
}

// NewIndexedGame creates a new indexed game
func NewIndexedGame(index string, game *Game) IndexedGame {
	return IndexedGame{
		Index: index,
		Game:  game,
	}
}

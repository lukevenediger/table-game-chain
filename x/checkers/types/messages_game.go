package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(
	creator string,
	index string,
	board string,
	red string,
	black string,
	turn string,
	state string,
	startTime int32,
	endTime int32,

) *MsgCreateGame {
	return &MsgCreateGame{
		Creator: creator,
		Index:   index,
		Red:     red,
		Black:   black,
	}
}

func (msg *MsgCreateGame) ValidateBasic() error {
	// Validate creator address
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validate red address
	_, err = sdk.AccAddressFromBech32(msg.Red)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidRed, "invalid red address (%s)", err)
	}

	//  Validate black address
	_, err = sdk.AccAddressFromBech32(msg.Black)
	if err != nil {
		return errorsmod.Wrapf(ErrInvalidBlack, "invalid black address (%s)", err)
	}

	return nil
}

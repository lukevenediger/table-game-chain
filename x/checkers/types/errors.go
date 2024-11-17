package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidSigner       = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidBlack        = sdkerrors.Register(ModuleName, 1102, "invalid black address: %s")
	ErrInvalidRed          = sdkerrors.Register(ModuleName, 1103, "invalid red address: %s")
	ErrCannotParseGameTime = sdkerrors.Register(ModuleName, 1104, "cannot parse game time: %s")
	ErrCannotParseDeadline = sdkerrors.Register(ModuleName, 1105, "cannot parse deadline: %s")
)

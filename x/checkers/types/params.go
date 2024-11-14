package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyInactivityTimeout = []byte("InactivityTimeout")
	// TODO: Determine the default value
	DefaultInactivityTimeout string = "inactivity_timeout"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	inactivityTimeout string,
) Params {
	return Params{
		InactivityTimeout: inactivityTimeout,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultInactivityTimeout,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyInactivityTimeout, &p.InactivityTimeout, validateInactivityTimeout),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateInactivityTimeout(p.InactivityTimeout); err != nil {
		return err
	}

	return nil
}

// validateInactivityTimeout validates the InactivityTimeout param
func validateInactivityTimeout(v interface{}) error {
	inactivityTimeout, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = inactivityTimeout

	return nil
}

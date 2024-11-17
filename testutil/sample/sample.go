package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// AccAddresses returns a slice of sample account addresses
func AccAddresses(n int) []string {
	addresses := make([]string, n)
	for i := range addresses {
		addresses[i] = AccAddress()
	}
	return addresses
}

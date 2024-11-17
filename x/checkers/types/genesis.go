package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GameList: []IndexedGame{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in game
	gameIndexMap := make(map[string]struct{})

	for _, elem := range gs.GameList {
		index := string(GameKey(elem.Index))
		if _, ok := gameIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for game")
		}
		gameIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

package types_test

import (
	"testing"

	"table-game-chain/x/checkers/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				GameList: []types.IndexedGame{
					{
						Index: "0",
						Game:  NewValidGame(),
					},
					{
						Index: "1",
						Game:  NewValidGame(),
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated game",
			genState: &types.GenesisState{
				GameList: []types.IndexedGame{
					{
						Index: "0",
						Game:  NewValidGame(),
					},
					{
						Index: "0",
						Game:  NewValidGame(),
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

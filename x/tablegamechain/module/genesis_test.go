package tablegamechain_test

import (
	"testing"

	keepertest "table-game-chain/testutil/keeper"
	"table-game-chain/testutil/nullify"
	tablegamechain "table-game-chain/x/tablegamechain/module"
	"table-game-chain/x/tablegamechain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TablegamechainKeeper(t)
	tablegamechain.InitGenesis(ctx, k, genesisState)
	got := tablegamechain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package checkers

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"table-game-chain/testutil/sample"
	checkerssimulation "table-game-chain/x/checkers/simulation"
	"table-game-chain/x/checkers/types"
)

// avoid unused import issue
var (
	_ = checkerssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgUpdateGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGame int = 100

	opWeightMsgDeleteGame = "op_weight_msg_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGame int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	checkersGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		GameList: []types.IndexedGame{
			{
				Index: "0",
				Game: types.NewGame(
					accs[0],
					accs[0],
					accs[1],
					simState.GenTimestamp,
				),
			},
			{
				Index: "1",
				Game: types.NewGame(
					accs[0],
					accs[0],
					accs[1],
					simState.GenTimestamp,
				),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&checkersGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		checkerssimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGame,
			defaultWeightMsgCreateGame,
			func(_ *rand.Rand, _ sdk.Context, _ []simtypes.Account) sdk.Msg {
				checkerssimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

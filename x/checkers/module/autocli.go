package checkers

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "table-game-chain/api/tablegamechain/checkers"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "GameAll",
					Use:       "list-game",
					Short:     "List all game",
				},
				{
					RpcMethod:      "Game",
					Use:            "show-game [id]",
					Short:          "Shows a game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateGame",
					Use:            "create-game [index] [board] [red] [black] [turn] [state] [startTime] [endTime]",
					Short:          "Create a new game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "board"}, {ProtoField: "red"}, {ProtoField: "black"}, {ProtoField: "turn"}, {ProtoField: "state"}, {ProtoField: "startTime"}, {ProtoField: "endTime"}},
				},
				{
					RpcMethod:      "UpdateGame",
					Use:            "update-game [index] [board] [red] [black] [turn] [state] [startTime] [endTime]",
					Short:          "Update game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "board"}, {ProtoField: "red"}, {ProtoField: "black"}, {ProtoField: "turn"}, {ProtoField: "state"}, {ProtoField: "startTime"}, {ProtoField: "endTime"}},
				},
				{
					RpcMethod:      "DeleteGame",
					Use:            "delete-game [index]",
					Short:          "Delete game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

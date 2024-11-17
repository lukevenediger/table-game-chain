package types

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"

	// GameTimeLayout defines the time layout for start and end game time
	GameTimeLayout = "2006-01-02 15:04:05.999999999 +0000 UTC"

	// MaxTurnDuration
	MaxTurnDuration = "1w" // 1 week
)

var (
	ParamsKey = []byte("p_checkers")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

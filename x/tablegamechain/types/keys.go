package types

const (
	// ModuleName defines the module name
	ModuleName = "tablegamechain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tablegamechain"
)

var (
	ParamsKey = []byte("p_tablegamechain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "faucetchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_faucetchain"
)

var (
	ParamsKey = []byte("p_faucetchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyMaxPerRequest = []byte("MaxPerRequest")
	KeyMaxPerAddress = []byte("MaxPerAddress")

	DefaultMaxPerRequest uint64 = 1000
	DefaultMaxPerAddress uint64 = 10000
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(maxRequest, maxAddress uint64) Params {
	return Params{
		MaxPerRequest: maxRequest,
		MaxPerAddress: maxAddress,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMaxPerRequest, DefaultMaxPerAddress)
}

// ParamSetPairs get the params.ParamSet
func (p Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxPerRequest, &p.MaxPerRequest, validateUint64),
		paramtypes.NewParamSetPair(KeyMaxPerAddress, &p.MaxPerAddress, validateUint64),
	}
}

func validateUint64(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

package types

import (
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
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

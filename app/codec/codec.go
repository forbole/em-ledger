package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/supply"
	supplye "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

var (
	_ auth.Codec   = new(Codec)
	_ supply.Codec = new(Codec)
)

// Codec defines the application-level codec. This codec contains all the
// required module-specific codecs that are to be provided upon initialization.
type Codec struct {
	codec.Marshaler

	// Keep reference to the amino codec to allow backwards compatibility along
	// with type, and interface registration.
	Amino *codec.Codec
}

func (c *Codec) MarshalSupply(supplyI supplye.SupplyI) ([]byte, error) {
	supply := &Supply{}
	if err := supply.SetSupplyI(supplyI); err != nil {
		return nil, err
	}

	return c.Marshaler.MarshalBinaryLengthPrefixed(supply)
}

func (c *Codec) UnmarshalSupply(bz []byte) (supplye.SupplyI, error) {
	supply := &Supply{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, supply); err != nil {
		return nil, err
	}

	return supply.GetSupplyI(), nil
}

func (c *Codec) MarshalSupplyJSON(supply supplye.SupplyI) ([]byte, error) {
	return c.Amino.MarshalJSON(supply)
}

func (c *Codec) UnmarshalSupplyJSON(bz []byte) (supplye.SupplyI, error) {
	var s supplye.SupplyI
	err := c.Amino.UnmarshalJSON(bz, s)
	return s, err
}

func (c *Codec) MarshalAccount(accI exported.Account) ([]byte, error) {
	acc := &Account{}
	if err := acc.SetAccount(accI); err != nil {
		return nil, err
	}

	return c.Marshaler.MarshalBinaryLengthPrefixed(acc)
}

func (c *Codec) UnmarshalAccount(bz []byte) (exported.Account, error) {
	acc := &Account{}
	if err := c.Marshaler.UnmarshalBinaryLengthPrefixed(bz, acc); err != nil {
		return nil, err
	}

	return acc.GetAccount(), nil
}

func (c *Codec) MarshalAccountJSON(acc exported.Account) ([]byte, error) {
	return c.Marshaler.MarshalJSON(acc)
}

func (c *Codec) UnmarshalAccountJSON(bz []byte) (exported.Account, error) {
	var acc exported.Account
	err := c.Marshaler.UnmarshalJSON(bz, acc)
	return acc, err
}

func NewAppCodec(amino *codec.Codec) *Codec {
	return &Codec{codec.NewHybridCodec(amino), amino}
}

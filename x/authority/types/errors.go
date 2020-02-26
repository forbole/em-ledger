// This software is Copyright (c) 2019 e-Money A/S. It is not offered under an open source license.
//
// Please contact partners@e-money.com for licensing related questions.

package types

import (
	//sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//type CodeType = sdk.CodeType

//const (
//	Codespace sdk.CodespaceType = ModuleName
//
//	CodeNotAuthority        sdk.CodeType = 1
//	CodeMissingDenomination sdk.CodeType = 2
//	CodeInvalidDenomination sdk.CodeType = 3
//	CodeNoAuthority         sdk.CodeType = 4
//	CodeInvalidGasPrices    sdk.CodeType = 5
//	CodeUnknownDenomination sdk.CodeType = 6
//)

var (
	ErrNotAuthority          = sdkerrors.Register(ModuleName, 1, "not an authority")
	ErrNoDenomsSpecified     = sdkerrors.Register(ModuleName, 2, "No denominations specified in authority call")
	ErrInvalidDenom          = sdkerrors.Register(ModuleName, 3, "Invalid denomination found")
	ErrNoAuthorityConfigured = sdkerrors.Register(ModuleName, 4, "No authority configured")
	ErrInvalidGasPrices      = sdkerrors.Register(ModuleName, 5, "Invalid gas prices")
	ErrUnknownDenom          = sdkerrors.Register(ModuleName, 6, "Unknown denomination specified")
)

//
//func ErrNotAuthority(address string) sdk.Error {
//	return sdk.NewError(Codespace, CodeMissingDenomination, "%v is not the authority", address)
//}
//
//func ErrNoDenomsSpecified() sdk.Error {
//	return sdk.NewError(Codespace, CodeMissingDenomination, "No denominations specified in authority call")
//}
//
//func ErrInvalidDenom(denom string) sdk.Error {
//	return sdk.NewError(Codespace, CodeInvalidDenomination, "Invalid denomination found: %v", denom)
//}
//
//func ErrNoAuthorityConfigured() sdk.Error {
//	return sdk.NewError(Codespace, CodeNoAuthority, "No authority configured")
//}
//
//func ErrInvalidGasPrices(amt string) sdk.Error {
//	return sdk.NewError(Codespace, CodeInvalidGasPrices, "Invalid gas prices : %v", amt)
//}
//
//func ErrUnknownDenom(denom string) sdk.Error {
//	return sdk.NewError(Codespace, CodeUnknownDenomination, "Unknown denomination specified: %v", denom)
//}

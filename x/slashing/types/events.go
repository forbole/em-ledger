//noalias
package types

// Slashing module event types
const (
	EventTypeSlash         = "slash"
	EventTypeLiveness      = "liveness"
	EventTypePenaltyPayout = "penalty_payout"

	AttributeKeyAddress      = "address"
	AttributeKeyHeight       = "height"
	AttributeKeyPower        = "power"
	AttributeKeyReason       = "reason"
	AttributeKeyJailed       = "jailed"
	AttributeKeyMissedBlocks = "missed_blocks"
	AttributeKeyAmount       = "amount"

	AttributeValueDoubleSign       = "double_sign"
	AttributeValueMissingSignature = "missing_signature"
	AttributeValueCategory         = ModuleName
)

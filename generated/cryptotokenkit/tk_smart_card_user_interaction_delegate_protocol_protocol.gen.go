// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PTKSmartCardUserInteractionDelegate is the TKSmartCardUserInteractionDelegate protocol interface.
//
// The interface implemented by a Smart Card user interaction delegate to handle user interaction events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS +
//   - iPadOS +
//   - macOS 10.11+
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.cryptotokenkit/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate
type PTKSmartCardUserInteractionDelegate interface {
	// Optional methods
	CharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasCharacterEnteredInUserInteraction() bool
	CorrectionKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasCorrectionKeyPressedInUserInteraction() bool
	InvalidCharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasInvalidCharacterEnteredInUserInteraction() bool
	NewPINConfirmationRequestedInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasNewPINConfirmationRequestedInUserInteraction() bool
	NewPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasNewPINRequestedInUserInteraction() bool
	OldPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasOldPINRequestedInUserInteraction() bool
	ValidationKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction)
	HasValidationKeyPressedInUserInteraction() bool
}

// TKSmartCardUserInteractionDelegate is a delegate implementation builder for the PTKSmartCardUserInteractionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TKSmartCardUserInteractionDelegate struct {
	_CharacterEnteredInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_CorrectionKeyPressedInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_InvalidCharacterEnteredInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_NewPINConfirmationRequestedInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_NewPINRequestedInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_OldPINRequestedInUserInteraction func(interaction ITKSmartCardUserInteraction)
	_ValidationKeyPressedInUserInteraction func(interaction ITKSmartCardUserInteraction)
}

// SetCharacterEnteredInUserInteraction sets the handler for the CharacterEnteredInUserInteraction delegate method.
//
// Tells the delegate that a valid character has been entered.
func (d *TKSmartCardUserInteractionDelegate) SetCharacterEnteredInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._CharacterEnteredInUserInteraction = f
}

// SetCorrectionKeyPressedInUserInteraction sets the handler for the CorrectionKeyPressedInUserInteraction delegate method.
//
// Tells the delegate that a correction key has been pressed.
func (d *TKSmartCardUserInteractionDelegate) SetCorrectionKeyPressedInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._CorrectionKeyPressedInUserInteraction = f
}

// SetInvalidCharacterEnteredInUserInteraction sets the handler for the InvalidCharacterEnteredInUserInteraction delegate method.
//
// Tells the delegate that an invalid character has been entered.
func (d *TKSmartCardUserInteractionDelegate) SetInvalidCharacterEnteredInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._InvalidCharacterEnteredInUserInteraction = f
}

// SetNewPINConfirmationRequestedInUserInteraction sets the handler for the NewPINConfirmationRequestedInUserInteraction delegate method.
//
// Tells the delegate that the new PIN needs to be re-entered for confirmation.
func (d *TKSmartCardUserInteractionDelegate) SetNewPINConfirmationRequestedInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._NewPINConfirmationRequestedInUserInteraction = f
}

// SetNewPINRequestedInUserInteraction sets the handler for the NewPINRequestedInUserInteraction delegate method.
//
// Tells the delegate that the new PIN needs to be entered.
func (d *TKSmartCardUserInteractionDelegate) SetNewPINRequestedInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._NewPINRequestedInUserInteraction = f
}

// SetOldPINRequestedInUserInteraction sets the handler for the OldPINRequestedInUserInteraction delegate method.
//
// Tells the delegate that the old PIN needs to be entered.
func (d *TKSmartCardUserInteractionDelegate) SetOldPINRequestedInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._OldPINRequestedInUserInteraction = f
}

// SetValidationKeyPressedInUserInteraction sets the handler for the ValidationKeyPressedInUserInteraction delegate method.
//
// Tells the delegate that the validation key has been pressed, indicating the end of PIN entry.
func (d *TKSmartCardUserInteractionDelegate) SetValidationKeyPressedInUserInteraction(f func(interaction ITKSmartCardUserInteraction)) {
	d._ValidationKeyPressedInUserInteraction = f
}

// CharacterEnteredInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) CharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._CharacterEnteredInUserInteraction != nil {
		d._CharacterEnteredInUserInteraction(interaction)
	}
}

// HasCharacterEnteredInUserInteraction returns true if a handler for CharacterEnteredInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasCharacterEnteredInUserInteraction() bool {
	return d._CharacterEnteredInUserInteraction != nil
}

// CorrectionKeyPressedInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) CorrectionKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._CorrectionKeyPressedInUserInteraction != nil {
		d._CorrectionKeyPressedInUserInteraction(interaction)
	}
}

// HasCorrectionKeyPressedInUserInteraction returns true if a handler for CorrectionKeyPressedInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasCorrectionKeyPressedInUserInteraction() bool {
	return d._CorrectionKeyPressedInUserInteraction != nil
}

// InvalidCharacterEnteredInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) InvalidCharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._InvalidCharacterEnteredInUserInteraction != nil {
		d._InvalidCharacterEnteredInUserInteraction(interaction)
	}
}

// HasInvalidCharacterEnteredInUserInteraction returns true if a handler for InvalidCharacterEnteredInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasInvalidCharacterEnteredInUserInteraction() bool {
	return d._InvalidCharacterEnteredInUserInteraction != nil
}

// NewPINConfirmationRequestedInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) NewPINConfirmationRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._NewPINConfirmationRequestedInUserInteraction != nil {
		d._NewPINConfirmationRequestedInUserInteraction(interaction)
	}
}

// HasNewPINConfirmationRequestedInUserInteraction returns true if a handler for NewPINConfirmationRequestedInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasNewPINConfirmationRequestedInUserInteraction() bool {
	return d._NewPINConfirmationRequestedInUserInteraction != nil
}

// NewPINRequestedInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) NewPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._NewPINRequestedInUserInteraction != nil {
		d._NewPINRequestedInUserInteraction(interaction)
	}
}

// HasNewPINRequestedInUserInteraction returns true if a handler for NewPINRequestedInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasNewPINRequestedInUserInteraction() bool {
	return d._NewPINRequestedInUserInteraction != nil
}

// OldPINRequestedInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) OldPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._OldPINRequestedInUserInteraction != nil {
		d._OldPINRequestedInUserInteraction(interaction)
	}
}

// HasOldPINRequestedInUserInteraction returns true if a handler for OldPINRequestedInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasOldPINRequestedInUserInteraction() bool {
	return d._OldPINRequestedInUserInteraction != nil
}

// ValidationKeyPressedInUserInteraction implements the PTKSmartCardUserInteractionDelegate interface.
func (d *TKSmartCardUserInteractionDelegate) ValidationKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	if d._ValidationKeyPressedInUserInteraction != nil {
		d._ValidationKeyPressedInUserInteraction(interaction)
	}
}

// HasValidationKeyPressedInUserInteraction returns true if a handler for ValidationKeyPressedInUserInteraction has been set.
func (d *TKSmartCardUserInteractionDelegate) HasValidationKeyPressedInUserInteraction() bool {
	return d._ValidationKeyPressedInUserInteraction != nil
}

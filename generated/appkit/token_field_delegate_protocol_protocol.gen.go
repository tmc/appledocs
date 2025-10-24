// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PTokenFieldDelegate is the NSTokenFieldDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTokenFieldDelegate
type PTokenFieldDelegate interface {
	// Optional methods
	TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array
	HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool
	TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) foundation.String
	HasTokenFieldDisplayStringForRepresentedObject() bool
	TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) foundation.String
	HasTokenFieldEditingStringForRepresentedObject() bool
	TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool
	HasTokenFieldHasMenuForRepresentedObject() bool
	TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu
	HasTokenFieldMenuForRepresentedObject() bool
	TokenFieldReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) foundation.Array
	HasTokenFieldReadFromPasteboard() bool
	TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString objc.IObject /* cross-framework: NSString */) objc.ID
	HasTokenFieldRepresentedObjectForEditingString() bool
	TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array
	HasTokenFieldShouldAddObjectsAtIndex() bool
	TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle
	HasTokenFieldStyleForRepresentedObject() bool
	TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool
	HasTokenFieldWriteRepresentedObjectsToPasteboard() bool
}

// TokenFieldDelegate is a delegate implementation builder for the PTokenFieldDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TokenFieldDelegate struct {
	_TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem func(tokenField ITokenField, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array
	_TokenFieldDisplayStringForRepresentedObject func(tokenField ITokenField, representedObject objc.IObject) foundation.String
	_TokenFieldEditingStringForRepresentedObject func(tokenField ITokenField, representedObject objc.IObject) foundation.String
	_TokenFieldHasMenuForRepresentedObject func(tokenField ITokenField, representedObject objc.IObject) bool
	_TokenFieldMenuForRepresentedObject func(tokenField ITokenField, representedObject objc.IObject) Menu
	_TokenFieldReadFromPasteboard func(tokenField ITokenField, pboard IPasteboard) foundation.Array
	_TokenFieldRepresentedObjectForEditingString func(tokenField ITokenField, editingString objc.IObject /* cross-framework: NSString */) objc.ID
	_TokenFieldShouldAddObjectsAtIndex func(tokenField ITokenField, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array
	_TokenFieldStyleForRepresentedObject func(tokenField ITokenField, representedObject objc.IObject) TokenStyle
	_TokenFieldWriteRepresentedObjectsToPasteboard func(tokenField ITokenField, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool
}

// SetTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem sets the handler for the TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem delegate method.
//
// Allows the delegate to provide an array of appropriate completions for the contents of the receiver.
func (d *TokenFieldDelegate) SetTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(f func(tokenField ITokenField, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array) {
	d._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem = f
}

// SetTokenFieldDisplayStringForRepresentedObject sets the handler for the TokenFieldDisplayStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be displayed as a proxy for the given represented object.
func (d *TokenFieldDelegate) SetTokenFieldDisplayStringForRepresentedObject(f func(tokenField ITokenField, representedObject objc.IObject) foundation.String) {
	d._TokenFieldDisplayStringForRepresentedObject = f
}

// SetTokenFieldEditingStringForRepresentedObject sets the handler for the TokenFieldEditingStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be edited as a proxy for a represented object.
func (d *TokenFieldDelegate) SetTokenFieldEditingStringForRepresentedObject(f func(tokenField ITokenField, representedObject objc.IObject) foundation.String) {
	d._TokenFieldEditingStringForRepresentedObject = f
}

// SetTokenFieldHasMenuForRepresentedObject sets the handler for the TokenFieldHasMenuForRepresentedObject delegate method.
//
// Allows the delegate to specify whether the given represented object provides a menu.
func (d *TokenFieldDelegate) SetTokenFieldHasMenuForRepresentedObject(f func(tokenField ITokenField, representedObject objc.IObject) bool) {
	d._TokenFieldHasMenuForRepresentedObject = f
}

// SetTokenFieldMenuForRepresentedObject sets the handler for the TokenFieldMenuForRepresentedObject delegate method.
//
// Allows the delegate to provide a menu for the specified represented object.
func (d *TokenFieldDelegate) SetTokenFieldMenuForRepresentedObject(f func(tokenField ITokenField, representedObject objc.IObject) Menu) {
	d._TokenFieldMenuForRepresentedObject = f
}

// SetTokenFieldReadFromPasteboard sets the handler for the TokenFieldReadFromPasteboard delegate method.
//
// Allows the delegate to return an array of objects representing the data read from the specified pasteboard.
func (d *TokenFieldDelegate) SetTokenFieldReadFromPasteboard(f func(tokenField ITokenField, pboard IPasteboard) foundation.Array) {
	d._TokenFieldReadFromPasteboard = f
}

// SetTokenFieldRepresentedObjectForEditingString sets the handler for the TokenFieldRepresentedObjectForEditingString delegate method.
//
// Allows the delegate to provide a represented object for the given editing string.
func (d *TokenFieldDelegate) SetTokenFieldRepresentedObjectForEditingString(f func(tokenField ITokenField, editingString objc.IObject /* cross-framework: NSString */) objc.ID) {
	d._TokenFieldRepresentedObjectForEditingString = f
}

// SetTokenFieldShouldAddObjectsAtIndex sets the handler for the TokenFieldShouldAddObjectsAtIndex delegate method.
//
// Allows the delegate to validate the tokens to be added to the receiver at a particular location.
func (d *TokenFieldDelegate) SetTokenFieldShouldAddObjectsAtIndex(f func(tokenField ITokenField, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array) {
	d._TokenFieldShouldAddObjectsAtIndex = f
}

// SetTokenFieldStyleForRepresentedObject sets the handler for the TokenFieldStyleForRepresentedObject delegate method.
//
// Allows the delegate to return the token style for editing the specified represented object.
func (d *TokenFieldDelegate) SetTokenFieldStyleForRepresentedObject(f func(tokenField ITokenField, representedObject objc.IObject) TokenStyle) {
	d._TokenFieldStyleForRepresentedObject = f
}

// SetTokenFieldWriteRepresentedObjectsToPasteboard sets the handler for the TokenFieldWriteRepresentedObjectsToPasteboard delegate method.
//
// Sent so the delegate can write represented objects to the pasteboard corresponding to a given array of display strings.
func (d *TokenFieldDelegate) SetTokenFieldWriteRepresentedObjectsToPasteboard(f func(tokenField ITokenField, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool) {
	d._TokenFieldWriteRepresentedObjectsToPasteboard = f
}

// TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array {
	if d._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil {
		return d._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField, substring, tokenIndex, selectedIndex)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem returns true if a handler for TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem has been set.
func (d *TokenFieldDelegate) HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return d._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil
}

// TokenFieldDisplayStringForRepresentedObject implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) foundation.String {
	if d._TokenFieldDisplayStringForRepresentedObject != nil {
		return d._TokenFieldDisplayStringForRepresentedObject(tokenField, representedObject)
	}
	var zero foundation.String
	return zero
}

// HasTokenFieldDisplayStringForRepresentedObject returns true if a handler for TokenFieldDisplayStringForRepresentedObject has been set.
func (d *TokenFieldDelegate) HasTokenFieldDisplayStringForRepresentedObject() bool {
	return d._TokenFieldDisplayStringForRepresentedObject != nil
}

// TokenFieldEditingStringForRepresentedObject implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) foundation.String {
	if d._TokenFieldEditingStringForRepresentedObject != nil {
		return d._TokenFieldEditingStringForRepresentedObject(tokenField, representedObject)
	}
	var zero foundation.String
	return zero
}

// HasTokenFieldEditingStringForRepresentedObject returns true if a handler for TokenFieldEditingStringForRepresentedObject has been set.
func (d *TokenFieldDelegate) HasTokenFieldEditingStringForRepresentedObject() bool {
	return d._TokenFieldEditingStringForRepresentedObject != nil
}

// TokenFieldHasMenuForRepresentedObject implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool {
	if d._TokenFieldHasMenuForRepresentedObject != nil {
		return d._TokenFieldHasMenuForRepresentedObject(tokenField, representedObject)
	}
	var zero bool
	return zero
}

// HasTokenFieldHasMenuForRepresentedObject returns true if a handler for TokenFieldHasMenuForRepresentedObject has been set.
func (d *TokenFieldDelegate) HasTokenFieldHasMenuForRepresentedObject() bool {
	return d._TokenFieldHasMenuForRepresentedObject != nil
}

// TokenFieldMenuForRepresentedObject implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu {
	if d._TokenFieldMenuForRepresentedObject != nil {
		return d._TokenFieldMenuForRepresentedObject(tokenField, representedObject)
	}
	var zero Menu
	return zero
}

// HasTokenFieldMenuForRepresentedObject returns true if a handler for TokenFieldMenuForRepresentedObject has been set.
func (d *TokenFieldDelegate) HasTokenFieldMenuForRepresentedObject() bool {
	return d._TokenFieldMenuForRepresentedObject != nil
}

// TokenFieldReadFromPasteboard implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) foundation.Array {
	if d._TokenFieldReadFromPasteboard != nil {
		return d._TokenFieldReadFromPasteboard(tokenField, pboard)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldReadFromPasteboard returns true if a handler for TokenFieldReadFromPasteboard has been set.
func (d *TokenFieldDelegate) HasTokenFieldReadFromPasteboard() bool {
	return d._TokenFieldReadFromPasteboard != nil
}

// TokenFieldRepresentedObjectForEditingString implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString objc.IObject /* cross-framework: NSString */) objc.ID {
	if d._TokenFieldRepresentedObjectForEditingString != nil {
		return d._TokenFieldRepresentedObjectForEditingString(tokenField, editingString)
	}
	var zero objc.ID
	return zero
}

// HasTokenFieldRepresentedObjectForEditingString returns true if a handler for TokenFieldRepresentedObjectForEditingString has been set.
func (d *TokenFieldDelegate) HasTokenFieldRepresentedObjectForEditingString() bool {
	return d._TokenFieldRepresentedObjectForEditingString != nil
}

// TokenFieldShouldAddObjectsAtIndex implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array {
	if d._TokenFieldShouldAddObjectsAtIndex != nil {
		return d._TokenFieldShouldAddObjectsAtIndex(tokenField, tokens, index)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldShouldAddObjectsAtIndex returns true if a handler for TokenFieldShouldAddObjectsAtIndex has been set.
func (d *TokenFieldDelegate) HasTokenFieldShouldAddObjectsAtIndex() bool {
	return d._TokenFieldShouldAddObjectsAtIndex != nil
}

// TokenFieldStyleForRepresentedObject implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle {
	if d._TokenFieldStyleForRepresentedObject != nil {
		return d._TokenFieldStyleForRepresentedObject(tokenField, representedObject)
	}
	var zero TokenStyle
	return zero
}

// HasTokenFieldStyleForRepresentedObject returns true if a handler for TokenFieldStyleForRepresentedObject has been set.
func (d *TokenFieldDelegate) HasTokenFieldStyleForRepresentedObject() bool {
	return d._TokenFieldStyleForRepresentedObject != nil
}

// TokenFieldWriteRepresentedObjectsToPasteboard implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool {
	if d._TokenFieldWriteRepresentedObjectsToPasteboard != nil {
		return d._TokenFieldWriteRepresentedObjectsToPasteboard(tokenField, objects, pboard)
	}
	var zero bool
	return zero
}

// HasTokenFieldWriteRepresentedObjectsToPasteboard returns true if a handler for TokenFieldWriteRepresentedObjectsToPasteboard has been set.
func (d *TokenFieldDelegate) HasTokenFieldWriteRepresentedObjectsToPasteboard() bool {
	return d._TokenFieldWriteRepresentedObjectsToPasteboard != nil
}

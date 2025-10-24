// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PTokenFieldCellDelegate is the NSTokenFieldCellDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects to work with tokenized strings.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTokenFieldCellDelegate
type PTokenFieldCellDelegate interface {
	// Optional methods
	TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array
	HasTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool
	TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String
	HasTokenFieldCellDisplayStringForRepresentedObject() bool
	TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String
	HasTokenFieldCellEditingStringForRepresentedObject() bool
	TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool
	HasTokenFieldCellHasMenuForRepresentedObject() bool
	TokenFieldCellMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu
	HasTokenFieldCellMenuForRepresentedObject() bool
	TokenFieldCellReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) foundation.Array
	HasTokenFieldCellReadFromPasteboard() bool
	TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString objc.IObject /* cross-framework: NSString */) objc.ID
	HasTokenFieldCellRepresentedObjectForEditingString() bool
	TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell ITokenFieldCell, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array
	HasTokenFieldCellShouldAddObjectsAtIndex() bool
	TokenFieldCellStyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle
	HasTokenFieldCellStyleForRepresentedObject() bool
	TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell ITokenFieldCell, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool
	HasTokenFieldCellWriteRepresentedObjectsToPasteboard() bool
}

// TokenFieldCellDelegate is a delegate implementation builder for the PTokenFieldCellDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TokenFieldCellDelegate struct {
	_TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem func(tokenFieldCell ITokenFieldCell, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array
	_TokenFieldCellDisplayStringForRepresentedObject func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String
	_TokenFieldCellEditingStringForRepresentedObject func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String
	_TokenFieldCellHasMenuForRepresentedObject func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool
	_TokenFieldCellMenuForRepresentedObject func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu
	_TokenFieldCellReadFromPasteboard func(tokenFieldCell ITokenFieldCell, pboard IPasteboard) foundation.Array
	_TokenFieldCellRepresentedObjectForEditingString func(tokenFieldCell ITokenFieldCell, editingString objc.IObject /* cross-framework: NSString */) objc.ID
	_TokenFieldCellShouldAddObjectsAtIndex func(tokenFieldCell ITokenFieldCell, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array
	_TokenFieldCellStyleForRepresentedObject func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle
	_TokenFieldCellWriteRepresentedObjectsToPasteboard func(tokenFieldCell ITokenFieldCell, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool
}

// SetTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem sets the handler for the TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem delegate method.
//
// Allows the delegate to provide an array of appropriate completions for the contents of the receiver.
func (d *TokenFieldCellDelegate) SetTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(f func(tokenFieldCell ITokenFieldCell, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array) {
	d._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem = f
}

// SetTokenFieldCellDisplayStringForRepresentedObject sets the handler for the TokenFieldCellDisplayStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be displayed as a proxy for the represented object.
func (d *TokenFieldCellDelegate) SetTokenFieldCellDisplayStringForRepresentedObject(f func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String) {
	d._TokenFieldCellDisplayStringForRepresentedObject = f
}

// SetTokenFieldCellEditingStringForRepresentedObject sets the handler for the TokenFieldCellEditingStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be edited as a proxy for the represented object.
func (d *TokenFieldCellDelegate) SetTokenFieldCellEditingStringForRepresentedObject(f func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String) {
	d._TokenFieldCellEditingStringForRepresentedObject = f
}

// SetTokenFieldCellHasMenuForRepresentedObject sets the handler for the TokenFieldCellHasMenuForRepresentedObject delegate method.
//
// Allows the delegate to specify whether the represented object provides a menu.
func (d *TokenFieldCellDelegate) SetTokenFieldCellHasMenuForRepresentedObject(f func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool) {
	d._TokenFieldCellHasMenuForRepresentedObject = f
}

// SetTokenFieldCellMenuForRepresentedObject sets the handler for the TokenFieldCellMenuForRepresentedObject delegate method.
//
// Allows the delegate to provide a menu for the specified represented object.
func (d *TokenFieldCellDelegate) SetTokenFieldCellMenuForRepresentedObject(f func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu) {
	d._TokenFieldCellMenuForRepresentedObject = f
}

// SetTokenFieldCellReadFromPasteboard sets the handler for the TokenFieldCellReadFromPasteboard delegate method.
//
// Allows the delegate to return an array of objects representing the data read from  .
func (d *TokenFieldCellDelegate) SetTokenFieldCellReadFromPasteboard(f func(tokenFieldCell ITokenFieldCell, pboard IPasteboard) foundation.Array) {
	d._TokenFieldCellReadFromPasteboard = f
}

// SetTokenFieldCellRepresentedObjectForEditingString sets the handler for the TokenFieldCellRepresentedObjectForEditingString delegate method.
//
// Allows the delegate to provide a represented object for the string being edited.
func (d *TokenFieldCellDelegate) SetTokenFieldCellRepresentedObjectForEditingString(f func(tokenFieldCell ITokenFieldCell, editingString objc.IObject /* cross-framework: NSString */) objc.ID) {
	d._TokenFieldCellRepresentedObjectForEditingString = f
}

// SetTokenFieldCellShouldAddObjectsAtIndex sets the handler for the TokenFieldCellShouldAddObjectsAtIndex delegate method.
//
// Allows the delegate to validate the tokens to be added to the receiver at a given index.
func (d *TokenFieldCellDelegate) SetTokenFieldCellShouldAddObjectsAtIndex(f func(tokenFieldCell ITokenFieldCell, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array) {
	d._TokenFieldCellShouldAddObjectsAtIndex = f
}

// SetTokenFieldCellStyleForRepresentedObject sets the handler for the TokenFieldCellStyleForRepresentedObject delegate method.
//
// Allows the delegate to return the token style for editing the specified represented object.
func (d *TokenFieldCellDelegate) SetTokenFieldCellStyleForRepresentedObject(f func(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle) {
	d._TokenFieldCellStyleForRepresentedObject = f
}

// SetTokenFieldCellWriteRepresentedObjectsToPasteboard sets the handler for the TokenFieldCellWriteRepresentedObjectsToPasteboard delegate method.
//
// Allows the delegate the opportunity to write custom pasteboard types to the pasteboard for the represented objects in  .
func (d *TokenFieldCellDelegate) SetTokenFieldCellWriteRepresentedObjectsToPasteboard(f func(tokenFieldCell ITokenFieldCell, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool) {
	d._TokenFieldCellWriteRepresentedObjectsToPasteboard = f
}

// TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring objc.IObject /* cross-framework: NSString */, tokenIndex int, selectedIndex int) foundation.Array {
	if d._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil {
		return d._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell, substring, tokenIndex, selectedIndex)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem returns true if a handler for TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return d._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil
}

// TokenFieldCellDisplayStringForRepresentedObject implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String {
	if d._TokenFieldCellDisplayStringForRepresentedObject != nil {
		return d._TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell, representedObject)
	}
	var zero foundation.String
	return zero
}

// HasTokenFieldCellDisplayStringForRepresentedObject returns true if a handler for TokenFieldCellDisplayStringForRepresentedObject has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellDisplayStringForRepresentedObject() bool {
	return d._TokenFieldCellDisplayStringForRepresentedObject != nil
}

// TokenFieldCellEditingStringForRepresentedObject implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) foundation.String {
	if d._TokenFieldCellEditingStringForRepresentedObject != nil {
		return d._TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell, representedObject)
	}
	var zero foundation.String
	return zero
}

// HasTokenFieldCellEditingStringForRepresentedObject returns true if a handler for TokenFieldCellEditingStringForRepresentedObject has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellEditingStringForRepresentedObject() bool {
	return d._TokenFieldCellEditingStringForRepresentedObject != nil
}

// TokenFieldCellHasMenuForRepresentedObject implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool {
	if d._TokenFieldCellHasMenuForRepresentedObject != nil {
		return d._TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell, representedObject)
	}
	var zero bool
	return zero
}

// HasTokenFieldCellHasMenuForRepresentedObject returns true if a handler for TokenFieldCellHasMenuForRepresentedObject has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellHasMenuForRepresentedObject() bool {
	return d._TokenFieldCellHasMenuForRepresentedObject != nil
}

// TokenFieldCellMenuForRepresentedObject implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu {
	if d._TokenFieldCellMenuForRepresentedObject != nil {
		return d._TokenFieldCellMenuForRepresentedObject(tokenFieldCell, representedObject)
	}
	var zero Menu
	return zero
}

// HasTokenFieldCellMenuForRepresentedObject returns true if a handler for TokenFieldCellMenuForRepresentedObject has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellMenuForRepresentedObject() bool {
	return d._TokenFieldCellMenuForRepresentedObject != nil
}

// TokenFieldCellReadFromPasteboard implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) foundation.Array {
	if d._TokenFieldCellReadFromPasteboard != nil {
		return d._TokenFieldCellReadFromPasteboard(tokenFieldCell, pboard)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldCellReadFromPasteboard returns true if a handler for TokenFieldCellReadFromPasteboard has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellReadFromPasteboard() bool {
	return d._TokenFieldCellReadFromPasteboard != nil
}

// TokenFieldCellRepresentedObjectForEditingString implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString objc.IObject /* cross-framework: NSString */) objc.ID {
	if d._TokenFieldCellRepresentedObjectForEditingString != nil {
		return d._TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell, editingString)
	}
	var zero objc.ID
	return zero
}

// HasTokenFieldCellRepresentedObjectForEditingString returns true if a handler for TokenFieldCellRepresentedObjectForEditingString has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellRepresentedObjectForEditingString() bool {
	return d._TokenFieldCellRepresentedObjectForEditingString != nil
}

// TokenFieldCellShouldAddObjectsAtIndex implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell ITokenFieldCell, tokens objc.IObject /* cross-framework: NSArray */, index uint) foundation.Array {
	if d._TokenFieldCellShouldAddObjectsAtIndex != nil {
		return d._TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell, tokens, index)
	}
	var zero foundation.Array
	return zero
}

// HasTokenFieldCellShouldAddObjectsAtIndex returns true if a handler for TokenFieldCellShouldAddObjectsAtIndex has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellShouldAddObjectsAtIndex() bool {
	return d._TokenFieldCellShouldAddObjectsAtIndex != nil
}

// TokenFieldCellStyleForRepresentedObject implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellStyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle {
	if d._TokenFieldCellStyleForRepresentedObject != nil {
		return d._TokenFieldCellStyleForRepresentedObject(tokenFieldCell, representedObject)
	}
	var zero TokenStyle
	return zero
}

// HasTokenFieldCellStyleForRepresentedObject returns true if a handler for TokenFieldCellStyleForRepresentedObject has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellStyleForRepresentedObject() bool {
	return d._TokenFieldCellStyleForRepresentedObject != nil
}

// TokenFieldCellWriteRepresentedObjectsToPasteboard implements the PTokenFieldCellDelegate interface.
func (d *TokenFieldCellDelegate) TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell ITokenFieldCell, objects objc.IObject /* cross-framework: NSArray */, pboard IPasteboard) bool {
	if d._TokenFieldCellWriteRepresentedObjectsToPasteboard != nil {
		return d._TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell, objects, pboard)
	}
	var zero bool
	return zero
}

// HasTokenFieldCellWriteRepresentedObjectsToPasteboard returns true if a handler for TokenFieldCellWriteRepresentedObjectsToPasteboard has been set.
func (d *TokenFieldCellDelegate) HasTokenFieldCellWriteRepresentedObjectsToPasteboard() bool {
	return d._TokenFieldCellWriteRepresentedObjectsToPasteboard != nil
}

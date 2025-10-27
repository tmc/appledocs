// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring foundation.foundation.INSString, tokenIndex int, selectedIndex int) foundation.Array
	HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool
	TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String
	HasTokenFieldDisplayStringForRepresentedObject() bool
	TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String
	HasTokenFieldEditingStringForRepresentedObject() bool
	TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) bool
	HasTokenFieldHasMenuForRepresentedObject() bool
	TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) IMenu
	HasTokenFieldMenuForRepresentedObject() bool
	TokenFieldReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) foundation.Array
	HasTokenFieldReadFromPasteboard() bool
	TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString foundation.foundation.INSString) objc.ID
	HasTokenFieldRepresentedObjectForEditingString() bool
	TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens foundation.foundation.INSArray, index uint) foundation.Array
	HasTokenFieldShouldAddObjectsAtIndex() bool
	TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) TokenStyle
	HasTokenFieldStyleForRepresentedObject() bool
	TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects foundation.foundation.INSArray, pboard IPasteboard) bool
	HasTokenFieldWriteRepresentedObjectsToPasteboard() bool
}

// TokenFieldDelegate is a delegate implementation builder for the PTokenFieldDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TokenFieldDelegate struct {
	_TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem func(tokenField ITokenField, substring foundation.foundation.INSString, tokenIndex int, selectedIndex int) foundation.Array
	_TokenFieldDisplayStringForRepresentedObject func(tokenField ITokenField, representedObject objectivec.IObject) foundation.String
	_TokenFieldEditingStringForRepresentedObject func(tokenField ITokenField, representedObject objectivec.IObject) foundation.String
	_TokenFieldHasMenuForRepresentedObject func(tokenField ITokenField, representedObject objectivec.IObject) bool
	_TokenFieldMenuForRepresentedObject func(tokenField ITokenField, representedObject objectivec.IObject) IMenu
	_TokenFieldReadFromPasteboard func(tokenField ITokenField, pboard IPasteboard) foundation.Array
	_TokenFieldRepresentedObjectForEditingString func(tokenField ITokenField, editingString foundation.foundation.INSString) objc.ID
	_TokenFieldShouldAddObjectsAtIndex func(tokenField ITokenField, tokens foundation.foundation.INSArray, index uint) foundation.Array
	_TokenFieldStyleForRepresentedObject func(tokenField ITokenField, representedObject objectivec.IObject) TokenStyle
	_TokenFieldWriteRepresentedObjectsToPasteboard func(tokenField ITokenField, objects foundation.foundation.INSArray, pboard IPasteboard) bool
}

// SetTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem sets the handler for the TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem delegate method.
//
// Allows the delegate to provide an array of appropriate completions for the contents of the receiver.
func (d *TokenFieldDelegate) SetTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(f func(tokenField ITokenField, substring foundation.foundation.INSString, tokenIndex int, selectedIndex int) foundation.Array) {
	d._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem = f
}

// SetTokenFieldDisplayStringForRepresentedObject sets the handler for the TokenFieldDisplayStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be displayed as a proxy for the given represented object.
func (d *TokenFieldDelegate) SetTokenFieldDisplayStringForRepresentedObject(f func(tokenField ITokenField, representedObject objectivec.IObject) foundation.String) {
	d._TokenFieldDisplayStringForRepresentedObject = f
}

// SetTokenFieldEditingStringForRepresentedObject sets the handler for the TokenFieldEditingStringForRepresentedObject delegate method.
//
// Allows the delegate to provide a string to be edited as a proxy for a represented object.
func (d *TokenFieldDelegate) SetTokenFieldEditingStringForRepresentedObject(f func(tokenField ITokenField, representedObject objectivec.IObject) foundation.String) {
	d._TokenFieldEditingStringForRepresentedObject = f
}

// SetTokenFieldHasMenuForRepresentedObject sets the handler for the TokenFieldHasMenuForRepresentedObject delegate method.
//
// Allows the delegate to specify whether the given represented object provides a menu.
func (d *TokenFieldDelegate) SetTokenFieldHasMenuForRepresentedObject(f func(tokenField ITokenField, representedObject objectivec.IObject) bool) {
	d._TokenFieldHasMenuForRepresentedObject = f
}

// SetTokenFieldMenuForRepresentedObject sets the handler for the TokenFieldMenuForRepresentedObject delegate method.
//
// Allows the delegate to provide a menu for the specified represented object.
func (d *TokenFieldDelegate) SetTokenFieldMenuForRepresentedObject(f func(tokenField ITokenField, representedObject objectivec.IObject) IMenu) {
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
func (d *TokenFieldDelegate) SetTokenFieldRepresentedObjectForEditingString(f func(tokenField ITokenField, editingString foundation.foundation.INSString) objc.ID) {
	d._TokenFieldRepresentedObjectForEditingString = f
}

// SetTokenFieldShouldAddObjectsAtIndex sets the handler for the TokenFieldShouldAddObjectsAtIndex delegate method.
//
// Allows the delegate to validate the tokens to be added to the receiver at a particular location.
func (d *TokenFieldDelegate) SetTokenFieldShouldAddObjectsAtIndex(f func(tokenField ITokenField, tokens foundation.foundation.INSArray, index uint) foundation.Array) {
	d._TokenFieldShouldAddObjectsAtIndex = f
}

// SetTokenFieldStyleForRepresentedObject sets the handler for the TokenFieldStyleForRepresentedObject delegate method.
//
// Allows the delegate to return the token style for editing the specified represented object.
func (d *TokenFieldDelegate) SetTokenFieldStyleForRepresentedObject(f func(tokenField ITokenField, representedObject objectivec.IObject) TokenStyle) {
	d._TokenFieldStyleForRepresentedObject = f
}

// SetTokenFieldWriteRepresentedObjectsToPasteboard sets the handler for the TokenFieldWriteRepresentedObjectsToPasteboard delegate method.
//
// Sent so the delegate can write represented objects to the pasteboard corresponding to a given array of display strings.
func (d *TokenFieldDelegate) SetTokenFieldWriteRepresentedObjectsToPasteboard(f func(tokenField ITokenField, objects foundation.foundation.INSArray, pboard IPasteboard) bool) {
	d._TokenFieldWriteRepresentedObjectsToPasteboard = f
}

// TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem implements the PTokenFieldDelegate interface.
func (d *TokenFieldDelegate) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring foundation.foundation.INSString, tokenIndex int, selectedIndex int) foundation.Array {
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
func (d *TokenFieldDelegate) TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String {
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
func (d *TokenFieldDelegate) TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String {
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
func (d *TokenFieldDelegate) TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) bool {
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
func (d *TokenFieldDelegate) TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) IMenu {
	if d._TokenFieldMenuForRepresentedObject != nil {
		return d._TokenFieldMenuForRepresentedObject(tokenField, representedObject)
	}
	var zero IMenu
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
func (d *TokenFieldDelegate) TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString foundation.foundation.INSString) objc.ID {
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
func (d *TokenFieldDelegate) TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens foundation.foundation.INSArray, index uint) foundation.Array {
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
func (d *TokenFieldDelegate) TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) TokenStyle {
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
func (d *TokenFieldDelegate) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects foundation.foundation.INSArray, pboard IPasteboard) bool {
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

// TokenFieldDelegateObject wraps an existing Objective-C object that conforms to the PTokenFieldDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TokenFieldDelegateObject struct {
	objectivec.Object
}

// NewTokenFieldDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTokenFieldDelegate protocol.
func NewTokenFieldDelegateObject(obj objectivec.Object) *TokenFieldDelegateObject {
	return &TokenFieldDelegateObject{obj}
}

// Make sure TokenFieldDelegateObject implements PTokenFieldDelegate.
var _ PTokenFieldDelegate = (*TokenFieldDelegateObject)(nil)

// TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring foundation.foundation.INSString, tokenIndex int, selectedIndex int) foundation.Array {
	return objc.Send[foundation.Array](o.ID, objc.Sel("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), tokenField, substring, tokenIndex, selectedIndex)
}

// HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldDisplayStringForRepresentedObject implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("tokenField:displayStringForRepresentedObject:"), tokenField, representedObject)
}

// HasTokenFieldDisplayStringForRepresentedObject returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldDisplayStringForRepresentedObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldEditingStringForRepresentedObject implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("tokenField:editingStringForRepresentedObject:"), tokenField, representedObject)
}

// HasTokenFieldEditingStringForRepresentedObject returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldEditingStringForRepresentedObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldHasMenuForRepresentedObject implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("tokenField:hasMenuForRepresentedObject:"), tokenField, representedObject)
}

// HasTokenFieldHasMenuForRepresentedObject returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldHasMenuForRepresentedObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldMenuForRepresentedObject implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) IMenu {
	return objc.Send[IMenu](o.ID, objc.Sel("tokenField:menuForRepresentedObject:"), tokenField, representedObject)
}

// HasTokenFieldMenuForRepresentedObject returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldMenuForRepresentedObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldReadFromPasteboard implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) foundation.Array {
	return objc.Send[foundation.Array](o.ID, objc.Sel("tokenField:readFromPasteboard:"), tokenField, pboard)
}

// HasTokenFieldReadFromPasteboard returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldReadFromPasteboard() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldRepresentedObjectForEditingString implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString foundation.foundation.INSString) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("tokenField:representedObjectForEditingString:"), tokenField, editingString)
}

// HasTokenFieldRepresentedObjectForEditingString returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldRepresentedObjectForEditingString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldShouldAddObjectsAtIndex implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens foundation.foundation.INSArray, index uint) foundation.Array {
	return objc.Send[foundation.Array](o.ID, objc.Sel("tokenField:shouldAddObjects:atIndex:"), tokenField, tokens, index)
}

// HasTokenFieldShouldAddObjectsAtIndex returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldShouldAddObjectsAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldStyleForRepresentedObject implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objectivec.IObject) TokenStyle {
	return objc.Send[TokenStyle](o.ID, objc.Sel("tokenField:styleForRepresentedObject:"), tokenField, representedObject)
}

// HasTokenFieldStyleForRepresentedObject returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldStyleForRepresentedObject() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TokenFieldWriteRepresentedObjectsToPasteboard implements the PTokenFieldDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TokenFieldDelegateObject) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects foundation.foundation.INSArray, pboard IPasteboard) bool {
	return objc.Send[bool](o.ID, objc.Sel("tokenField:writeRepresentedObjects:toPasteboard:"), tokenField, objects, pboard)
}

// HasTokenFieldWriteRepresentedObjectsToPasteboard returns true; this is a placeholder for optional method checks.
func (o *TokenFieldDelegateObject) HasTokenFieldWriteRepresentedObjectsToPasteboard() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

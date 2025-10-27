// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PComboBoxDataSource is the NSComboBoxDataSource protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSComboBoxDataSource
type PComboBoxDataSource interface {
	// Optional methods
	ComboBoxCompletedString(comboBox IComboBox, string_ foundation.foundation.INSString) foundation.String
	HasComboBoxCompletedString() bool
	ComboBoxIndexOfItemWithStringValue(comboBox IComboBox, string_ foundation.foundation.INSString) uint
	HasComboBoxIndexOfItemWithStringValue() bool
	ComboBoxObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.ID
	HasComboBoxObjectValueForItemAtIndex() bool
	NumberOfItemsInComboBox(comboBox IComboBox) int
	HasNumberOfItemsInComboBox() bool
}

// ComboBoxDataSource is a delegate implementation builder for the PComboBoxDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ComboBoxDataSource struct {
	_ComboBoxCompletedString func(comboBox IComboBox, string_ foundation.foundation.INSString) foundation.String
	_ComboBoxIndexOfItemWithStringValue func(comboBox IComboBox, string_ foundation.foundation.INSString) uint
	_ComboBoxObjectValueForItemAtIndex func(comboBox IComboBox, index int) objc.ID
	_NumberOfItemsInComboBox func(comboBox IComboBox) int
}

// SetComboBoxCompletedString sets the handler for the ComboBoxCompletedString delegate method.
//
// Returns the first item from the pop-up list that starts with the text the user has typed.
func (d *ComboBoxDataSource) SetComboBoxCompletedString(f func(comboBox IComboBox, string_ foundation.foundation.INSString) foundation.String) {
	d._ComboBoxCompletedString = f
}

// SetComboBoxIndexOfItemWithStringValue sets the handler for the ComboBoxIndexOfItemWithStringValue delegate method.
//
// Returns the index of the combo box item matching the specified string.
func (d *ComboBoxDataSource) SetComboBoxIndexOfItemWithStringValue(f func(comboBox IComboBox, string_ foundation.foundation.INSString) uint) {
	d._ComboBoxIndexOfItemWithStringValue = f
}

// SetComboBoxObjectValueForItemAtIndex sets the handler for the ComboBoxObjectValueForItemAtIndex delegate method.
//
// Returns the object that corresponds to the item at the specified index in the combo box.
func (d *ComboBoxDataSource) SetComboBoxObjectValueForItemAtIndex(f func(comboBox IComboBox, index int) objc.ID) {
	d._ComboBoxObjectValueForItemAtIndex = f
}

// SetNumberOfItemsInComboBox sets the handler for the NumberOfItemsInComboBox delegate method.
//
// Returns the number of items that the data source manages for the combo box.
func (d *ComboBoxDataSource) SetNumberOfItemsInComboBox(f func(comboBox IComboBox) int) {
	d._NumberOfItemsInComboBox = f
}

// ComboBoxCompletedString implements the PComboBoxDataSource interface.
func (d *ComboBoxDataSource) ComboBoxCompletedString(comboBox IComboBox, string_ foundation.foundation.INSString) foundation.String {
	if d._ComboBoxCompletedString != nil {
		return d._ComboBoxCompletedString(comboBox, string_)
	}
	var zero foundation.String
	return zero
}

// HasComboBoxCompletedString returns true if a handler for ComboBoxCompletedString has been set.
func (d *ComboBoxDataSource) HasComboBoxCompletedString() bool {
	return d._ComboBoxCompletedString != nil
}

// ComboBoxIndexOfItemWithStringValue implements the PComboBoxDataSource interface.
func (d *ComboBoxDataSource) ComboBoxIndexOfItemWithStringValue(comboBox IComboBox, string_ foundation.foundation.INSString) uint {
	if d._ComboBoxIndexOfItemWithStringValue != nil {
		return d._ComboBoxIndexOfItemWithStringValue(comboBox, string_)
	}
	var zero uint
	return zero
}

// HasComboBoxIndexOfItemWithStringValue returns true if a handler for ComboBoxIndexOfItemWithStringValue has been set.
func (d *ComboBoxDataSource) HasComboBoxIndexOfItemWithStringValue() bool {
	return d._ComboBoxIndexOfItemWithStringValue != nil
}

// ComboBoxObjectValueForItemAtIndex implements the PComboBoxDataSource interface.
func (d *ComboBoxDataSource) ComboBoxObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.ID {
	if d._ComboBoxObjectValueForItemAtIndex != nil {
		return d._ComboBoxObjectValueForItemAtIndex(comboBox, index)
	}
	var zero objc.ID
	return zero
}

// HasComboBoxObjectValueForItemAtIndex returns true if a handler for ComboBoxObjectValueForItemAtIndex has been set.
func (d *ComboBoxDataSource) HasComboBoxObjectValueForItemAtIndex() bool {
	return d._ComboBoxObjectValueForItemAtIndex != nil
}

// NumberOfItemsInComboBox implements the PComboBoxDataSource interface.
func (d *ComboBoxDataSource) NumberOfItemsInComboBox(comboBox IComboBox) int {
	if d._NumberOfItemsInComboBox != nil {
		return d._NumberOfItemsInComboBox(comboBox)
	}
	var zero int
	return zero
}

// HasNumberOfItemsInComboBox returns true if a handler for NumberOfItemsInComboBox has been set.
func (d *ComboBoxDataSource) HasNumberOfItemsInComboBox() bool {
	return d._NumberOfItemsInComboBox != nil
}

// ComboBoxDataSourceObject wraps an existing Objective-C object that conforms to the PComboBoxDataSource protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type ComboBoxDataSourceObject struct {
	objectivec.Object
}

// NewComboBoxDataSourceObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSComboBoxDataSource protocol.
func NewComboBoxDataSourceObject(obj objectivec.Object) *ComboBoxDataSourceObject {
	return &ComboBoxDataSourceObject{obj}
}

// Make sure ComboBoxDataSourceObject implements PComboBoxDataSource.
var _ PComboBoxDataSource = (*ComboBoxDataSourceObject)(nil)

// ComboBoxCompletedString implements the PComboBoxDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDataSourceObject) ComboBoxCompletedString(comboBox IComboBox, string_ foundation.foundation.INSString) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("comboBox:completedString:"), comboBox, string_)
}

// HasComboBoxCompletedString returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDataSourceObject) HasComboBoxCompletedString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ComboBoxIndexOfItemWithStringValue implements the PComboBoxDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDataSourceObject) ComboBoxIndexOfItemWithStringValue(comboBox IComboBox, string_ foundation.foundation.INSString) uint {
	return objc.Send[uint](o.ID, objc.Sel("comboBox:indexOfItemWithStringValue:"), comboBox, string_)
}

// HasComboBoxIndexOfItemWithStringValue returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDataSourceObject) HasComboBoxIndexOfItemWithStringValue() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ComboBoxObjectValueForItemAtIndex implements the PComboBoxDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDataSourceObject) ComboBoxObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.ID {
	return objc.Send[objc.ID](o.ID, objc.Sel("comboBox:objectValueForItemAtIndex:"), comboBox, index)
}

// HasComboBoxObjectValueForItemAtIndex returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDataSourceObject) HasComboBoxObjectValueForItemAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NumberOfItemsInComboBox implements the PComboBoxDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDataSourceObject) NumberOfItemsInComboBox(comboBox IComboBox) int {
	return objc.Send[int](o.ID, objc.Sel("numberOfItemsInComboBox:"), comboBox)
}

// HasNumberOfItemsInComboBox returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDataSourceObject) HasNumberOfItemsInComboBox() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

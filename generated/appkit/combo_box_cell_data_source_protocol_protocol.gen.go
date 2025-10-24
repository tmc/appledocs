// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PComboBoxCellDataSource is the NSComboBoxCellDataSource protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSComboBoxCellDataSource
type PComboBoxCellDataSource interface {
	// Optional methods
	ComboBoxCellCompletedString(comboBoxCell IComboBoxCell, uncompletedString objc.IObject /* cross-framework: NSString */) foundation.String
	HasComboBoxCellCompletedString() bool
	ComboBoxCellIndexOfItemWithStringValue(comboBoxCell IComboBoxCell, string_ objc.IObject /* cross-framework: NSString */) uint
	HasComboBoxCellIndexOfItemWithStringValue() bool
	ComboBoxCellObjectValueForItemAtIndex(comboBoxCell IComboBoxCell, index int) objc.ID
	HasComboBoxCellObjectValueForItemAtIndex() bool
	NumberOfItemsInComboBoxCell(comboBoxCell IComboBoxCell) int
	HasNumberOfItemsInComboBoxCell() bool
}

// ComboBoxCellDataSource is a delegate implementation builder for the PComboBoxCellDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ComboBoxCellDataSource struct {
	_ComboBoxCellCompletedString func(comboBoxCell IComboBoxCell, uncompletedString objc.IObject /* cross-framework: NSString */) foundation.String
	_ComboBoxCellIndexOfItemWithStringValue func(comboBoxCell IComboBoxCell, string_ objc.IObject /* cross-framework: NSString */) uint
	_ComboBoxCellObjectValueForItemAtIndex func(comboBoxCell IComboBoxCell, index int) objc.ID
	_NumberOfItemsInComboBoxCell func(comboBoxCell IComboBoxCell) int
}

// SetComboBoxCellCompletedString sets the handler for the ComboBoxCellCompletedString delegate method.
//
// Returns the item from the combo box’s pop-up list that matches the text entered by the user.
func (d *ComboBoxCellDataSource) SetComboBoxCellCompletedString(f func(comboBoxCell IComboBoxCell, uncompletedString objc.IObject /* cross-framework: NSString */) foundation.String) {
	d._ComboBoxCellCompletedString = f
}

// SetComboBoxCellIndexOfItemWithStringValue sets the handler for the ComboBoxCellIndexOfItemWithStringValue delegate method.
//
// Invoked by an   object to synchronize the pop-up list’s selected item with the text field’s contents.
func (d *ComboBoxCellDataSource) SetComboBoxCellIndexOfItemWithStringValue(f func(comboBoxCell IComboBoxCell, string_ objc.IObject /* cross-framework: NSString */) uint) {
	d._ComboBoxCellIndexOfItemWithStringValue = f
}

// SetComboBoxCellObjectValueForItemAtIndex sets the handler for the ComboBoxCellObjectValueForItemAtIndex delegate method.
//
// Returns the object that corresponds to the item at the given index in the combo box cell.
func (d *ComboBoxCellDataSource) SetComboBoxCellObjectValueForItemAtIndex(f func(comboBoxCell IComboBoxCell, index int) objc.ID) {
	d._ComboBoxCellObjectValueForItemAtIndex = f
}

// SetNumberOfItemsInComboBoxCell sets the handler for the NumberOfItemsInComboBoxCell delegate method.
//
// Returns the number of items managed for the combo box cell by your data source object.
func (d *ComboBoxCellDataSource) SetNumberOfItemsInComboBoxCell(f func(comboBoxCell IComboBoxCell) int) {
	d._NumberOfItemsInComboBoxCell = f
}

// ComboBoxCellCompletedString implements the PComboBoxCellDataSource interface.
func (d *ComboBoxCellDataSource) ComboBoxCellCompletedString(comboBoxCell IComboBoxCell, uncompletedString objc.IObject /* cross-framework: NSString */) foundation.String {
	if d._ComboBoxCellCompletedString != nil {
		return d._ComboBoxCellCompletedString(comboBoxCell, uncompletedString)
	}
	var zero foundation.String
	return zero
}

// HasComboBoxCellCompletedString returns true if a handler for ComboBoxCellCompletedString has been set.
func (d *ComboBoxCellDataSource) HasComboBoxCellCompletedString() bool {
	return d._ComboBoxCellCompletedString != nil
}

// ComboBoxCellIndexOfItemWithStringValue implements the PComboBoxCellDataSource interface.
func (d *ComboBoxCellDataSource) ComboBoxCellIndexOfItemWithStringValue(comboBoxCell IComboBoxCell, string_ objc.IObject /* cross-framework: NSString */) uint {
	if d._ComboBoxCellIndexOfItemWithStringValue != nil {
		return d._ComboBoxCellIndexOfItemWithStringValue(comboBoxCell, string_)
	}
	var zero uint
	return zero
}

// HasComboBoxCellIndexOfItemWithStringValue returns true if a handler for ComboBoxCellIndexOfItemWithStringValue has been set.
func (d *ComboBoxCellDataSource) HasComboBoxCellIndexOfItemWithStringValue() bool {
	return d._ComboBoxCellIndexOfItemWithStringValue != nil
}

// ComboBoxCellObjectValueForItemAtIndex implements the PComboBoxCellDataSource interface.
func (d *ComboBoxCellDataSource) ComboBoxCellObjectValueForItemAtIndex(comboBoxCell IComboBoxCell, index int) objc.ID {
	if d._ComboBoxCellObjectValueForItemAtIndex != nil {
		return d._ComboBoxCellObjectValueForItemAtIndex(comboBoxCell, index)
	}
	var zero objc.ID
	return zero
}

// HasComboBoxCellObjectValueForItemAtIndex returns true if a handler for ComboBoxCellObjectValueForItemAtIndex has been set.
func (d *ComboBoxCellDataSource) HasComboBoxCellObjectValueForItemAtIndex() bool {
	return d._ComboBoxCellObjectValueForItemAtIndex != nil
}

// NumberOfItemsInComboBoxCell implements the PComboBoxCellDataSource interface.
func (d *ComboBoxCellDataSource) NumberOfItemsInComboBoxCell(comboBoxCell IComboBoxCell) int {
	if d._NumberOfItemsInComboBoxCell != nil {
		return d._NumberOfItemsInComboBoxCell(comboBoxCell)
	}
	var zero int
	return zero
}

// HasNumberOfItemsInComboBoxCell returns true if a handler for NumberOfItemsInComboBoxCell has been set.
func (d *ComboBoxCellDataSource) HasNumberOfItemsInComboBoxCell() bool {
	return d._NumberOfItemsInComboBoxCell != nil
}

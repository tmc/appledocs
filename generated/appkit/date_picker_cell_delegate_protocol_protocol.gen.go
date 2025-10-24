// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PDatePickerCellDelegate is the NSDatePickerCellDelegate protocol interface.
//
// A set of optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDatePickerCellDelegate
type PDatePickerCellDelegate interface {
	// Optional methods
	DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue objc.IObject /* cross-framework: NSDate */, proposedTimeInterval float64)
	HasDatePickerCellValidateProposedDateValueTimeInterval() bool
}

// DatePickerCellDelegate is a delegate implementation builder for the PDatePickerCellDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DatePickerCellDelegate struct {
	_DatePickerCellValidateProposedDateValueTimeInterval func(datePickerCell IDatePickerCell, proposedDateValue objc.IObject /* cross-framework: NSDate */, proposedTimeInterval float64)
}

// SetDatePickerCellValidateProposedDateValueTimeInterval sets the handler for the DatePickerCellValidateProposedDateValueTimeInterval delegate method.
//
// The delegate receives this message each time the user attempts to change the receiver’s value, allowing the delegate the opportunity to override the change.
func (d *DatePickerCellDelegate) SetDatePickerCellValidateProposedDateValueTimeInterval(f func(datePickerCell IDatePickerCell, proposedDateValue objc.IObject /* cross-framework: NSDate */, proposedTimeInterval float64)) {
	d._DatePickerCellValidateProposedDateValueTimeInterval = f
}

// DatePickerCellValidateProposedDateValueTimeInterval implements the PDatePickerCellDelegate interface.
func (d *DatePickerCellDelegate) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue objc.IObject /* cross-framework: NSDate */, proposedTimeInterval float64) {
	if d._DatePickerCellValidateProposedDateValueTimeInterval != nil {
		d._DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell, proposedDateValue, proposedTimeInterval)
	}
}

// HasDatePickerCellValidateProposedDateValueTimeInterval returns true if a handler for DatePickerCellValidateProposedDateValueTimeInterval has been set.
func (d *DatePickerCellDelegate) HasDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return d._DatePickerCellValidateProposedDateValueTimeInterval != nil
}

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue foundation.foundation.INSDate, proposedTimeInterval float64)
	HasDatePickerCellValidateProposedDateValueTimeInterval() bool
}

// DatePickerCellDelegate is a delegate implementation builder for the PDatePickerCellDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type DatePickerCellDelegate struct {
	_DatePickerCellValidateProposedDateValueTimeInterval func(datePickerCell IDatePickerCell, proposedDateValue foundation.foundation.INSDate, proposedTimeInterval float64)
}

// SetDatePickerCellValidateProposedDateValueTimeInterval sets the handler for the DatePickerCellValidateProposedDateValueTimeInterval delegate method.
//
// The delegate receives this message each time the user attempts to change the receiver’s value, allowing the delegate the opportunity to override the change.
func (d *DatePickerCellDelegate) SetDatePickerCellValidateProposedDateValueTimeInterval(f func(datePickerCell IDatePickerCell, proposedDateValue foundation.foundation.INSDate, proposedTimeInterval float64)) {
	d._DatePickerCellValidateProposedDateValueTimeInterval = f
}

// DatePickerCellValidateProposedDateValueTimeInterval implements the PDatePickerCellDelegate interface.
func (d *DatePickerCellDelegate) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue foundation.foundation.INSDate, proposedTimeInterval float64) {
	if d._DatePickerCellValidateProposedDateValueTimeInterval != nil {
		d._DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell, proposedDateValue, proposedTimeInterval)
	}
}

// HasDatePickerCellValidateProposedDateValueTimeInterval returns true if a handler for DatePickerCellValidateProposedDateValueTimeInterval has been set.
func (d *DatePickerCellDelegate) HasDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return d._DatePickerCellValidateProposedDateValueTimeInterval != nil
}

// DatePickerCellDelegateObject wraps an existing Objective-C object that conforms to the PDatePickerCellDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type DatePickerCellDelegateObject struct {
	objectivec.Object
}

// NewDatePickerCellDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSDatePickerCellDelegate protocol.
func NewDatePickerCellDelegateObject(obj objectivec.Object) *DatePickerCellDelegateObject {
	return &DatePickerCellDelegateObject{obj}
}

// Make sure DatePickerCellDelegateObject implements PDatePickerCellDelegate.
var _ PDatePickerCellDelegate = (*DatePickerCellDelegateObject)(nil)

// DatePickerCellValidateProposedDateValueTimeInterval implements the PDatePickerCellDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *DatePickerCellDelegateObject) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue foundation.foundation.INSDate, proposedTimeInterval float64) {
	objc.Send[objc.ID](o.ID, objc.Sel("datePickerCell:validateProposedDateValue:timeInterval:"), datePickerCell, proposedDateValue, proposedTimeInterval)
}

// HasDatePickerCellValidateProposedDateValueTimeInterval returns true; this is a placeholder for optional method checks.
func (o *DatePickerCellDelegateObject) HasDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

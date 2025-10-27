// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PComboBoxDelegate is the NSComboBoxDelegate protocol interface.
//
// A set of optional methods implemented by delegates of combo box objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSComboBoxDelegate
type PComboBoxDelegate interface {
	// Optional methods
	ComboBoxSelectionDidChange(notification foundation.foundation.INSNotification)
	HasComboBoxSelectionDidChange() bool
	ComboBoxSelectionIsChanging(notification foundation.foundation.INSNotification)
	HasComboBoxSelectionIsChanging() bool
	ComboBoxWillDismiss(notification foundation.foundation.INSNotification)
	HasComboBoxWillDismiss() bool
	ComboBoxWillPopUp(notification foundation.foundation.INSNotification)
	HasComboBoxWillPopUp() bool
}

// ComboBoxDelegate is a delegate implementation builder for the PComboBoxDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ComboBoxDelegate struct {
	_ComboBoxSelectionDidChange func(notification foundation.foundation.INSNotification)
	_ComboBoxSelectionIsChanging func(notification foundation.foundation.INSNotification)
	_ComboBoxWillDismiss func(notification foundation.foundation.INSNotification)
	_ComboBoxWillPopUp func(notification foundation.foundation.INSNotification)
}

// SetComboBoxSelectionDidChange sets the handler for the ComboBoxSelectionDidChange delegate method.
//
// Informs the delegate that the pop-up list selection has finished changing.
func (d *ComboBoxDelegate) SetComboBoxSelectionDidChange(f func(notification foundation.foundation.INSNotification)) {
	d._ComboBoxSelectionDidChange = f
}

// SetComboBoxSelectionIsChanging sets the handler for the ComboBoxSelectionIsChanging delegate method.
//
// Informs the delegate that the pop-up list selection is changing.
func (d *ComboBoxDelegate) SetComboBoxSelectionIsChanging(f func(notification foundation.foundation.INSNotification)) {
	d._ComboBoxSelectionIsChanging = f
}

// SetComboBoxWillDismiss sets the handler for the ComboBoxWillDismiss delegate method.
//
// Informs the delegate that the pop-up list is about to be dismissed.
func (d *ComboBoxDelegate) SetComboBoxWillDismiss(f func(notification foundation.foundation.INSNotification)) {
	d._ComboBoxWillDismiss = f
}

// SetComboBoxWillPopUp sets the handler for the ComboBoxWillPopUp delegate method.
//
// Informs the delegate that the pop-up list is about to be displayed.
func (d *ComboBoxDelegate) SetComboBoxWillPopUp(f func(notification foundation.foundation.INSNotification)) {
	d._ComboBoxWillPopUp = f
}

// ComboBoxSelectionDidChange implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxSelectionDidChange(notification foundation.foundation.INSNotification) {
	if d._ComboBoxSelectionDidChange != nil {
		d._ComboBoxSelectionDidChange(notification)
	}
}

// HasComboBoxSelectionDidChange returns true if a handler for ComboBoxSelectionDidChange has been set.
func (d *ComboBoxDelegate) HasComboBoxSelectionDidChange() bool {
	return d._ComboBoxSelectionDidChange != nil
}

// ComboBoxSelectionIsChanging implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxSelectionIsChanging(notification foundation.foundation.INSNotification) {
	if d._ComboBoxSelectionIsChanging != nil {
		d._ComboBoxSelectionIsChanging(notification)
	}
}

// HasComboBoxSelectionIsChanging returns true if a handler for ComboBoxSelectionIsChanging has been set.
func (d *ComboBoxDelegate) HasComboBoxSelectionIsChanging() bool {
	return d._ComboBoxSelectionIsChanging != nil
}

// ComboBoxWillDismiss implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxWillDismiss(notification foundation.foundation.INSNotification) {
	if d._ComboBoxWillDismiss != nil {
		d._ComboBoxWillDismiss(notification)
	}
}

// HasComboBoxWillDismiss returns true if a handler for ComboBoxWillDismiss has been set.
func (d *ComboBoxDelegate) HasComboBoxWillDismiss() bool {
	return d._ComboBoxWillDismiss != nil
}

// ComboBoxWillPopUp implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxWillPopUp(notification foundation.foundation.INSNotification) {
	if d._ComboBoxWillPopUp != nil {
		d._ComboBoxWillPopUp(notification)
	}
}

// HasComboBoxWillPopUp returns true if a handler for ComboBoxWillPopUp has been set.
func (d *ComboBoxDelegate) HasComboBoxWillPopUp() bool {
	return d._ComboBoxWillPopUp != nil
}

// ComboBoxDelegateObject wraps an existing Objective-C object that conforms to the PComboBoxDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type ComboBoxDelegateObject struct {
	objectivec.Object
}

// NewComboBoxDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSComboBoxDelegate protocol.
func NewComboBoxDelegateObject(obj objectivec.Object) *ComboBoxDelegateObject {
	return &ComboBoxDelegateObject{obj}
}

// Make sure ComboBoxDelegateObject implements PComboBoxDelegate.
var _ PComboBoxDelegate = (*ComboBoxDelegateObject)(nil)

// ComboBoxSelectionDidChange implements the PComboBoxDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDelegateObject) ComboBoxSelectionDidChange(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("comboBoxSelectionDidChange:"), notification)
}

// HasComboBoxSelectionDidChange returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDelegateObject) HasComboBoxSelectionDidChange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ComboBoxSelectionIsChanging implements the PComboBoxDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDelegateObject) ComboBoxSelectionIsChanging(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("comboBoxSelectionIsChanging:"), notification)
}

// HasComboBoxSelectionIsChanging returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDelegateObject) HasComboBoxSelectionIsChanging() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ComboBoxWillDismiss implements the PComboBoxDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDelegateObject) ComboBoxWillDismiss(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("comboBoxWillDismiss:"), notification)
}

// HasComboBoxWillDismiss returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDelegateObject) HasComboBoxWillDismiss() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// ComboBoxWillPopUp implements the PComboBoxDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ComboBoxDelegateObject) ComboBoxWillPopUp(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("comboBoxWillPopUp:"), notification)
}

// HasComboBoxWillPopUp returns true; this is a placeholder for optional method checks.
func (o *ComboBoxDelegateObject) HasComboBoxWillPopUp() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

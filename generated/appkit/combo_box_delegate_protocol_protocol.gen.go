// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	ComboBoxSelectionDidChange(notification foundation.Notification)
	HasComboBoxSelectionDidChange() bool
	ComboBoxSelectionIsChanging(notification foundation.Notification)
	HasComboBoxSelectionIsChanging() bool
	ComboBoxWillDismiss(notification foundation.Notification)
	HasComboBoxWillDismiss() bool
	ComboBoxWillPopUp(notification foundation.Notification)
	HasComboBoxWillPopUp() bool
}

// ComboBoxDelegate is a delegate implementation builder for the PComboBoxDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ComboBoxDelegate struct {
	_ComboBoxSelectionDidChange func(notification foundation.Notification)
	_ComboBoxSelectionIsChanging func(notification foundation.Notification)
	_ComboBoxWillDismiss func(notification foundation.Notification)
	_ComboBoxWillPopUp func(notification foundation.Notification)
}

// SetComboBoxSelectionDidChange sets the handler for the ComboBoxSelectionDidChange delegate method.
//
// Informs the delegate that the pop-up list selection has finished changing.
func (d *ComboBoxDelegate) SetComboBoxSelectionDidChange(f func(notification foundation.Notification)) {
	d._ComboBoxSelectionDidChange = f
}

// SetComboBoxSelectionIsChanging sets the handler for the ComboBoxSelectionIsChanging delegate method.
//
// Informs the delegate that the pop-up list selection is changing.
func (d *ComboBoxDelegate) SetComboBoxSelectionIsChanging(f func(notification foundation.Notification)) {
	d._ComboBoxSelectionIsChanging = f
}

// SetComboBoxWillDismiss sets the handler for the ComboBoxWillDismiss delegate method.
//
// Informs the delegate that the pop-up list is about to be dismissed.
func (d *ComboBoxDelegate) SetComboBoxWillDismiss(f func(notification foundation.Notification)) {
	d._ComboBoxWillDismiss = f
}

// SetComboBoxWillPopUp sets the handler for the ComboBoxWillPopUp delegate method.
//
// Informs the delegate that the pop-up list is about to be displayed.
func (d *ComboBoxDelegate) SetComboBoxWillPopUp(f func(notification foundation.Notification)) {
	d._ComboBoxWillPopUp = f
}

// ComboBoxSelectionDidChange implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxSelectionDidChange(notification foundation.Notification) {
	if d._ComboBoxSelectionDidChange != nil {
		d._ComboBoxSelectionDidChange(notification)
	}
}

// HasComboBoxSelectionDidChange returns true if a handler for ComboBoxSelectionDidChange has been set.
func (d *ComboBoxDelegate) HasComboBoxSelectionDidChange() bool {
	return d._ComboBoxSelectionDidChange != nil
}

// ComboBoxSelectionIsChanging implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxSelectionIsChanging(notification foundation.Notification) {
	if d._ComboBoxSelectionIsChanging != nil {
		d._ComboBoxSelectionIsChanging(notification)
	}
}

// HasComboBoxSelectionIsChanging returns true if a handler for ComboBoxSelectionIsChanging has been set.
func (d *ComboBoxDelegate) HasComboBoxSelectionIsChanging() bool {
	return d._ComboBoxSelectionIsChanging != nil
}

// ComboBoxWillDismiss implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxWillDismiss(notification foundation.Notification) {
	if d._ComboBoxWillDismiss != nil {
		d._ComboBoxWillDismiss(notification)
	}
}

// HasComboBoxWillDismiss returns true if a handler for ComboBoxWillDismiss has been set.
func (d *ComboBoxDelegate) HasComboBoxWillDismiss() bool {
	return d._ComboBoxWillDismiss != nil
}

// ComboBoxWillPopUp implements the PComboBoxDelegate interface.
func (d *ComboBoxDelegate) ComboBoxWillPopUp(notification foundation.Notification) {
	if d._ComboBoxWillPopUp != nil {
		d._ComboBoxWillPopUp(notification)
	}
}

// HasComboBoxWillPopUp returns true if a handler for ComboBoxWillPopUp has been set.
func (d *ComboBoxDelegate) HasComboBoxWillPopUp() bool {
	return d._ComboBoxWillPopUp != nil
}

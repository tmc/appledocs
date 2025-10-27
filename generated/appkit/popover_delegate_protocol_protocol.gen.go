// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PPopoverDelegate is the NSPopoverDelegate protocol interface.
//
// A set of optional methods that a popover delegate can implement to provide additional or custom functionality.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPopoverDelegate
type PPopoverDelegate interface {
	// Optional methods
	PopoverDidClose(notification foundation.foundation.INSNotification)
	HasPopoverDidClose() bool
}

// PopoverDelegate is a delegate implementation builder for the PPopoverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PopoverDelegate struct {
	_PopoverDidClose func(notification foundation.foundation.INSNotification)
}

// SetPopoverDidClose sets the handler for the PopoverDidClose delegate method.
//
// Invoked when the popover did close.
func (d *PopoverDelegate) SetPopoverDidClose(f func(notification foundation.foundation.INSNotification)) {
	d._PopoverDidClose = f
}

// PopoverDidClose implements the PPopoverDelegate interface.
func (d *PopoverDelegate) PopoverDidClose(notification foundation.foundation.INSNotification) {
	if d._PopoverDidClose != nil {
		d._PopoverDidClose(notification)
	}
}

// HasPopoverDidClose returns true if a handler for PopoverDidClose has been set.
func (d *PopoverDelegate) HasPopoverDidClose() bool {
	return d._PopoverDidClose != nil
}

// PopoverDelegateObject wraps an existing Objective-C object that conforms to the PPopoverDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type PopoverDelegateObject struct {
	objectivec.Object
}

// NewPopoverDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSPopoverDelegate protocol.
func NewPopoverDelegateObject(obj objectivec.Object) *PopoverDelegateObject {
	return &PopoverDelegateObject{obj}
}

// Make sure PopoverDelegateObject implements PPopoverDelegate.
var _ PPopoverDelegate = (*PopoverDelegateObject)(nil)

// PopoverDidClose implements the PPopoverDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *PopoverDelegateObject) PopoverDidClose(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("popoverDidClose:"), notification)
}

// HasPopoverDidClose returns true; this is a placeholder for optional method checks.
func (o *PopoverDelegateObject) HasPopoverDidClose() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

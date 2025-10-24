// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	PopoverDidClose(notification foundation.Notification)
	HasPopoverDidClose() bool
}

// PopoverDelegate is a delegate implementation builder for the PPopoverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PopoverDelegate struct {
	_PopoverDidClose func(notification foundation.Notification)
}

// SetPopoverDidClose sets the handler for the PopoverDidClose delegate method.
//
// Invoked when the popover did close.
func (d *PopoverDelegate) SetPopoverDidClose(f func(notification foundation.Notification)) {
	d._PopoverDidClose = f
}

// PopoverDidClose implements the PPopoverDelegate interface.
func (d *PopoverDelegate) PopoverDidClose(notification foundation.Notification) {
	if d._PopoverDidClose != nil {
		d._PopoverDidClose(notification)
	}
}

// HasPopoverDidClose returns true if a handler for PopoverDidClose has been set.
func (d *PopoverDelegate) HasPopoverDidClose() bool {
	return d._PopoverDidClose != nil
}

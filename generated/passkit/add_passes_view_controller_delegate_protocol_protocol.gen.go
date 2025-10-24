// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAddPassesViewControllerDelegate is the PKAddPassesViewControllerDelegate protocol interface.
//
// Methods that an add-passes view controller’s delegate implements.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - visionOS +
//
// See: doc://com.apple.passkit/documentation/PassKit/PKAddPassesViewControllerDelegate
type PAddPassesViewControllerDelegate interface {
	// Optional methods
	AddPassesViewControllerDidFinish(controller AddPassesViewController /* not a class type */)
	HasAddPassesViewControllerDidFinish() bool
}

// AddPassesViewControllerDelegate is a delegate implementation builder for the PAddPassesViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AddPassesViewControllerDelegate struct {
	_AddPassesViewControllerDidFinish func(controller AddPassesViewController /* not a class type */)
}

// SetAddPassesViewControllerDidFinish sets the handler for the AddPassesViewControllerDidFinish delegate method.
//
// Sent to the delegate after the add-passes view controller has finished.
func (d *AddPassesViewControllerDelegate) SetAddPassesViewControllerDidFinish(f func(controller AddPassesViewController /* not a class type */)) {
	d._AddPassesViewControllerDidFinish = f
}

// AddPassesViewControllerDidFinish implements the PAddPassesViewControllerDelegate interface.
func (d *AddPassesViewControllerDelegate) AddPassesViewControllerDidFinish(controller AddPassesViewController /* not a class type */) {
	if d._AddPassesViewControllerDidFinish != nil {
		d._AddPassesViewControllerDidFinish(controller)
	}
}

// HasAddPassesViewControllerDidFinish returns true if a handler for AddPassesViewControllerDidFinish has been set.
func (d *AddPassesViewControllerDelegate) HasAddPassesViewControllerDidFinish() bool {
	return d._AddPassesViewControllerDidFinish != nil
}

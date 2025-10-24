// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PRPBroadcastActivityControllerDelegate is the RPBroadcastActivityControllerDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to selection events from a broadcast activity controller.
//
// Availability:
//   - macOS 11.0+
//
// See: doc://com.apple.replaykit/documentation/ReplayKit/RPBroadcastActivityControllerDelegate
type PRPBroadcastActivityControllerDelegate interface {
	// Required methods
	BroadcastActivityControllerDidFinishWithBroadcastControllerError(broadcastActivityController IRPBroadcastActivityController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: BroadcastActivityControllerDidFinishWithBroadcastControllerError */
}

// RPBroadcastActivityControllerDelegate is a delegate implementation builder for the PRPBroadcastActivityControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RPBroadcastActivityControllerDelegate struct {
	_BroadcastActivityControllerDidFinishWithBroadcastControllerError func(broadcastActivityController IRPBroadcastActivityController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)
}

// SetBroadcastActivityControllerDidFinishWithBroadcastControllerError sets the handler for the BroadcastActivityControllerDidFinishWithBroadcastControllerError delegate method.
//
// Tells the delegate that a user selected a broadcast.
func (d *RPBroadcastActivityControllerDelegate) SetBroadcastActivityControllerDidFinishWithBroadcastControllerError(f func(broadcastActivityController IRPBroadcastActivityController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)) {
	d._BroadcastActivityControllerDidFinishWithBroadcastControllerError = f
}

// BroadcastActivityControllerDidFinishWithBroadcastControllerError implements the PRPBroadcastActivityControllerDelegate interface.
func (d *RPBroadcastActivityControllerDelegate) BroadcastActivityControllerDidFinishWithBroadcastControllerError(broadcastActivityController IRPBroadcastActivityController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */) {
	if d._BroadcastActivityControllerDidFinishWithBroadcastControllerError != nil {
		d._BroadcastActivityControllerDidFinishWithBroadcastControllerError(broadcastActivityController, broadcastController, error_)
	}
}

// HasBroadcastActivityControllerDidFinishWithBroadcastControllerError returns true if a handler for BroadcastActivityControllerDidFinishWithBroadcastControllerError has been set.
func (d *RPBroadcastActivityControllerDelegate) HasBroadcastActivityControllerDidFinishWithBroadcastControllerError() bool {
	return d._BroadcastActivityControllerDidFinishWithBroadcastControllerError != nil
}

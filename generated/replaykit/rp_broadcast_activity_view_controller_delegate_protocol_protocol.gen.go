// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PRPBroadcastActivityViewControllerDelegate is the RPBroadcastActivityViewControllerDelegate protocol interface.
//
// The protocol you implement to respond to changes to a broadcast activity user interface.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.replaykit/documentation/ReplayKit/RPBroadcastActivityViewControllerDelegate
type PRPBroadcastActivityViewControllerDelegate interface {
	// Required methods
	BroadcastActivityViewControllerDidFinishWithBroadcastControllerError(broadcastActivityViewController IRPBroadcastActivityViewController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: BroadcastActivityViewControllerDidFinishWithBroadcastControllerError */
}

// RPBroadcastActivityViewControllerDelegate is a delegate implementation builder for the PRPBroadcastActivityViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RPBroadcastActivityViewControllerDelegate struct {
	_BroadcastActivityViewControllerDidFinishWithBroadcastControllerError func(broadcastActivityViewController IRPBroadcastActivityViewController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)
}

// SetBroadcastActivityViewControllerDidFinishWithBroadcastControllerError sets the handler for the BroadcastActivityViewControllerDidFinishWithBroadcastControllerError delegate method.
//
// Indicates that the broadcast activity view controller is ready to be dismissed.
func (d *RPBroadcastActivityViewControllerDelegate) SetBroadcastActivityViewControllerDidFinishWithBroadcastControllerError(f func(broadcastActivityViewController IRPBroadcastActivityViewController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)) {
	d._BroadcastActivityViewControllerDidFinishWithBroadcastControllerError = f
}

// BroadcastActivityViewControllerDidFinishWithBroadcastControllerError implements the PRPBroadcastActivityViewControllerDelegate interface.
func (d *RPBroadcastActivityViewControllerDelegate) BroadcastActivityViewControllerDidFinishWithBroadcastControllerError(broadcastActivityViewController IRPBroadcastActivityViewController, broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */) {
	if d._BroadcastActivityViewControllerDidFinishWithBroadcastControllerError != nil {
		d._BroadcastActivityViewControllerDidFinishWithBroadcastControllerError(broadcastActivityViewController, broadcastController, error_)
	}
}

// HasBroadcastActivityViewControllerDidFinishWithBroadcastControllerError returns true if a handler for BroadcastActivityViewControllerDidFinishWithBroadcastControllerError has been set.
func (d *RPBroadcastActivityViewControllerDelegate) HasBroadcastActivityViewControllerDidFinishWithBroadcastControllerError() bool {
	return d._BroadcastActivityViewControllerDidFinishWithBroadcastControllerError != nil
}

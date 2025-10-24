// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PRPBroadcastControllerDelegate is the RPBroadcastControllerDelegate protocol interface.
//
// The protocol you implement to respond to changes in a live broadcast.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.replaykit/documentation/ReplayKit/RPBroadcastControllerDelegate
type PRPBroadcastControllerDelegate interface {
	// Optional methods
	BroadcastControllerDidFinishWithError(broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)
	HasBroadcastControllerDidFinishWithError() bool
	BroadcastControllerDidUpdateBroadcastURL(broadcastController IRPBroadcastController, broadcastURL objc.IObject /* cross-framework: NSURL */)
	HasBroadcastControllerDidUpdateBroadcastURL() bool
	BroadcastControllerDidUpdateServiceInfo(broadcastController IRPBroadcastController, serviceInfo foundation.IDictionary)
	HasBroadcastControllerDidUpdateServiceInfo() bool
}

// RPBroadcastControllerDelegate is a delegate implementation builder for the PRPBroadcastControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RPBroadcastControllerDelegate struct {
	_BroadcastControllerDidFinishWithError func(broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)
	_BroadcastControllerDidUpdateBroadcastURL func(broadcastController IRPBroadcastController, broadcastURL objc.IObject /* cross-framework: NSURL */)
	_BroadcastControllerDidUpdateServiceInfo func(broadcastController IRPBroadcastController, serviceInfo foundation.IDictionary)
}

// SetBroadcastControllerDidFinishWithError sets the handler for the BroadcastControllerDidFinishWithError delegate method.
//
// Tells the delegate that a broadcast ended due to an error.
func (d *RPBroadcastControllerDelegate) SetBroadcastControllerDidFinishWithError(f func(broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */)) {
	d._BroadcastControllerDidFinishWithError = f
}

// SetBroadcastControllerDidUpdateBroadcastURL sets the handler for the BroadcastControllerDidUpdateBroadcastURL delegate method.
//
// Tells the broadcast service the broadcast URL has been updated.
func (d *RPBroadcastControllerDelegate) SetBroadcastControllerDidUpdateBroadcastURL(f func(broadcastController IRPBroadcastController, broadcastURL objc.IObject /* cross-framework: NSURL */)) {
	d._BroadcastControllerDidUpdateBroadcastURL = f
}

// SetBroadcastControllerDidUpdateServiceInfo sets the handler for the BroadcastControllerDidUpdateServiceInfo delegate method.
//
// Tells the delegate the broadcast service has data to pass back to the broadcasting app.
func (d *RPBroadcastControllerDelegate) SetBroadcastControllerDidUpdateServiceInfo(f func(broadcastController IRPBroadcastController, serviceInfo foundation.IDictionary)) {
	d._BroadcastControllerDidUpdateServiceInfo = f
}

// BroadcastControllerDidFinishWithError implements the PRPBroadcastControllerDelegate interface.
func (d *RPBroadcastControllerDelegate) BroadcastControllerDidFinishWithError(broadcastController IRPBroadcastController, error_ objc.IObject /* cross-framework: Error */) {
	if d._BroadcastControllerDidFinishWithError != nil {
		d._BroadcastControllerDidFinishWithError(broadcastController, error_)
	}
}

// HasBroadcastControllerDidFinishWithError returns true if a handler for BroadcastControllerDidFinishWithError has been set.
func (d *RPBroadcastControllerDelegate) HasBroadcastControllerDidFinishWithError() bool {
	return d._BroadcastControllerDidFinishWithError != nil
}

// BroadcastControllerDidUpdateBroadcastURL implements the PRPBroadcastControllerDelegate interface.
func (d *RPBroadcastControllerDelegate) BroadcastControllerDidUpdateBroadcastURL(broadcastController IRPBroadcastController, broadcastURL objc.IObject /* cross-framework: NSURL */) {
	if d._BroadcastControllerDidUpdateBroadcastURL != nil {
		d._BroadcastControllerDidUpdateBroadcastURL(broadcastController, broadcastURL)
	}
}

// HasBroadcastControllerDidUpdateBroadcastURL returns true if a handler for BroadcastControllerDidUpdateBroadcastURL has been set.
func (d *RPBroadcastControllerDelegate) HasBroadcastControllerDidUpdateBroadcastURL() bool {
	return d._BroadcastControllerDidUpdateBroadcastURL != nil
}

// BroadcastControllerDidUpdateServiceInfo implements the PRPBroadcastControllerDelegate interface.
func (d *RPBroadcastControllerDelegate) BroadcastControllerDidUpdateServiceInfo(broadcastController IRPBroadcastController, serviceInfo foundation.IDictionary) {
	if d._BroadcastControllerDidUpdateServiceInfo != nil {
		d._BroadcastControllerDidUpdateServiceInfo(broadcastController, serviceInfo)
	}
}

// HasBroadcastControllerDidUpdateServiceInfo returns true if a handler for BroadcastControllerDidUpdateServiceInfo has been set.
func (d *RPBroadcastControllerDelegate) HasBroadcastControllerDidUpdateServiceInfo() bool {
	return d._BroadcastControllerDidUpdateServiceInfo != nil
}

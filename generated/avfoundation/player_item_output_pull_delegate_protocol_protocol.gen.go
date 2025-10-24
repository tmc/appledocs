// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPlayerItemOutputPullDelegate is the AVPlayerItemOutputPullDelegate protocol interface.
//
// Methods you can implement to respond to pixel buffer changes.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemOutputPullDelegate
type PPlayerItemOutputPullDelegate interface {
	// Optional methods
	OutputMediaDataWillChange(sender IAVPlayerItemOutput)
	HasOutputMediaDataWillChange() bool
	OutputSequenceWasFlushed(output IAVPlayerItemOutput)
	HasOutputSequenceWasFlushed() bool
}

// PlayerItemOutputPullDelegate is a delegate implementation builder for the PPlayerItemOutputPullDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerItemOutputPullDelegate struct {
	_OutputMediaDataWillChange func(sender IAVPlayerItemOutput)
	_OutputSequenceWasFlushed func(output IAVPlayerItemOutput)
}

// SetOutputMediaDataWillChange sets the handler for the OutputMediaDataWillChange delegate method.
//
// Tells the delegate that new samples are about to arrive.
func (d *PlayerItemOutputPullDelegate) SetOutputMediaDataWillChange(f func(sender IAVPlayerItemOutput)) {
	d._OutputMediaDataWillChange = f
}

// SetOutputSequenceWasFlushed sets the handler for the OutputSequenceWasFlushed delegate method.
//
// Tells the delegate that a new sample sequence is commencing.
func (d *PlayerItemOutputPullDelegate) SetOutputSequenceWasFlushed(f func(output IAVPlayerItemOutput)) {
	d._OutputSequenceWasFlushed = f
}

// OutputMediaDataWillChange implements the PPlayerItemOutputPullDelegate interface.
func (d *PlayerItemOutputPullDelegate) OutputMediaDataWillChange(sender IAVPlayerItemOutput) {
	if d._OutputMediaDataWillChange != nil {
		d._OutputMediaDataWillChange(sender)
	}
}

// HasOutputMediaDataWillChange returns true if a handler for OutputMediaDataWillChange has been set.
func (d *PlayerItemOutputPullDelegate) HasOutputMediaDataWillChange() bool {
	return d._OutputMediaDataWillChange != nil
}

// OutputSequenceWasFlushed implements the PPlayerItemOutputPullDelegate interface.
func (d *PlayerItemOutputPullDelegate) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	if d._OutputSequenceWasFlushed != nil {
		d._OutputSequenceWasFlushed(output)
	}
}

// HasOutputSequenceWasFlushed returns true if a handler for OutputSequenceWasFlushed has been set.
func (d *PlayerItemOutputPullDelegate) HasOutputSequenceWasFlushed() bool {
	return d._OutputSequenceWasFlushed != nil
}

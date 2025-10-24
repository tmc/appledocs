// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPlayerItemOutputPushDelegate is the AVPlayerItemOutputPushDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to changes in the media data sequence.
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
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemOutputPushDelegate
type PPlayerItemOutputPushDelegate interface {
	// Optional methods
	OutputSequenceWasFlushed(output IAVPlayerItemOutput)
	HasOutputSequenceWasFlushed() bool
}

// PlayerItemOutputPushDelegate is a delegate implementation builder for the PPlayerItemOutputPushDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerItemOutputPushDelegate struct {
	_OutputSequenceWasFlushed func(output IAVPlayerItemOutput)
}

// SetOutputSequenceWasFlushed sets the handler for the OutputSequenceWasFlushed delegate method.
//
// Tells the delegate that the output is starting a new sequence of media data.
func (d *PlayerItemOutputPushDelegate) SetOutputSequenceWasFlushed(f func(output IAVPlayerItemOutput)) {
	d._OutputSequenceWasFlushed = f
}

// OutputSequenceWasFlushed implements the PPlayerItemOutputPushDelegate interface.
func (d *PlayerItemOutputPushDelegate) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	if d._OutputSequenceWasFlushed != nil {
		d._OutputSequenceWasFlushed(output)
	}
}

// HasOutputSequenceWasFlushed returns true if a handler for OutputSequenceWasFlushed has been set.
func (d *PlayerItemOutputPushDelegate) HasOutputSequenceWasFlushed() bool {
	return d._OutputSequenceWasFlushed != nil
}

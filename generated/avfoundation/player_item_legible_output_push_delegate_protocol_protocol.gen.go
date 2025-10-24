// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corevideo"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayerItemLegibleOutputPushDelegate is the AVPlayerItemLegibleOutputPushDelegate protocol interface.
//
// Methods you can implement to provide alternative attributed-string output.
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
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemLegibleOutputPushDelegate
type PPlayerItemLegibleOutputPushDelegate interface {
	// Optional methods
	LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output IAVPlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples objc.IObject /* cross-framework: NSArray */, itemTime objc.IObject /* cross-framework: Time */)
	HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool
}

// PlayerItemLegibleOutputPushDelegate is a delegate implementation builder for the PPlayerItemLegibleOutputPushDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerItemLegibleOutputPushDelegate struct {
	_LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime func(output IAVPlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples objc.IObject /* cross-framework: NSArray */, itemTime objc.IObject /* cross-framework: Time */)
}

// SetLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime sets the handler for the LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime delegate method.
//
// Asks the delegate to process the delivery of new textual samples.
func (d *PlayerItemLegibleOutputPushDelegate) SetLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(f func(output IAVPlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples objc.IObject /* cross-framework: NSArray */, itemTime objc.IObject /* cross-framework: Time */)) {
	d._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime = f
}

// LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime implements the PPlayerItemLegibleOutputPushDelegate interface.
func (d *PlayerItemLegibleOutputPushDelegate) LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output IAVPlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples objc.IObject /* cross-framework: NSArray */, itemTime objc.IObject /* cross-framework: Time */) {
	if d._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime != nil {
		d._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output, strings, nativeSamples, itemTime)
	}
}

// HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime returns true if a handler for LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime has been set.
func (d *PlayerItemLegibleOutputPushDelegate) HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool {
	return d._LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime != nil
}

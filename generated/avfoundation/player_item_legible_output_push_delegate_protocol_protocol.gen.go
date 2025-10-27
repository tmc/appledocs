// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	LegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime(output IAVPlayerItemLegibleOutput, strings []foundation.AttributedString, nativeSamples foundation.foundation.INSArray, itemTime objectivec.IObject)
	HasLegibleOutputDidOutputAttributedStringsNativeSampleBuffersForItemTime() bool
}

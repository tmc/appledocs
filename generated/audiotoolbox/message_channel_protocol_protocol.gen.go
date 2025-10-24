// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMessageChannel is the AUMessageChannel protocol interface.
//
// A specification for a bidirectional communication message channel.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.audiotoolbox/documentation/AudioToolbox/AUMessageChannel
type PMessageChannel interface {
	// Optional methods
	CallAudioUnit(message objc.IObject /* cross-framework: NSDictionary */) foundation.Dictionary
	HasCallAudioUnit() bool
}

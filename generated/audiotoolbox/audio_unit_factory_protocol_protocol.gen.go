// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PAudioUnitFactory is the AUAudioUnitFactory protocol interface.
//
// An object that creates a version 3 audio unit.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.audiotoolbox/documentation/AudioToolbox/AUAudioUnitFactory
type PAudioUnitFactory interface {
	// Required methods
	CreateAudioUnitWithComponentDescriptionError(desc objc.IObject /* cross-framework: AudioComponentDescription */, error_ objectivec.IObject) AudioUnit
}

//go:build darwin && ios

// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/coremidi"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioUnit


// iOS-only properties

// The AUAudioUnit’s intended spatial experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/intendedSpatialExperience-1dvhd
func (a_ AudioUnit) IntendedSpatialExperience() ISpatialAudioExperience {
	rv := objc.Send[SpatialAudioExperience](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a_ AudioUnit) SetIntendedSpatialExperience(value ISpatialAudioExperience) {
	a_.ID.Send(objc.RegisterName("setIntendedSpatialExperience:"), value)
}





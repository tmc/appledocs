//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for AudioOutputNode


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode/intendedSpatialExperience-3uznq
func (a_ AudioOutputNode) IntendedSpatialExperience() audiotoolbox.SpatialAudioExperience {
	rv := objc.Send[audiotoolbox.SpatialAudioExperience](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a_ AudioOutputNode) SetIntendedSpatialExperience(value audiotoolbox.SpatialAudioExperience) {
	a_.ID.Send(objc.RegisterName("setIntendedSpatialExperience:"), value)
}






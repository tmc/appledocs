//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioPlayer


// iOS-only properties

// An array of channel descriptions for the audio player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a_ AudioPlayer) ChannelAssignments() []AudioSessionChannelDescription {
	rv := objc.Send[[]AudioSessionChannelDescription](a_.ID, objc.Sel("channelAssignments"))
	return rv
}
func (a_ AudioPlayer) SetChannelAssignments(value []AudioSessionChannelDescription) {
	a_.ID.Send(objc.RegisterName("setChannelAssignments:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a_ AudioPlayer) IntendedSpatialExperience() audiotoolbox.SpatialAudioExperience {
	rv := objc.Send[audiotoolbox.SpatialAudioExperience](a_.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a_ AudioPlayer) SetIntendedSpatialExperience(value audiotoolbox.SpatialAudioExperience) {
	a_.ID.Send(objc.RegisterName("setIntendedSpatialExperience:"), value)
}





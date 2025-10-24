//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioRecorder


// iOS-only properties

// An array of channel descriptions associated with the audio recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/channelAssignments
func (a_ AudioRecorder) ChannelAssignments() []AudioSessionChannelDescription {
	rv := objc.Send[[]AudioSessionChannelDescription](a_.ID, objc.Sel("channelAssignments"))
	return rv
}
func (a_ AudioRecorder) SetChannelAssignments(value []AudioSessionChannelDescription) {
	a_.ID.Send(objc.RegisterName("setChannelAssignments:"), value)
}




